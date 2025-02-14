package inmem

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/model/flow/factory"
	"github.com/onflow/flow-go/model/flow/filter"
	"github.com/onflow/flow-go/module/signature"
	"github.com/onflow/flow-go/state/cluster"
	"github.com/onflow/flow-go/state/protocol"
)

// Epochs provides access to epoch data, backed by a rich epoch protocol state entry.
type Epochs struct {
	entry flow.RichEpochStateEntry
}

var _ protocol.EpochQuery = (*Epochs)(nil)

// Previous returns the previous epoch as of this snapshot. Valid snapshots
// must have a previous epoch for all epochs except that immediately after the root block.
// Error returns:
//   - [protocol.ErrNoPreviousEpoch] - if the previous epoch does not exist.
//     This happens when the previous epoch is queried within the first epoch of a spork.
func (eq Epochs) Previous() (protocol.CommittedEpoch, error) {
	if eq.entry.PreviousEpoch == nil {
		return nil, protocol.ErrNoPreviousEpoch
	}
	return NewCommittedEpoch(eq.entry.PreviousEpochSetup, eq.entry.PreviousEpoch.EpochExtensions, eq.entry.PreviousEpochCommit), nil
}

// Current returns the current epoch as of this snapshot. All valid snapshots have a current epoch.
func (eq Epochs) Current() (protocol.CommittedEpoch, error) {
	return NewCommittedEpoch(eq.entry.CurrentEpochSetup, eq.entry.CurrentEpoch.EpochExtensions, eq.entry.CurrentEpochCommit), nil
}

// NextUnsafe should only be used by components that are actively involved in advancing
// the epoch from [flow.EpochPhaseSetup] to [flow.EpochPhaseCommitted].
// NextUnsafe returns the tentative configuration for the next epoch as of this snapshot.
// Valid snapshots make such configuration available during the Epoch Setup Phase, which
// generally is the case only after an `EpochSetupPhaseStarted` notification has been emitted.
// CAUTION: epoch transition might not happen as described by the tentative configuration!
//
// Error returns:
//   - [ErrNextEpochNotSetup] in the case that this method is queried w.r.t. a snapshot
//     within the [flow.EpochPhaseStaking] phase or when we are in Epoch Fallback Mode.
//   - [ErrNextEpochAlreadyCommitted] if the tentative epoch is requested from
//     a snapshot within the [flow.EpochPhaseCommitted] phase.
//   - generic error in case of unexpected critical internal corruption or bugs
func (eq Epochs) NextUnsafe() (protocol.TentativeEpoch, error) {
	switch eq.entry.EpochPhase() {
	case flow.EpochPhaseStaking, flow.EpochPhaseFallback:
		return nil, protocol.ErrNextEpochNotSetup
	case flow.EpochPhaseSetup:
		return NewSetupEpoch(eq.entry.NextEpochSetup, eq.entry.NextEpoch.EpochExtensions), nil
	case flow.EpochPhaseCommitted:
		return nil, protocol.ErrNextEpochAlreadyCommitted
	}
	return nil, fmt.Errorf("unexpected unknown phase in protocol state entry")
}

// NextCommitted returns the next epoch as of this snapshot, only if it has
// been committed already - generally that is the case only after an
// `EpochCommittedPhaseStarted` notification has been emitted.
//
// Error returns:
//   - [ErrNextEpochNotCommitted] - in the case that committed epoch has been requested w.r.t a snapshot within
//     the [flow.EpochPhaseStaking] or [flow.EpochPhaseSetup] phases.
//   - generic error in case of unexpected critical internal corruption or bugs
func (eq Epochs) NextCommitted() (protocol.CommittedEpoch, error) {
	switch eq.entry.EpochPhase() {
	case flow.EpochPhaseStaking, flow.EpochPhaseFallback, flow.EpochPhaseSetup:
		return nil, protocol.ErrNextEpochNotCommitted
	case flow.EpochPhaseCommitted:
		return NewCommittedEpoch(eq.entry.NextEpochSetup, eq.entry.NextEpoch.EpochExtensions, eq.entry.NextEpochCommit), nil
	}
	return nil, fmt.Errorf("unexpected unknown phase in protocol state entry")
}

