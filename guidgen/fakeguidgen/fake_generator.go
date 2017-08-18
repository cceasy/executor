// Code generated by counterfeiter. DO NOT EDIT.
package fakeguidgen

import (
	"sync"

	"github.com/cceasy/executor/guidgen"
	"code.cloudfoundry.org/lager"
)

type FakeGenerator struct {
	GuidStub        func(lager.Logger) string
	guidMutex       sync.RWMutex
	guidArgsForCall []struct {
		arg1 lager.Logger
	}
	guidReturns struct {
		result1 string
	}
	guidReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGenerator) Guid(arg1 lager.Logger) string {
	fake.guidMutex.Lock()
	ret, specificReturn := fake.guidReturnsOnCall[len(fake.guidArgsForCall)]
	fake.guidArgsForCall = append(fake.guidArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Guid", []interface{}{arg1})
	fake.guidMutex.Unlock()
	if fake.GuidStub != nil {
		return fake.GuidStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.guidReturns.result1
}

func (fake *FakeGenerator) GuidCallCount() int {
	fake.guidMutex.RLock()
	defer fake.guidMutex.RUnlock()
	return len(fake.guidArgsForCall)
}

func (fake *FakeGenerator) GuidArgsForCall(i int) lager.Logger {
	fake.guidMutex.RLock()
	defer fake.guidMutex.RUnlock()
	return fake.guidArgsForCall[i].arg1
}

func (fake *FakeGenerator) GuidReturns(result1 string) {
	fake.GuidStub = nil
	fake.guidReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGenerator) GuidReturnsOnCall(i int, result1 string) {
	fake.GuidStub = nil
	if fake.guidReturnsOnCall == nil {
		fake.guidReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.guidReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.guidMutex.RLock()
	defer fake.guidMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGenerator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ guidgen.Generator = new(FakeGenerator)
