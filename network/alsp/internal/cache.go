package internal

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/mempool/stdmap"
	"github.com/onflow/flow-go/network/alsp"
)

// SpamRecordCache is a cache that stores spam records.
type SpamRecordCache struct {
	recordFactory func(flow.Identifier) alsp.ProtocolSpamRecord // recordFactory is a factory function that creates a new spam record.
	c             *stdmap.Backend                               // c is the underlying cache.
}

var _ alsp.SpamRecordCache = (*SpamRecordCache)(nil)

func NewSpamRecordCache(recordFactory func(flow.Identifier) alsp.ProtocolSpamRecord) *SpamRecordCache {
	return &SpamRecordCache{
		recordFactory: recordFactory,
		c:             stdmap.NewBackend(),
	}
}

// Init initializes the spam record cache for the given origin id if it does not exist.
// Returns true if the record is initialized, false otherwise (i.e., the record already exists).
// Args:
// - originId: the origin id of the spam record.
// Returns:
// - true if the record is initialized, false otherwise (i.e., the record already exists).
func (s *SpamRecordCache) Init(originId flow.Identifier) bool {
	return s.c.Add(ProtocolSpamRecordEntity{s.recordFactory(originId)})
}

// Adjust applies the given adjust function to the spam record of the given origin id.
// Returns the Penalty value of the record after the adjustment.
// It returns an error if the adjustFunc returns an error or if the record does not exist.
// Assuming that adjust is always called when the record exists, the error is irrecoverable and indicates a bug.
// Args:
// - originId: the origin id of the spam record.
// - adjustFunc: the function that adjusts the spam record.
// Returns:
// - Penalty value of the record after the adjustment.
func (s *SpamRecordCache) Adjust(originId flow.Identifier, adjustFunc alsp.RecordAdjustFunc) (float64, error) {
	var rErr error
	adjustedEntity, adjusted := s.c.Adjust(originId, func(entity flow.Entity) flow.Entity {
		record, ok := entity.(ProtocolSpamRecordEntity)
		if !ok {
			// sanity check
			// This should never happen, because the cache only contains ProtocolSpamRecordEntity entities.
			panic("invalid entity type, expected ProtocolSpamRecordEntity type")
		}

		// Adjust the record.
		adjustedRecord, err := adjustFunc(record.ProtocolSpamRecord)
		if err != nil {
			rErr = fmt.Errorf("adjust function failed: %w", err)
			return entity // returns the original entity (reverse the adjustment).
		}

		// Return the adjusted record.
		return ProtocolSpamRecordEntity{adjustedRecord}
	})

	if rErr != nil {
		return 0, fmt.Errorf("failed to adjust record: %w", rErr)
	}

	if !adjusted {
		return 0, fmt.Errorf("record does not exist")
	}

	return adjustedEntity.(ProtocolSpamRecordEntity).Penalty, nil
}

// Identities returns the list of identities of the nodes that have a spam record in the cache.
func (s *SpamRecordCache) Identities() []flow.Identifier {
	return flow.GetIDs(s.c.All())
}

// Remove removes the spam record of the given origin id from the cache.
// Returns true if the record is removed, false otherwise (i.e., the record does not exist).
// Args:
// - originId: the origin id of the spam record.
// Returns:
// - true if the record is removed, false otherwise (i.e., the record does not exist).
func (s *SpamRecordCache) Remove(originId flow.Identifier) bool {
	return s.c.Remove(originId)
}
