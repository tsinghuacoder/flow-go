package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/onflow/flow-go/cmd/bootstrap/run"
	"github.com/onflow/flow-go/cmd/util/cmd/common"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/bootstrap"
	"github.com/onflow/flow-go/model/dkg"
	"github.com/onflow/flow-go/model/flow"
)

// constructRootQC constructs root QC based on root block, votes and dkg info
func constructRootQC(block *flow.Block, votes []*model.Vote, allNodes, internalNodes []bootstrap.NodeInfo, randomBeaconData dkg.ThresholdKeySet) *flow.QuorumCertificate {

	identities := bootstrap.ToIdentityList(allNodes)
	participantData, err := run.GenerateQCParticipantData(allNodes, internalNodes, randomBeaconData)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate QC participant data")
	}

	qc, invalidVotesErr, err := run.GenerateRootQC(block, votes, participantData, identities)
	if len(invalidVotesErr) > 0 {
		for _, err := range invalidVotesErr {
			log.Warn().Err(err).Msg("invalid vote")
		}
	}

	if err != nil {
		log.Fatal().Err(err).Msg("generating root QC failed")
	}

	return qc
}

// NOTE: allNodes must be in the same order as when generating the DKG
func constructRootVotes(block *flow.Block, allNodes, internalNodes []bootstrap.NodeInfo, randomBeaconData dkg.ThresholdKeySet) {
	participantData, err := run.GenerateQCParticipantData(allNodes, internalNodes, randomBeaconData)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate QC participant data")
	}

	votes, err := run.GenerateRootBlockVotes(block, participantData)
	if err != nil {
		log.Fatal().Err(err).Msg("generating votes for root block failed")
	}

	for _, vote := range votes {
		path := filepath.Join(bootstrap.DirnameRootBlockVotes, fmt.Sprintf(bootstrap.FilenameRootBlockVote, vote.SignerID))
		err = common.WriteJSON(path, flagOutdir, vote)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to write json")
		}
		log.Info().Msgf("wrote file %s/%s", flagOutdir, path)
	}
}
