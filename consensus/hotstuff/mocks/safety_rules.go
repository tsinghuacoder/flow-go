// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"
)

// SafetyRules is an autogenerated mock type for the SafetyRules type
type SafetyRules struct {
	mock.Mock
}

// ProduceTimeout provides a mock function with given fields: curView, highestQC, highestTC
func (_m *SafetyRules) ProduceTimeout(curView uint64, highestQC *flow.QuorumCertificate, highestTC *flow.TimeoutCertificate) (*model.TimeoutObject, error) {
	ret := _m.Called(curView, highestQC, highestTC)

	var r0 *model.TimeoutObject
	if rf, ok := ret.Get(0).(func(uint64, *flow.QuorumCertificate, *flow.TimeoutCertificate) *model.TimeoutObject); ok {
		r0 = rf(curView, highestQC, highestTC)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TimeoutObject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, *flow.QuorumCertificate, *flow.TimeoutCertificate) error); ok {
		r1 = rf(curView, highestQC, highestTC)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProduceVote provides a mock function with given fields: proposal, curView
func (_m *SafetyRules) ProduceVote(proposal *model.Proposal, curView uint64) (*model.Vote, error) {
	ret := _m.Called(proposal, curView)

	var r0 *model.Vote
	if rf, ok := ret.Get(0).(func(*model.Proposal, uint64) *model.Vote); ok {
		r0 = rf(proposal, curView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Vote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Proposal, uint64) error); ok {
		r1 = rf(proposal, curView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