// setupEpoch is an implementation of protocol.TentativeEpoch backed by an EpochSetup service event.
// Includes any extensions which have been included as of the reference block.
type setupEpoch struct {
	// EpochSetup service event
	setupEvent *flow.EpochSetup
	extensions []flow.EpochExtension
}

var _ protocol.TentativeEpoch = (*setupEpoch)(nil)

func (es *setupEpoch) Counter() (uint64, error) {
	return es.setupEvent.Counter, nil
}

func (es *setupEpoch) FirstView() (uint64, error) {
	return es.setupEvent.FirstView, nil
}

func (es *setupEpoch) DKGPhase1FinalView() (uint64, error) {
	return es.setupEvent.DKGPhase1FinalView, nil
}

func (es *setupEpoch) DKGPhase2FinalView() (uint64, error) {
	return es.setupEvent.DKGPhase2FinalView, nil
}

func (es *setupEpoch) DKGPhase3FinalView() (uint64, error) {
	return es.setupEvent.DKGPhase3FinalView, nil
}

// FinalView returns the final view of the epoch, taking into account possible epoch extensions.
// If there are no epoch extensions, the final view is the final view of the current epoch setup,
// otherwise it is the final view of the last epoch extension.
func (es *setupEpoch) FinalView() (uint64, error) {
	if len(es.extensions) > 0 {
		return es.extensions[len(es.extensions)-1].FinalView, nil
	}
	return es.setupEvent.FinalView, nil
}

// TargetDuration returns the desired real-world duration for this epoch, in seconds.
// This target is specified by the FlowEpoch smart contract in the EpochSetup event
// and used by the Cruise Control system to moderate the block rate.
// In case the epoch has extensions, the target duration is calculated based on the last extension, by calculating how many
// views were added by the extension and adding the proportional time to the target duration.
func (es *setupEpoch) TargetDuration() (uint64, error) {
	if len(es.extensions) == 0 {
		return es.setupEvent.TargetDuration, nil
	} else {
		viewDuration := float64(es.setupEvent.TargetDuration) / float64(es.setupEvent.FinalView-es.setupEvent.FirstView+1)
		lastExtension := es.extensions[len(es.extensions)-1]
		return es.setupEvent.TargetDuration + uint64(float64(lastExtension.FinalView-es.setupEvent.FinalView)*viewDuration), nil
	}
}

// TargetEndTime returns the desired real-world end time for this epoch, represented as
// Unix Time (in units of seconds). This target is specified by the FlowEpoch smart contract in
// the EpochSetup event and used by the Cruise Control system to moderate the block rate.
// In case the epoch has extensions, the target end time is calculated based on the last extension, by calculating how many
// views were added by the extension and adding the proportional time to the target end time.
func (es *setupEpoch) TargetEndTime() (uint64, error) {
	if len(es.extensions) == 0 {
		return es.setupEvent.TargetEndTime, nil
	} else {
		viewDuration := float64(es.setupEvent.TargetDuration) / float64(es.setupEvent.FinalView-es.setupEvent.FirstView+1)
		lastExtension := es.extensions[len(es.extensions)-1]
		return es.setupEvent.TargetEndTime + uint64(float64(lastExtension.FinalView-es.setupEvent.FinalView)*viewDuration), nil
	}
}

func (es *setupEpoch) RandomSource() ([]byte, error) {
	return es.setupEvent.RandomSource, nil
}

func (es *setupEpoch) InitialIdentities() (flow.IdentitySkeletonList, error) {
	return es.setupEvent.Participants, nil
}

func (es *setupEpoch) Clustering() (flow.ClusterList, error) {
	return ClusteringFromSetupEvent(es.setupEvent)
}

