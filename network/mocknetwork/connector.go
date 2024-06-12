// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocknetwork

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// Connector is an autogenerated mock type for the Connector type
type Connector struct {
	mock.Mock
}

// Connect provides a mock function with given fields: ctx, peerChan
func (_m *Connector) Connect(ctx context.Context, peerChan <-chan peer.AddrInfo) {
	_m.Called(ctx, peerChan)
}

// NewConnector creates a new instance of Connector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnector(t interface {
	mock.TestingT
	Cleanup(func())
}) *Connector {
	mock := &Connector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
