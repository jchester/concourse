// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	sync "sync"

	atc "github.com/concourse/concourse/atc"
	v2 "github.com/concourse/concourse/atc/resource/v2"
)

type FakeCheckEventHandler struct {
	FinishStub        func() error
	finishMutex       sync.RWMutex
	finishArgsForCall []struct {
	}
	finishReturns struct {
		result1 error
	}
	finishReturnsOnCall map[int]struct {
		result1 error
	}
	SaveStub        func(atc.Space, atc.Version, atc.Metadata) error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		arg1 atc.Space
		arg2 atc.Version
		arg3 atc.Metadata
	}
	saveReturns struct {
		result1 error
	}
	saveReturnsOnCall map[int]struct {
		result1 error
	}
	SaveDefaultStub        func(atc.Space) error
	saveDefaultMutex       sync.RWMutex
	saveDefaultArgsForCall []struct {
		arg1 atc.Space
	}
	saveDefaultReturns struct {
		result1 error
	}
	saveDefaultReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckEventHandler) Finish() error {
	fake.finishMutex.Lock()
	ret, specificReturn := fake.finishReturnsOnCall[len(fake.finishArgsForCall)]
	fake.finishArgsForCall = append(fake.finishArgsForCall, struct {
	}{})
	fake.recordInvocation("Finish", []interface{}{})
	fake.finishMutex.Unlock()
	if fake.FinishStub != nil {
		return fake.FinishStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.finishReturns
	return fakeReturns.result1
}

func (fake *FakeCheckEventHandler) FinishCallCount() int {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return len(fake.finishArgsForCall)
}

func (fake *FakeCheckEventHandler) FinishCalls(stub func() error) {
	fake.finishMutex.Lock()
	defer fake.finishMutex.Unlock()
	fake.FinishStub = stub
}

func (fake *FakeCheckEventHandler) FinishReturns(result1 error) {
	fake.finishMutex.Lock()
	defer fake.finishMutex.Unlock()
	fake.FinishStub = nil
	fake.finishReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckEventHandler) FinishReturnsOnCall(i int, result1 error) {
	fake.finishMutex.Lock()
	defer fake.finishMutex.Unlock()
	fake.FinishStub = nil
	if fake.finishReturnsOnCall == nil {
		fake.finishReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.finishReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckEventHandler) Save(arg1 atc.Space, arg2 atc.Version, arg3 atc.Metadata) error {
	fake.saveMutex.Lock()
	ret, specificReturn := fake.saveReturnsOnCall[len(fake.saveArgsForCall)]
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		arg1 atc.Space
		arg2 atc.Version
		arg3 atc.Metadata
	}{arg1, arg2, arg3})
	fake.recordInvocation("Save", []interface{}{arg1, arg2, arg3})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveReturns
	return fakeReturns.result1
}

func (fake *FakeCheckEventHandler) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeCheckEventHandler) SaveCalls(stub func(atc.Space, atc.Version, atc.Metadata) error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = stub
}

func (fake *FakeCheckEventHandler) SaveArgsForCall(i int) (atc.Space, atc.Version, atc.Metadata) {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	argsForCall := fake.saveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCheckEventHandler) SaveReturns(result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckEventHandler) SaveReturnsOnCall(i int, result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	if fake.saveReturnsOnCall == nil {
		fake.saveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckEventHandler) SaveDefault(arg1 atc.Space) error {
	fake.saveDefaultMutex.Lock()
	ret, specificReturn := fake.saveDefaultReturnsOnCall[len(fake.saveDefaultArgsForCall)]
	fake.saveDefaultArgsForCall = append(fake.saveDefaultArgsForCall, struct {
		arg1 atc.Space
	}{arg1})
	fake.recordInvocation("SaveDefault", []interface{}{arg1})
	fake.saveDefaultMutex.Unlock()
	if fake.SaveDefaultStub != nil {
		return fake.SaveDefaultStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveDefaultReturns
	return fakeReturns.result1
}

func (fake *FakeCheckEventHandler) SaveDefaultCallCount() int {
	fake.saveDefaultMutex.RLock()
	defer fake.saveDefaultMutex.RUnlock()
	return len(fake.saveDefaultArgsForCall)
}

func (fake *FakeCheckEventHandler) SaveDefaultCalls(stub func(atc.Space) error) {
	fake.saveDefaultMutex.Lock()
	defer fake.saveDefaultMutex.Unlock()
	fake.SaveDefaultStub = stub
}

func (fake *FakeCheckEventHandler) SaveDefaultArgsForCall(i int) atc.Space {
	fake.saveDefaultMutex.RLock()
	defer fake.saveDefaultMutex.RUnlock()
	argsForCall := fake.saveDefaultArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckEventHandler) SaveDefaultReturns(result1 error) {
	fake.saveDefaultMutex.Lock()
	defer fake.saveDefaultMutex.Unlock()
	fake.SaveDefaultStub = nil
	fake.saveDefaultReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckEventHandler) SaveDefaultReturnsOnCall(i int, result1 error) {
	fake.saveDefaultMutex.Lock()
	defer fake.saveDefaultMutex.Unlock()
	fake.SaveDefaultStub = nil
	if fake.saveDefaultReturnsOnCall == nil {
		fake.saveDefaultReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveDefaultReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckEventHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	fake.saveDefaultMutex.RLock()
	defer fake.saveDefaultMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckEventHandler) recordInvocation(key string, args []interface{}) {
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

var _ v2.CheckEventHandler = new(FakeCheckEventHandler)
