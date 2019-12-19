// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package logging

import (
	"encoding/hex"

	"github.com/dapperlabs/flow-go/model"
)

func HexSlice(nodeIDs []model.Identifier) []string {
	ss := make([]string, 0, len(nodeIDs))
	for _, nodeID := range nodeIDs {
		ss = append(ss, hex.EncodeToString(nodeID[:]))
	}
	return ss
}
