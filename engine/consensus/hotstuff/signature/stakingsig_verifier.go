// +build relic

package signature

import (
	"fmt"

	"github.com/dapperlabs/flow-go/crypto"
	model "github.com/dapperlabs/flow-go/model/hotstuff"
)

// StakingSigVerifier verifies signatures generated with staking keys. Specifically, it verifies
// individual signaturesz (e.g. from a vote) and aggregated signatures (e.g. from a Quorum Certificate).
type StakingSigVerifier struct {
	stakingHasher crypto.Hasher // the hasher for staking signature
}

// NewStakingSigVerifier constructs a new StakingSigVerifier
// The tag used for identifying the vote is different between collector and consensus nodes.
func NewStakingSigVerifier(stakingSigTag string) StakingSigVerifier {
	return StakingSigVerifier{
		stakingHasher: crypto.NewBLS_KMAC(stakingSigTag),
	}
}

// VerifyStakingSig verifies a single BLS staking signature for a block using signer's public key
// sig - the signature to be verified
// block - the block that the signature was signed for.
// signerKey - the public key of the signer who signed the block.
//
// Note: we are specifically choosing safety over performance here.
//   * The vote itself contains all the information for verifying the signature: the blockID and the block's view
//   * We could use the vote to verify that the signature is valid for the information contained in the vote's message
//   * However, for security, we are explicitly verifying that the vote matches the full block.
//     We do this by converting the block to the byte-sequence which we expect an honest voter to have signed
//     and then check the provided signature against this self-computed byte-sequence.
func (s *StakingSigVerifier) VerifyStakingSig(sig crypto.Signature, block *model.Block, signerKey crypto.PublicKey) (bool, error) {
	msg := BlockToBytesForSign(block)
	valid, err := signerKey.Verify(sig, msg, s.stakingHasher)
	if err != nil {
		return false, fmt.Errorf("cannot verify staking sig: %w", err)
	}
	return valid, nil
}

// VerifyAggregatedStakingSignature verifies an aggregated BLS signature.
// Inputs:
//    aggStakingSig - the aggregated staking signature to be verified
//    block - the block that the signature was signed for.
//    signerKeys - the signer's public staking key
//
// Note: we are specifically choosing safety over performance here.
//   * The vote itself contains all the information for verifying the signature: the blockID and the block's view
//   * We could use the vote to verify that the signature is valid for the information contained in the vote's message
//   * However, for security, we are explicitly verifying that the vote matches the full block.
//     We do this by converting the block to the byte-sequence which we expect an honest voter to have signed
//     and then check the provided signature against this self-computed byte-sequence.
//
// For now, the aggregated BLS staking signature is implemented as a slice of individual signatures.
// To verify it, we just verify every single signature. The implementation (and method signature)
// will later be updated, once full BLS sigmnature aggregation is implemented.
func (s *StakingSigVerifier) VerifyAggregatedStakingSignature(aggStakingSig []crypto.Signature, block *model.Block, signerKeys []crypto.PublicKey) (bool, error) {
	// check that the number of keys and signatures should match
	if len(aggStakingSig) != len(signerKeys) {
		return false, nil
	}

	msg := BlockToBytesForSign(block)

	// check each signature
	for i, sig := range aggStakingSig {
		signerKey := signerKeys[i]

		// validate the staking signature
		valid, err := signerKey.Verify(sig, msg, s.stakingHasher)
		if err != nil {
			return false, fmt.Errorf("cannot verify aggregated staking sig for (%d)-th sig: %w", i, err)
		}
		if !valid {
			return false, nil
		}
	}
	return true, nil
}
