// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cceasy/executor/depot/steps"
)

type FakeStep struct {
	PerformStub        func() error
	performMutex       sync.RWMutex
	performArgsForCall []struct{}
	performReturns     struct {
		result1 error
	}
	CancelStub        func()
	cancelMutex       sync.RWMutex
	cancelArgsForCall []struct{}
}

func (fake *FakeStep) Perform() error {
	fake.performMutex.Lock()
	fake.performArgsForCall = append(fake.performArgsForCall, struct{}{})
	fake.performMutex.Unlock()
	if fake.PerformStub != nil {
		return fake.PerformStub()
	} else {
		return fake.performReturns.result1
	}
}

func (fake *FakeStep) PerformCallCount() int {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return len(fake.performArgsForCall)
}

func (fake *FakeStep) PerformReturns(result1 error) {
	fake.PerformStub = nil
	fake.performReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStep) Cancel() {
	fake.cancelMutex.Lock()
	fake.cancelArgsForCall = append(fake.cancelArgsForCall, struct{}{})
	fake.cancelMutex.Unlock()
	if fake.CancelStub != nil {
		fake.CancelStub()
	}
}

func (fake *FakeStep) CancelCallCount() int {
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	return len(fake.cancelArgsForCall)
}

var _ steps.Step = new(FakeStep)
