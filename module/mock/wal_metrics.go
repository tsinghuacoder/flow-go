// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// WALMetrics is an autogenerated mock type for the WALMetrics type
type WALMetrics struct {
	mock.Mock
}

// DiskSize provides a mock function with given fields: _a0
func (_m *WALMetrics) DiskSize(_a0 uint64) {
	_m.Called(_a0)
}

type NewWALMetricsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewWALMetrics creates a new instance of WALMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWALMetrics(t NewWALMetricsT) *WALMetrics {
	mock := &WALMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
