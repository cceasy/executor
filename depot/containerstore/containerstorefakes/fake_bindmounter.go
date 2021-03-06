// Code generated by counterfeiter. DO NOT EDIT.
package containerstorefakes

import (
	"sync"

	"github.com/cceasy/executor"
	"github.com/cceasy/executor/depot/containerstore"
	"github.com/cceasy/executor/depot/log_streamer"
	"code.cloudfoundry.org/lager"
)

type FakeDependencyManager struct {
	DownloadCachedDependenciesStub        func(logger lager.Logger, mounts []executor.CachedDependency, logStreamer log_streamer.LogStreamer) (containerstore.BindMounts, error)
	downloadCachedDependenciesMutex       sync.RWMutex
	downloadCachedDependenciesArgsForCall []struct {
		logger      lager.Logger
		mounts      []executor.CachedDependency
		logStreamer log_streamer.LogStreamer
	}
	downloadCachedDependenciesReturns struct {
		result1 containerstore.BindMounts
		result2 error
	}
	downloadCachedDependenciesReturnsOnCall map[int]struct {
		result1 containerstore.BindMounts
		result2 error
	}
	ReleaseCachedDependenciesStub        func(logger lager.Logger, keys []containerstore.BindMountCacheKey) error
	releaseCachedDependenciesMutex       sync.RWMutex
	releaseCachedDependenciesArgsForCall []struct {
		logger lager.Logger
		keys   []containerstore.BindMountCacheKey
	}
	releaseCachedDependenciesReturns struct {
		result1 error
	}
	releaseCachedDependenciesReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func(logger lager.Logger)
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		logger lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDependencyManager) DownloadCachedDependencies(logger lager.Logger, mounts []executor.CachedDependency, logStreamer log_streamer.LogStreamer) (containerstore.BindMounts, error) {
	var mountsCopy []executor.CachedDependency
	if mounts != nil {
		mountsCopy = make([]executor.CachedDependency, len(mounts))
		copy(mountsCopy, mounts)
	}
	fake.downloadCachedDependenciesMutex.Lock()
	ret, specificReturn := fake.downloadCachedDependenciesReturnsOnCall[len(fake.downloadCachedDependenciesArgsForCall)]
	fake.downloadCachedDependenciesArgsForCall = append(fake.downloadCachedDependenciesArgsForCall, struct {
		logger      lager.Logger
		mounts      []executor.CachedDependency
		logStreamer log_streamer.LogStreamer
	}{logger, mountsCopy, logStreamer})
	fake.recordInvocation("DownloadCachedDependencies", []interface{}{logger, mountsCopy, logStreamer})
	fake.downloadCachedDependenciesMutex.Unlock()
	if fake.DownloadCachedDependenciesStub != nil {
		return fake.DownloadCachedDependenciesStub(logger, mounts, logStreamer)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.downloadCachedDependenciesReturns.result1, fake.downloadCachedDependenciesReturns.result2
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesCallCount() int {
	fake.downloadCachedDependenciesMutex.RLock()
	defer fake.downloadCachedDependenciesMutex.RUnlock()
	return len(fake.downloadCachedDependenciesArgsForCall)
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesArgsForCall(i int) (lager.Logger, []executor.CachedDependency, log_streamer.LogStreamer) {
	fake.downloadCachedDependenciesMutex.RLock()
	defer fake.downloadCachedDependenciesMutex.RUnlock()
	return fake.downloadCachedDependenciesArgsForCall[i].logger, fake.downloadCachedDependenciesArgsForCall[i].mounts, fake.downloadCachedDependenciesArgsForCall[i].logStreamer
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesReturns(result1 containerstore.BindMounts, result2 error) {
	fake.DownloadCachedDependenciesStub = nil
	fake.downloadCachedDependenciesReturns = struct {
		result1 containerstore.BindMounts
		result2 error
	}{result1, result2}
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesReturnsOnCall(i int, result1 containerstore.BindMounts, result2 error) {
	fake.DownloadCachedDependenciesStub = nil
	if fake.downloadCachedDependenciesReturnsOnCall == nil {
		fake.downloadCachedDependenciesReturnsOnCall = make(map[int]struct {
			result1 containerstore.BindMounts
			result2 error
		})
	}
	fake.downloadCachedDependenciesReturnsOnCall[i] = struct {
		result1 containerstore.BindMounts
		result2 error
	}{result1, result2}
}

func (fake *FakeDependencyManager) ReleaseCachedDependencies(logger lager.Logger, keys []containerstore.BindMountCacheKey) error {
	var keysCopy []containerstore.BindMountCacheKey
	if keys != nil {
		keysCopy = make([]containerstore.BindMountCacheKey, len(keys))
		copy(keysCopy, keys)
	}
	fake.releaseCachedDependenciesMutex.Lock()
	ret, specificReturn := fake.releaseCachedDependenciesReturnsOnCall[len(fake.releaseCachedDependenciesArgsForCall)]
	fake.releaseCachedDependenciesArgsForCall = append(fake.releaseCachedDependenciesArgsForCall, struct {
		logger lager.Logger
		keys   []containerstore.BindMountCacheKey
	}{logger, keysCopy})
	fake.recordInvocation("ReleaseCachedDependencies", []interface{}{logger, keysCopy})
	fake.releaseCachedDependenciesMutex.Unlock()
	if fake.ReleaseCachedDependenciesStub != nil {
		return fake.ReleaseCachedDependenciesStub(logger, keys)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.releaseCachedDependenciesReturns.result1
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesCallCount() int {
	fake.releaseCachedDependenciesMutex.RLock()
	defer fake.releaseCachedDependenciesMutex.RUnlock()
	return len(fake.releaseCachedDependenciesArgsForCall)
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesArgsForCall(i int) (lager.Logger, []containerstore.BindMountCacheKey) {
	fake.releaseCachedDependenciesMutex.RLock()
	defer fake.releaseCachedDependenciesMutex.RUnlock()
	return fake.releaseCachedDependenciesArgsForCall[i].logger, fake.releaseCachedDependenciesArgsForCall[i].keys
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesReturns(result1 error) {
	fake.ReleaseCachedDependenciesStub = nil
	fake.releaseCachedDependenciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesReturnsOnCall(i int, result1 error) {
	fake.ReleaseCachedDependenciesStub = nil
	if fake.releaseCachedDependenciesReturnsOnCall == nil {
		fake.releaseCachedDependenciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.releaseCachedDependenciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencyManager) Stop(logger lager.Logger) {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("Stop", []interface{}{logger})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub(logger)
	}
}

func (fake *FakeDependencyManager) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeDependencyManager) StopArgsForCall(i int) lager.Logger {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].logger
}

func (fake *FakeDependencyManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadCachedDependenciesMutex.RLock()
	defer fake.downloadCachedDependenciesMutex.RUnlock()
	fake.releaseCachedDependenciesMutex.RLock()
	defer fake.releaseCachedDependenciesMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDependencyManager) recordInvocation(key string, args []interface{}) {
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

var _ containerstore.DependencyManager = new(FakeDependencyManager)
