package dkg

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/onflow/crypto"

	"github.com/onflow/flow-go/model/flow"
	msig "github.com/onflow/flow-go/module/signature"
	"github.com/onflow/flow-go/utils/unittest"
)

func TestWithEmulator(t *testing.T) {
	suite.Run(t, new(EmulatorSuite))
}

func (s *EmulatorSuite) runTest(goodNodes int, emulatorProblems bool) {

	nodes := s.nodes[:goodNodes]

	// The EpochSetup event is received at view 100.  The current epoch is
	// configured with phase transitions at views 150, 200, and 250. In between
	// phase transitions, the controller calls the DKG smart-contract every 10
	// views.
	//
	// VIEWS
	// setup      : 100
	// polling    : 110 120 130 140 150
	// Phase1Final: 150
	// polling    : 160 170 180 190 200
	// Phase2Final: 200
	// polling    : 210 220 230 240 250
	// Phase3Final: 250
	// final

	// we arbitrarily use 999 as the current epoch counter
	currentCounter := uint64(999)

	currentEpochSetup := flow.EpochSetup{
		Counter:            currentCounter,
		FirstView:          0,
		DKGPhase1FinalView: 150,
		DKGPhase2FinalView: 200,
		DKGPhase3FinalView: 250,
		FinalView:          300,
		Participants:       s.netIDs.ToSkeleton(),
		RandomSource:       unittest.EpochSetupRandomSourceFixture(),
	}

	// create the EpochSetup that will trigger the next DKG run with all the
	// desired parameters
	nextEpochSetup := flow.EpochSetup{
		Counter:      currentCounter + 1,
		Participants: s.netIDs.ToSkeleton(),
		RandomSource: unittest.EpochSetupRandomSourceFixture(),
		FirstView:    301,
		FinalView:    600,
	}

	firstBlock := &flow.Header{View: 100}

	for _, node := range nodes {
		node.setEpochs(s.T(), currentEpochSetup, nextEpochSetup, firstBlock)
		node.Start()
		unittest.RequireCloseBefore(s.T(), node.Ready(), time.Second, "failed to start up")
	}

	// trigger the EpochSetupPhaseStarted event for all nodes, effectively
	// starting the next DKG run
	for _, n := range nodes {
		n.ProtocolEvents.EpochSetupPhaseStarted(currentCounter, firstBlock)
	}

	// submit a lot of dummy transactions to force the creation of blocks and
	// views
	view := 0
	for view < 300 {
		time.Sleep(100 * time.Millisecond)

		// if we are testing situations where the DKG smart-contract is not
		// reachable, disable the DKG client for intervals of 10 views
		if emulatorProblems {
			for _, node := range nodes {
				if view%20 >= 10 {
					node.dkgContractClient.Disable()
				} else {
					node.dkgContractClient.Enable()
				}
			}
		}

		// deliver private messages
		s.hub.DeliverAll()

		// submit a tx to force the emulator to create and finalize a block
		block, err := s.sendDummyTx()

		if err == nil {
			for _, node := range nodes {
				node.ProtocolEvents.BlockFinalized(block.Header)
			}
			view = int(block.Header.View)
		}
	}

	// before ending the test and awaiting successful completion, ensure we leave
	// the dkg client in an enabled state
	for _, node := range nodes {
		node.dkgContractClient.Enable()
	}

	for _, n := range nodes {
		n.Stop()
		unittest.RequireCloseBefore(s.T(), n.Done(), time.Second, "nodes did not shutdown")
	}

	// DKG is completed if one value was proposed by a majority of nodes
	completed := s.isDKGCompleted()
	assert.True(s.T(), completed)

	// the result is an array of public keys where the first item is the group
	// public key
	_, groupPubKey, pubKeys := s.getParametersAndResult()

	tag := "some tag"
	hasher := msig.NewBLSHasher(tag)
	// create and test a threshold signature with the keys computed by dkg
	sigData := []byte("message to be signed")
	signatures := []crypto.Signature{}
	indices := []int{}
	for i, n := range nodes {
		beaconKey, err := n.dkgState.UnsafeRetrieveMyBeaconPrivateKey(nextEpochSetup.Counter)
		require.NoError(s.T(), err)

		signature, err := beaconKey.Sign(sigData, hasher)
		require.NoError(s.T(), err)

		signatures = append(signatures, signature)
		indices = append(indices, i)

		ok, err := pubKeys[i].Verify(signature, sigData, hasher)
		require.NoError(s.T(), err)
		assert.True(s.T(), ok, fmt.Sprintf("signature %d share doesn't verify under the public key share", i))
	}

	// shuffle the signatures and indices before constructing the group
	// signature (since it only uses the first half signatures)
	rand.Shuffle(len(signatures), func(i, j int) {
		signatures[i], signatures[j] = signatures[j], signatures[i]
		indices[i], indices[j] = indices[j], indices[i]
	})

	threshold := msig.RandomBeaconThreshold(numberOfNodes)
	groupSignature, err := crypto.BLSReconstructThresholdSignature(numberOfNodes, threshold, signatures, indices)
	require.NoError(s.T(), err)

	ok, err := groupPubKey.Verify(groupSignature, sigData, hasher)
	require.NoError(s.T(), err)
	assert.True(s.T(), ok, "failed to verify threshold signature")
}

// TestHappyPath checks that DKG works when all nodes are good
func (s *EmulatorSuite) TestHappyPath() {
	s.runTest(numberOfNodes, false)
}

// TestNodesDown checks that DKG still works with the maximum number of bad
// nodes.
func (s *EmulatorSuite) TestNodesDown() {
	minHonestNodes := numberOfNodes - msig.RandomBeaconThreshold(numberOfNodes)
	s.runTest(minHonestNodes, false)
}

// TestEmulatorProblems checks that DKG is resilient to transient problems
// between the node and the DKG smart-contract ( this covers connection issues
// between consensus node and access node, as well as connection issues between
// access node and execution node, or the execution node being down).
func (s *EmulatorSuite) TestEmulatorProblems() {
	s.runTest(numberOfNodes, true)
}
