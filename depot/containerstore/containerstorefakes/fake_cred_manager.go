// Code generated by counterfeiter. DO NOT EDIT.
package containerstorefakes

import (
	"sync"

	"github.com/cceasy/executor"
	"github.com/cceasy/executor/depot/containerstore"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/ifrit"
)

type FakeCredManager struct {
	CreateCredDirStub        func(lager.Logger, executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error)
	createCredDirMutex       sync.RWMutex
	createCredDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	createCredDirReturns struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}
	createCredDirReturnsOnCall map[int]struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}
	RunnerStub        func(lager.Logger, executor.Container) ifrit.Runner
	runnerMutex       sync.RWMutex
	runnerArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	runnerReturns struct {
		result1 ifrit.Runner
	}
	runnerReturnsOnCall map[int]struct {
		result1 ifrit.Runner
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredManager) CreateCredDir(arg1 lager.Logger, arg2 executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error) {
	fake.createCredDirMutex.Lock()
	ret, specificReturn := fake.createCredDirReturnsOnCall[len(fake.createCredDirArgsForCall)]
	fake.createCredDirArgsForCall = append(fake.createCredDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("CreateCredDir", []interface{}{arg1, arg2})
	fake.createCredDirMutex.Unlock()
	if fake.CreateCredDirStub != nil {
		return fake.CreateCredDirStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createCredDirReturns.result1, fake.createCredDirReturns.result2, fake.createCredDirReturns.result3
}

func (fake *FakeCredManager) CreateCredDirCallCount() int {
	fake.createCredDirMutex.RLock()
	defer fake.createCredDirMutex.RUnlock()
	return len(fake.createCredDirArgsForCall)
}

func (fake *FakeCredManager) CreateCredDirArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.createCredDirMutex.RLock()
	defer fake.createCredDirMutex.RUnlock()
	return fake.createCredDirArgsForCall[i].arg1, fake.createCredDirArgsForCall[i].arg2
}

func (fake *FakeCredManager) CreateCredDirReturns(result1 []garden.BindMount, result2 []executor.EnvironmentVariable, result3 error) {
	fake.CreateCredDirStub = nil
	fake.createCredDirReturns = struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCredManager) CreateCredDirReturnsOnCall(i int, result1 []garden.BindMount, result2 []executor.EnvironmentVariable, result3 error) {
	fake.CreateCredDirStub = nil
	if fake.createCredDirReturnsOnCall == nil {
		fake.createCredDirReturnsOnCall = make(map[int]struct {
			result1 []garden.BindMount
			result2 []executor.EnvironmentVariable
			result3 error
		})
	}
	fake.createCredDirReturnsOnCall[i] = struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCredManager) Runner(arg1 lager.Logger, arg2 executor.Container) ifrit.Runner {
	fake.runnerMutex.Lock()
	ret, specificReturn := fake.runnerReturnsOnCall[len(fake.runnerArgsForCall)]
	fake.runnerArgsForCall = append(fake.runnerArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("Runner", []interface{}{arg1, arg2})
	fake.runnerMutex.Unlock()
	if fake.RunnerStub != nil {
		return fake.RunnerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runnerReturns.result1
}

func (fake *FakeCredManager) RunnerCallCount() int {
	fake.runnerMutex.RLock()
	defer fake.runnerMutex.RUnlock()
	return len(fake.runnerArgsForCall)
}

func (fake *FakeCredManager) RunnerArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.runnerMutex.RLock()
	defer fake.runnerMutex.RUnlock()
	return fake.runnerArgsForCall[i].arg1, fake.runnerArgsForCall[i].arg2
}

func (fake *FakeCredManager) RunnerReturns(result1 ifrit.Runner) {
	fake.RunnerStub = nil
	fake.runnerReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeCredManager) RunnerReturnsOnCall(i int, result1 ifrit.Runner) {
	fake.RunnerStub = nil
	if fake.runnerReturnsOnCall == nil {
		fake.runnerReturnsOnCall = make(map[int]struct {
			result1 ifrit.Runner
		})
	}
	fake.runnerReturnsOnCall[i] = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeCredManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCredDirMutex.RLock()
	defer fake.createCredDirMutex.RUnlock()
	fake.runnerMutex.RLock()
	defer fake.runnerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredManager) recordInvocation(key string, args []interface{}) {
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

var _ containerstore.CredManager = new(FakeCredManager)
