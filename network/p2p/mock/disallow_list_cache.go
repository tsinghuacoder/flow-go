// Code generated by mockery v2.43.2. DO NOT EDIT.

package mockp2p

import (
	network "github.com/onflow/flow-go/network"
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// DisallowListCache is an autogenerated mock type for the DisallowListCache type
type DisallowListCache struct {
	mock.Mock
}

// AllowFor provides a mock function with given fields: peerID, cause
func (_m *DisallowListCache) AllowFor(peerID peer.ID, cause network.DisallowListedCause) []network.DisallowListedCause {
	ret := _m.Called(peerID, cause)

	if len(ret) == 0 {
		panic("no return value specified for AllowFor")
	}

	var r0 []network.DisallowListedCause
	if rf, ok := ret.Get(0).(func(peer.ID, network.DisallowListedCause) []network.DisallowListedCause); ok {
		r0 = rf(peerID, cause)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]network.DisallowListedCause)
		}
	}

	return r0
}

// DisallowFor provides a mock function with given fields: peerID, cause
func (_m *DisallowListCache) DisallowFor(peerID peer.ID, cause network.DisallowListedCause) ([]network.DisallowListedCause, error) {
	ret := _m.Called(peerID, cause)

	if len(ret) == 0 {
		panic("no return value specified for DisallowFor")
	}

	var r0 []network.DisallowListedCause
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID, network.DisallowListedCause) ([]network.DisallowListedCause, error)); ok {
		return rf(peerID, cause)
	}
	if rf, ok := ret.Get(0).(func(peer.ID, network.DisallowListedCause) []network.DisallowListedCause); ok {
		r0 = rf(peerID, cause)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]network.DisallowListedCause)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID, network.DisallowListedCause) error); ok {
		r1 = rf(peerID, cause)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDisallowListed provides a mock function with given fields: peerID
func (_m *DisallowListCache) IsDisallowListed(peerID peer.ID) ([]network.DisallowListedCause, bool) {
	ret := _m.Called(peerID)

	if len(ret) == 0 {
		panic("no return value specified for IsDisallowListed")
	}

	var r0 []network.DisallowListedCause
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) ([]network.DisallowListedCause, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) []network.DisallowListedCause); ok {
		r0 = rf(peerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]network.DisallowListedCause)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewDisallowListCache creates a new instance of DisallowListCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisallowListCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisallowListCache {
	mock := &DisallowListCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