func ClusteringFromSetupEvent(setupEvent *flow.EpochSetup) (flow.ClusterList, error) {
	collectorFilter := filter.HasRole[flow.IdentitySkeleton](flow.RoleCollection)
	clustering, err := factory.NewClusterList(setupEvent.Assignments, setupEvent.Participants.Filter(collectorFilter))
	if err != nil {
		return nil, fmt.Errorf("failed to generate ClusterList from collector identities: %w", err)
	}
	return clustering, nil
}

func (es *setupEpoch) Cluster(_ uint) (protocol.Cluster, error) {
	return nil, protocol.ErrNextEpochNotCommitted
}

func (es *setupEpoch) ClusterByChainID(_ flow.ChainID) (protocol.Cluster, error) {
	return nil, protocol.ErrNextEpochNotCommitted
}

func (es *setupEpoch) DKG() (protocol.DKG, error) {
	return nil, protocol.ErrNextEpochNotCommitted
}

func (es *setupEpoch) FirstHeight() (uint64, error) {
	return 0, protocol.ErrUnknownEpochBoundary
}

func (es *setupEpoch) FinalHeight() (uint64, error) {
	return 0, protocol.ErrUnknownEpochBoundary
}

// committedEpoch is an implementation of protocol.CommittedEpoch backed by an EpochSetup
// and EpochCommit service event.
type committedEpoch struct {
	setupEpoch
	commitEvent *flow.EpochCommit
}

var _ protocol.CommittedEpoch = (*committedEpoch)(nil)

func (es *committedEpoch) Cluster(index uint) (protocol.Cluster, error) {

	epochCounter := es.setupEvent.Counter

	clustering, err := es.Clustering()
	if err != nil {
		return nil, fmt.Errorf("failed to generate clustering: %w", err)
	}

	members, ok := clustering.ByIndex(index)
	if !ok {
		return nil, fmt.Errorf("no cluster with index %d: %w", index, protocol.ErrClusterNotFound)
	}

	qcs := es.commitEvent.ClusterQCs
	if uint(len(qcs)) <= index {
		return nil, fmt.Errorf("internal data inconsistency: cannot get qc at index %d - epoch has %d clusters and %d cluster QCs",
			index, len(clustering), len(qcs))
	}
	rootQCVoteData := qcs[index]

	signerIndices, err := signature.EncodeSignersToIndices(members.NodeIDs(), rootQCVoteData.VoterIDs)
	if err != nil {
		return nil, fmt.Errorf("could not encode signer indices for rootQCVoteData.VoterIDs: %w", err)
	}

	rootBlock := cluster.CanonicalRootBlock(epochCounter, members)
	rootQC := &flow.QuorumCertificate{
		View:          rootBlock.Header.View,
		BlockID:       rootBlock.ID(),
		SignerIndices: signerIndices,
		SigData:       rootQCVoteData.SigData,
	}

	cluster, err := ClusterFromEncodable(EncodableCluster{
		Index:     index,
		Counter:   epochCounter,
		Members:   members,
		RootBlock: rootBlock,
		RootQC:    rootQC,
	})
	return cluster, err
}

func (es *committedEpoch) ClusterByChainID(chainID flow.ChainID) (protocol.Cluster, error) {
	clustering, err := es.Clustering()
	if err != nil {
		return nil, fmt.Errorf("failed to generate clustering: %w", err)
	}

	for i, cl := range clustering {
		if cluster.CanonicalClusterID(es.setupEvent.Counter, cl.NodeIDs()) == chainID {
			cl, err := es.Cluster(uint(i))
			if err != nil {
				return nil, fmt.Errorf("could not retrieve known existing cluster (idx=%d, id=%s): %v", i, chainID, err)
			}
			return cl, nil
		}
	}
	return nil, protocol.ErrClusterNotFound
}

func (es *committedEpoch) DKG() (protocol.DKG, error) {
	encodable, err := EncodableDKGFromEvents(es.setupEvent, es.commitEvent)
	if err != nil {
		return nil, fmt.Errorf("could not build encodable DKG from epoch events")
	}
	return DKGFromEncodable(encodable)
}

