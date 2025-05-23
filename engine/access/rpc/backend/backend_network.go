package backend

import (
	"context"
	"errors"

	accessmodel "github.com/onflow/flow-go/model/access"
	"github.com/onflow/flow-go/state"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow-go/engine/common/rpc/convert"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/state/protocol"
	"github.com/onflow/flow-go/storage"
)

type backendNetwork struct {
	state                protocol.State
	chainID              flow.ChainID
	headers              storage.Headers
	snapshotHistoryLimit int
}

/*
NetworkAPI func

The observer and access nodes need to be able to handle GetNetworkParameters
and GetLatestProtocolStateSnapshot RPCs so this logic was split into
the backendNetwork so that we can ignore the rest of the backend logic
*/
func NewNetworkAPI(
	state protocol.State,
	chainID flow.ChainID,
	headers storage.Headers,
	snapshotHistoryLimit int,
) *backendNetwork {
	return &backendNetwork{
		state:                state,
		chainID:              chainID,
		headers:              headers,
		snapshotHistoryLimit: snapshotHistoryLimit,
	}
}

func (b *backendNetwork) GetNetworkParameters(_ context.Context) accessmodel.NetworkParameters {
	return accessmodel.NetworkParameters{
		ChainID: b.chainID,
	}
}

// GetLatestProtocolStateSnapshot returns the latest finalized snapshot.
func (b *backendNetwork) GetLatestProtocolStateSnapshot(_ context.Context) ([]byte, error) {
	snapshot := b.state.Final()
	data, err := convert.SnapshotToBytes(snapshot)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert snapshot to bytes: %v", err)
	}

	return data, nil
}

// GetProtocolStateSnapshotByBlockID returns serializable Snapshot for a block, by blockID.
// The requested block must be finalized, otherwise an error is returned.
// Expected errors during normal operation:
//   - status.Error[codes.NotFound] - No block with the given ID was found
//   - status.Error[codes.InvalidArgument] - Block ID is for an orphaned block and will never have a valid snapshot
//   - status.Error[codes.FailedPrecondition] - A block was found, but it is not finalized and is above the finalized height.
//     The block may or may not be finalized in the future; the client can retry later.
func (b *backendNetwork) GetProtocolStateSnapshotByBlockID(_ context.Context, blockID flow.Identifier) ([]byte, error) {
	snapshot := b.state.AtBlockID(blockID)
	snapshotHeadByBlockId, err := snapshot.Head()
	if err != nil {
		if errors.Is(err, state.ErrUnknownSnapshotReference) {
			return nil, status.Errorf(codes.NotFound, "failed to get a valid snapshot: block not found")
		}
		return nil, status.Errorf(codes.Internal, "could not get header by blockID: %v", err)
	}

	// Because there is no index from block ID to finalized height, we separately look up the finalized
	// block ID by the height of the queried block, then compare the queried ID to the finalized ID.
	// If they match, then the queried block must be finalized.
	blockIDFinalizedAtHeight, err := b.headers.BlockIDByHeight(snapshotHeadByBlockId.Height)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			// The block exists, but no block has been finalized at its height. Therefore, this block
			// may be finalized in the future, and the client can retry.
			return nil, status.Errorf(codes.FailedPrecondition,
				"failed to retrieve snapshot for block with height %d: block not finalized and is above finalized height",
				snapshotHeadByBlockId.Height)
		}
		return nil, status.Errorf(codes.Internal, "failed to lookup block id by height %d", snapshotHeadByBlockId.Height)
	}

	if blockIDFinalizedAtHeight != blockID {
		// A different block than what was queried has been finalized at this height.
		// Therefore, the queried block will never be finalized.
		return nil, status.Errorf(codes.InvalidArgument,
			"failed to retrieve snapshot for block: block not finalized and is below finalized height")
	}

	data, err := convert.SnapshotToBytes(snapshot)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert snapshot to bytes: %v", err)
	}
	return data, nil
}

// GetProtocolStateSnapshotByHeight returns serializable Snapshot by block height.
// The block must be finalized (otherwise the by-height query is ambiguous).
// Expected errors during normal operation:
//   - status.Error[codes.NotFound] - No block with the given height was found.
//     The block height may or may not be finalized in the future; the client can retry later.
func (b *backendNetwork) GetProtocolStateSnapshotByHeight(_ context.Context, blockHeight uint64) ([]byte, error) {
	snapshot := b.state.AtHeight(blockHeight)
	_, err := snapshot.Head()
	if err != nil {
		if errors.Is(err, state.ErrUnknownSnapshotReference) {
			return nil, status.Errorf(codes.NotFound, "failed to find snapshot: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get a valid snapshot: %v", err)
	}

	data, err := convert.SnapshotToBytes(snapshot)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert snapshot to bytes: %v", err)
	}
	return data, nil
}
