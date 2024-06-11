// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// BlockFinalized provides a mock function with given fields: block
func (_m *Consumer) BlockFinalized(block *flow.Header) {
	_m.Called(block)
}

// BlockProcessable provides a mock function with given fields: block, certifyingQC
func (_m *Consumer) BlockProcessable(block *flow.Header, certifyingQC *flow.QuorumCertificate) {
	_m.Called(block, certifyingQC)
}

// EpochCommittedPhaseStarted provides a mock function with given fields: currentEpochCounter, first
func (_m *Consumer) EpochCommittedPhaseStarted(currentEpochCounter uint64, first *flow.Header) {
	_m.Called(currentEpochCounter, first)
}

// EpochEmergencyFallbackTriggered provides a mock function with given fields:
func (_m *Consumer) EpochEmergencyFallbackTriggered() {
	_m.Called()
}

// EpochExtended provides a mock function with given fields: _a0
func (_m *Consumer) EpochExtended(_a0 flow.EpochExtension) {
	_m.Called(_a0)
}

// EpochSetupPhaseStarted provides a mock function with given fields: currentEpochCounter, first
func (_m *Consumer) EpochSetupPhaseStarted(currentEpochCounter uint64, first *flow.Header) {
	_m.Called(currentEpochCounter, first)
}

// EpochTransition provides a mock function with given fields: newEpochCounter, first
func (_m *Consumer) EpochTransition(newEpochCounter uint64, first *flow.Header) {
	_m.Called(newEpochCounter, first)
}

// NewConsumer creates a new instance of Consumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
