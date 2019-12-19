// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	"github.com/dapperlabs/flow-go/model"
	flow "github.com/dapperlabs/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Identities is an autogenerated mock type for the Identities type
type Identities struct {
	mock.Mock
}

// ByNodeID provides a mock function with given fields: _a0
func (_m *Identities) ByNodeID(_a0 model.Identifier) (flow.Identity, error) {
	ret := _m.Called(_a0)

	var r0 flow.Identity
	if rf, ok := ret.Get(0).(func(model.Identifier) flow.Identity); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(flow.Identity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Identifier) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
