// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package trickle

import (
	"github.com/dapperlabs/flow-go/model"
)

// Peer holds the information we know about a peer on the network, such as his
// ID and what events he has seen.
type Peer struct {
	ID   model.Identifier
	Seen map[string]struct{}
}

// PeerList represents a list of peers.
type PeerList []*Peer

// IDs returns the list of IDs for the peers in the list.
func (pl PeerList) NodeIDs() []model.Identifier {
	nodeIDs := make([]model.Identifier, 0, len(pl))
	for _, p := range pl {
		nodeIDs = append(nodeIDs, p.ID)
	}
	return nodeIDs
}

// PeerFilter is a function to filter peers by the information they hold.
type PeerFilter func(*Peer) bool
