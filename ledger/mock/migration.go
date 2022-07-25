// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import (
	ledger "github.com/onflow/flow-go/ledger"
	mock "github.com/stretchr/testify/mock"
)

// Migration is an autogenerated mock type for the Migration type
type Migration struct {
	mock.Mock
}

// Execute provides a mock function with given fields: payloads
func (_m *Migration) Execute(payloads []ledger.Payload) ([]ledger.Payload, error) {
	ret := _m.Called(payloads)

	var r0 []ledger.Payload
	if rf, ok := ret.Get(0).(func([]ledger.Payload) []ledger.Payload); ok {
		r0 = rf(payloads)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ledger.Payload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]ledger.Payload) error); ok {
		r1 = rf(payloads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewMigrationT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMigration creates a new instance of Migration. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMigration(t NewMigrationT) *Migration {
	mock := &Migration{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
