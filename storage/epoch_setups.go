// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package storage

import (
	"github.com/dgraph-io/badger/v2"

	"github.com/dapperlabs/flow-go/model/flow"
)

type EpochSetups interface {

	// StoreTx allows us to store a new epoch setup in a DB transaction while going through the cache.
	StoreTx(*flow.EpochSetup) func(*badger.Txn) error

	// BySetupID will return the EpochSetup event by its ID.
	BySetupID(flow.Identifier) (*flow.EpochSetup, error)
}