// heightBoundedEpoch represents an epoch (with counter N) for which we know either
// its start boundary, end boundary, or both. A boundary is included when:
//   - it occurred after this node's lowest known block AND
//   - it occurred before the latest finalized block (ie. the boundary is defined)
//
// heightBoundedEpoch has all the information of a committedEpoch, plus one or
// both height boundaries for the epoch.
type heightBoundedEpoch struct {
	committedEpoch
	firstHeight *uint64
	finalHeight *uint64
}

var _ protocol.CommittedEpoch = (*heightBoundedEpoch)(nil)

func (e *heightBoundedEpoch) FirstHeight() (uint64, error) {
	if e.firstHeight != nil {
		return *e.firstHeight, nil
	}
	return 0, protocol.ErrUnknownEpochBoundary
}

func (e *heightBoundedEpoch) FinalHeight() (uint64, error) {
	if e.finalHeight != nil {
		return *e.finalHeight, nil
	}
	return 0, protocol.ErrUnknownEpochBoundary
}

// NewSetupEpoch returns a memory-backed epoch implementation based on an EpochSetup event.
// Epoch information available after the setup phase will not be accessible in the resulting epoch instance.
// No errors are expected during normal operations.
func NewSetupEpoch(setupEvent *flow.EpochSetup, extensions []flow.EpochExtension) protocol.TentativeEpoch {
	return &setupEpoch{
		setupEvent: setupEvent,
		extensions: extensions,
	}
}

// NewCommittedEpoch returns a memory-backed epoch implementation based on an
// EpochSetup and EpochCommit events.
// No errors are expected during normal operations.
func NewCommittedEpoch(setupEvent *flow.EpochSetup, extensions []flow.EpochExtension, commitEvent *flow.EpochCommit) protocol.CommittedEpoch {
	return &committedEpoch{
		setupEpoch: setupEpoch{
			setupEvent: setupEvent,
			extensions: extensions,
		},
		commitEvent: commitEvent,
	}
}

// NewEpochWithStartBoundary returns a memory-backed epoch implementation based on an
// EpochSetup and EpochCommit events, and the epoch's first block height (start boundary).
// No errors are expected during normal operations.
func NewEpochWithStartBoundary(setupEvent *flow.EpochSetup, extensions []flow.EpochExtension, commitEvent *flow.EpochCommit, firstHeight uint64) protocol.CommittedEpoch {
	return &heightBoundedEpoch{
		committedEpoch: committedEpoch{
			setupEpoch: setupEpoch{
				setupEvent: setupEvent,
				extensions: extensions,
			},
			commitEvent: commitEvent,
		},
		firstHeight: &firstHeight,
		finalHeight: nil,
	}
}

// NewEpochWithEndBoundary returns a memory-backed epoch implementation based on an
// EpochSetup and EpochCommit events, and the epoch's final block height (end boundary).
// No errors are expected during normal operations.
func NewEpochWithEndBoundary(setupEvent *flow.EpochSetup, extensions []flow.EpochExtension, commitEvent *flow.EpochCommit, finalHeight uint64) protocol.CommittedEpoch {
	return &heightBoundedEpoch{
		committedEpoch: committedEpoch{
			setupEpoch: setupEpoch{
				setupEvent: setupEvent,
				extensions: extensions,
			},
			commitEvent: commitEvent,
		},
		firstHeight: nil,
		finalHeight: &finalHeight,
	}
}

// NewEpochWithStartAndEndBoundaries returns a memory-backed epoch implementation based on an
// EpochSetup and EpochCommit events, and the epoch's first and final block heights (start+end boundaries).
// No errors are expected during normal operations.
func NewEpochWithStartAndEndBoundaries(setupEvent *flow.EpochSetup, extensions []flow.EpochExtension, commitEvent *flow.EpochCommit, firstHeight, finalHeight uint64) protocol.CommittedEpoch {
	return &heightBoundedEpoch{
		committedEpoch: committedEpoch{
			setupEpoch: setupEpoch{
				setupEvent: setupEvent,
				extensions: extensions,
			},
			commitEvent: commitEvent,
		},
		firstHeight: &firstHeight,
		finalHeight: &finalHeight,
	}
}
