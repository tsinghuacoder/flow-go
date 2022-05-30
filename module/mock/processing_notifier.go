// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// ProcessingNotifier is an autogenerated mock type for the ProcessingNotifier type
type ProcessingNotifier struct {
	mock.Mock
}

// Notify provides a mock function with given fields: entityID
func (_m *ProcessingNotifier) Notify(entityID flow.Identifier) {
	_m.Called(entityID)
}

type NewProcessingNotifierT interface {
	mock.TestingT
	Cleanup(func())
}

// NewProcessingNotifier creates a new instance of ProcessingNotifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProcessingNotifier(t NewProcessingNotifierT) *ProcessingNotifier {
	mock := &ProcessingNotifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
