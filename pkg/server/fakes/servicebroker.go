// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/pivotal-cf/brokerapi/v8/domain"
)

type FakeServiceBroker struct {
	BindStub        func(context.Context, string, string, domain.BindDetails, bool) (domain.Binding, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.BindDetails
		arg5 bool
	}
	bindReturns struct {
		result1 domain.Binding
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 domain.Binding
		result2 error
	}
	DeprovisionStub        func(context.Context, string, domain.DeprovisionDetails, bool) (domain.DeprovisionServiceSpec, error)
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.DeprovisionDetails
		arg4 bool
	}
	deprovisionReturns struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}
	deprovisionReturnsOnCall map[int]struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}
	GetBindingStub        func(context.Context, string, string, domain.FetchBindingDetails) (domain.GetBindingSpec, error)
	getBindingMutex       sync.RWMutex
	getBindingArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.FetchBindingDetails
	}
	getBindingReturns struct {
		result1 domain.GetBindingSpec
		result2 error
	}
	getBindingReturnsOnCall map[int]struct {
		result1 domain.GetBindingSpec
		result2 error
	}
	GetInstanceStub        func(context.Context, string, domain.FetchInstanceDetails) (domain.GetInstanceDetailsSpec, error)
	getInstanceMutex       sync.RWMutex
	getInstanceArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.FetchInstanceDetails
	}
	getInstanceReturns struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}
	getInstanceReturnsOnCall map[int]struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}
	LastBindingOperationStub        func(context.Context, string, string, domain.PollDetails) (domain.LastOperation, error)
	lastBindingOperationMutex       sync.RWMutex
	lastBindingOperationArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.PollDetails
	}
	lastBindingOperationReturns struct {
		result1 domain.LastOperation
		result2 error
	}
	lastBindingOperationReturnsOnCall map[int]struct {
		result1 domain.LastOperation
		result2 error
	}
	LastOperationStub        func(context.Context, string, domain.PollDetails) (domain.LastOperation, error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.PollDetails
	}
	lastOperationReturns struct {
		result1 domain.LastOperation
		result2 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 domain.LastOperation
		result2 error
	}
	ProvisionStub        func(context.Context, string, domain.ProvisionDetails, bool) (domain.ProvisionedServiceSpec, error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.ProvisionDetails
		arg4 bool
	}
	provisionReturns struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}
	ServicesStub        func(context.Context) ([]domain.Service, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
		arg1 context.Context
	}
	servicesReturns struct {
		result1 []domain.Service
		result2 error
	}
	servicesReturnsOnCall map[int]struct {
		result1 []domain.Service
		result2 error
	}
	UnbindStub        func(context.Context, string, string, domain.UnbindDetails, bool) (domain.UnbindSpec, error)
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.UnbindDetails
		arg5 bool
	}
	unbindReturns struct {
		result1 domain.UnbindSpec
		result2 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 domain.UnbindSpec
		result2 error
	}
	UpdateStub        func(context.Context, string, domain.UpdateDetails, bool) (domain.UpdateServiceSpec, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.UpdateDetails
		arg4 bool
	}
	updateReturns struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBroker) Bind(arg1 context.Context, arg2 string, arg3 string, arg4 domain.BindDetails, arg5 bool) (domain.Binding, error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.BindDetails
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.BindStub
	fakeReturns := fake.bindReturns
	fake.recordInvocation("Bind", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.bindMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeServiceBroker) BindCalls(stub func(context.Context, string, string, domain.BindDetails, bool) (domain.Binding, error)) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = stub
}

func (fake *FakeServiceBroker) BindArgsForCall(i int) (context.Context, string, string, domain.BindDetails, bool) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	argsForCall := fake.bindArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeServiceBroker) BindReturns(result1 domain.Binding, result2 error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 domain.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) BindReturnsOnCall(i int, result1 domain.Binding, result2 error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 domain.Binding
			result2 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 domain.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) Deprovision(arg1 context.Context, arg2 string, arg3 domain.DeprovisionDetails, arg4 bool) (domain.DeprovisionServiceSpec, error) {
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.DeprovisionDetails
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.DeprovisionStub
	fakeReturns := fake.deprovisionReturns
	fake.recordInvocation("Deprovision", []interface{}{arg1, arg2, arg3, arg4})
	fake.deprovisionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeServiceBroker) DeprovisionCalls(stub func(context.Context, string, domain.DeprovisionDetails, bool) (domain.DeprovisionServiceSpec, error)) {
	fake.deprovisionMutex.Lock()
	defer fake.deprovisionMutex.Unlock()
	fake.DeprovisionStub = stub
}

func (fake *FakeServiceBroker) DeprovisionArgsForCall(i int) (context.Context, string, domain.DeprovisionDetails, bool) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	argsForCall := fake.deprovisionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeServiceBroker) DeprovisionReturns(result1 domain.DeprovisionServiceSpec, result2 error) {
	fake.deprovisionMutex.Lock()
	defer fake.deprovisionMutex.Unlock()
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) DeprovisionReturnsOnCall(i int, result1 domain.DeprovisionServiceSpec, result2 error) {
	fake.deprovisionMutex.Lock()
	defer fake.deprovisionMutex.Unlock()
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct {
			result1 domain.DeprovisionServiceSpec
			result2 error
		})
	}
	fake.deprovisionReturnsOnCall[i] = struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) GetBinding(arg1 context.Context, arg2 string, arg3 string, arg4 domain.FetchBindingDetails) (domain.GetBindingSpec, error) {
	fake.getBindingMutex.Lock()
	ret, specificReturn := fake.getBindingReturnsOnCall[len(fake.getBindingArgsForCall)]
	fake.getBindingArgsForCall = append(fake.getBindingArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.FetchBindingDetails
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetBindingStub
	fakeReturns := fake.getBindingReturns
	fake.recordInvocation("GetBinding", []interface{}{arg1, arg2, arg3, arg4})
	fake.getBindingMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) GetBindingCallCount() int {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return len(fake.getBindingArgsForCall)
}

func (fake *FakeServiceBroker) GetBindingCalls(stub func(context.Context, string, string, domain.FetchBindingDetails) (domain.GetBindingSpec, error)) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = stub
}

func (fake *FakeServiceBroker) GetBindingArgsForCall(i int) (context.Context, string, string, domain.FetchBindingDetails) {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	argsForCall := fake.getBindingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeServiceBroker) GetBindingReturns(result1 domain.GetBindingSpec, result2 error) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = nil
	fake.getBindingReturns = struct {
		result1 domain.GetBindingSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) GetBindingReturnsOnCall(i int, result1 domain.GetBindingSpec, result2 error) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = nil
	if fake.getBindingReturnsOnCall == nil {
		fake.getBindingReturnsOnCall = make(map[int]struct {
			result1 domain.GetBindingSpec
			result2 error
		})
	}
	fake.getBindingReturnsOnCall[i] = struct {
		result1 domain.GetBindingSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) GetInstance(arg1 context.Context, arg2 string, arg3 domain.FetchInstanceDetails) (domain.GetInstanceDetailsSpec, error) {
	fake.getInstanceMutex.Lock()
	ret, specificReturn := fake.getInstanceReturnsOnCall[len(fake.getInstanceArgsForCall)]
	fake.getInstanceArgsForCall = append(fake.getInstanceArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.FetchInstanceDetails
	}{arg1, arg2, arg3})
	stub := fake.GetInstanceStub
	fakeReturns := fake.getInstanceReturns
	fake.recordInvocation("GetInstance", []interface{}{arg1, arg2, arg3})
	fake.getInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) GetInstanceCallCount() int {
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	return len(fake.getInstanceArgsForCall)
}

func (fake *FakeServiceBroker) GetInstanceCalls(stub func(context.Context, string, domain.FetchInstanceDetails) (domain.GetInstanceDetailsSpec, error)) {
	fake.getInstanceMutex.Lock()
	defer fake.getInstanceMutex.Unlock()
	fake.GetInstanceStub = stub
}

func (fake *FakeServiceBroker) GetInstanceArgsForCall(i int) (context.Context, string, domain.FetchInstanceDetails) {
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	argsForCall := fake.getInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServiceBroker) GetInstanceReturns(result1 domain.GetInstanceDetailsSpec, result2 error) {
	fake.getInstanceMutex.Lock()
	defer fake.getInstanceMutex.Unlock()
	fake.GetInstanceStub = nil
	fake.getInstanceReturns = struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) GetInstanceReturnsOnCall(i int, result1 domain.GetInstanceDetailsSpec, result2 error) {
	fake.getInstanceMutex.Lock()
	defer fake.getInstanceMutex.Unlock()
	fake.GetInstanceStub = nil
	if fake.getInstanceReturnsOnCall == nil {
		fake.getInstanceReturnsOnCall = make(map[int]struct {
			result1 domain.GetInstanceDetailsSpec
			result2 error
		})
	}
	fake.getInstanceReturnsOnCall[i] = struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) LastBindingOperation(arg1 context.Context, arg2 string, arg3 string, arg4 domain.PollDetails) (domain.LastOperation, error) {
	fake.lastBindingOperationMutex.Lock()
	ret, specificReturn := fake.lastBindingOperationReturnsOnCall[len(fake.lastBindingOperationArgsForCall)]
	fake.lastBindingOperationArgsForCall = append(fake.lastBindingOperationArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.PollDetails
	}{arg1, arg2, arg3, arg4})
	stub := fake.LastBindingOperationStub
	fakeReturns := fake.lastBindingOperationReturns
	fake.recordInvocation("LastBindingOperation", []interface{}{arg1, arg2, arg3, arg4})
	fake.lastBindingOperationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) LastBindingOperationCallCount() int {
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	return len(fake.lastBindingOperationArgsForCall)
}

func (fake *FakeServiceBroker) LastBindingOperationCalls(stub func(context.Context, string, string, domain.PollDetails) (domain.LastOperation, error)) {
	fake.lastBindingOperationMutex.Lock()
	defer fake.lastBindingOperationMutex.Unlock()
	fake.LastBindingOperationStub = stub
}

func (fake *FakeServiceBroker) LastBindingOperationArgsForCall(i int) (context.Context, string, string, domain.PollDetails) {
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	argsForCall := fake.lastBindingOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeServiceBroker) LastBindingOperationReturns(result1 domain.LastOperation, result2 error) {
	fake.lastBindingOperationMutex.Lock()
	defer fake.lastBindingOperationMutex.Unlock()
	fake.LastBindingOperationStub = nil
	fake.lastBindingOperationReturns = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) LastBindingOperationReturnsOnCall(i int, result1 domain.LastOperation, result2 error) {
	fake.lastBindingOperationMutex.Lock()
	defer fake.lastBindingOperationMutex.Unlock()
	fake.LastBindingOperationStub = nil
	if fake.lastBindingOperationReturnsOnCall == nil {
		fake.lastBindingOperationReturnsOnCall = make(map[int]struct {
			result1 domain.LastOperation
			result2 error
		})
	}
	fake.lastBindingOperationReturnsOnCall[i] = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) LastOperation(arg1 context.Context, arg2 string, arg3 domain.PollDetails) (domain.LastOperation, error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.PollDetails
	}{arg1, arg2, arg3})
	stub := fake.LastOperationStub
	fakeReturns := fake.lastOperationReturns
	fake.recordInvocation("LastOperation", []interface{}{arg1, arg2, arg3})
	fake.lastOperationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *FakeServiceBroker) LastOperationCalls(stub func(context.Context, string, domain.PollDetails) (domain.LastOperation, error)) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = stub
}

func (fake *FakeServiceBroker) LastOperationArgsForCall(i int) (context.Context, string, domain.PollDetails) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	argsForCall := fake.lastOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServiceBroker) LastOperationReturns(result1 domain.LastOperation, result2 error) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) LastOperationReturnsOnCall(i int, result1 domain.LastOperation, result2 error) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = nil
	if fake.lastOperationReturnsOnCall == nil {
		fake.lastOperationReturnsOnCall = make(map[int]struct {
			result1 domain.LastOperation
			result2 error
		})
	}
	fake.lastOperationReturnsOnCall[i] = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) Provision(arg1 context.Context, arg2 string, arg3 domain.ProvisionDetails, arg4 bool) (domain.ProvisionedServiceSpec, error) {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.ProvisionDetails
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.ProvisionStub
	fakeReturns := fake.provisionReturns
	fake.recordInvocation("Provision", []interface{}{arg1, arg2, arg3, arg4})
	fake.provisionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeServiceBroker) ProvisionCalls(stub func(context.Context, string, domain.ProvisionDetails, bool) (domain.ProvisionedServiceSpec, error)) {
	fake.provisionMutex.Lock()
	defer fake.provisionMutex.Unlock()
	fake.ProvisionStub = stub
}

func (fake *FakeServiceBroker) ProvisionArgsForCall(i int) (context.Context, string, domain.ProvisionDetails, bool) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	argsForCall := fake.provisionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeServiceBroker) ProvisionReturns(result1 domain.ProvisionedServiceSpec, result2 error) {
	fake.provisionMutex.Lock()
	defer fake.provisionMutex.Unlock()
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) ProvisionReturnsOnCall(i int, result1 domain.ProvisionedServiceSpec, result2 error) {
	fake.provisionMutex.Lock()
	defer fake.provisionMutex.Unlock()
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 domain.ProvisionedServiceSpec
			result2 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) Services(arg1 context.Context) ([]domain.Service, error) {
	fake.servicesMutex.Lock()
	ret, specificReturn := fake.servicesReturnsOnCall[len(fake.servicesArgsForCall)]
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ServicesStub
	fakeReturns := fake.servicesReturns
	fake.recordInvocation("Services", []interface{}{arg1})
	fake.servicesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeServiceBroker) ServicesCalls(stub func(context.Context) ([]domain.Service, error)) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = stub
}

func (fake *FakeServiceBroker) ServicesArgsForCall(i int) context.Context {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	argsForCall := fake.servicesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceBroker) ServicesReturns(result1 []domain.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []domain.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) ServicesReturnsOnCall(i int, result1 []domain.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 []domain.Service
			result2 error
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 []domain.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) Unbind(arg1 context.Context, arg2 string, arg3 string, arg4 domain.UnbindDetails, arg5 bool) (domain.UnbindSpec, error) {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.UnbindDetails
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.UnbindStub
	fakeReturns := fake.unbindReturns
	fake.recordInvocation("Unbind", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.unbindMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeServiceBroker) UnbindCalls(stub func(context.Context, string, string, domain.UnbindDetails, bool) (domain.UnbindSpec, error)) {
	fake.unbindMutex.Lock()
	defer fake.unbindMutex.Unlock()
	fake.UnbindStub = stub
}

func (fake *FakeServiceBroker) UnbindArgsForCall(i int) (context.Context, string, string, domain.UnbindDetails, bool) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	argsForCall := fake.unbindArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeServiceBroker) UnbindReturns(result1 domain.UnbindSpec, result2 error) {
	fake.unbindMutex.Lock()
	defer fake.unbindMutex.Unlock()
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 domain.UnbindSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) UnbindReturnsOnCall(i int, result1 domain.UnbindSpec, result2 error) {
	fake.unbindMutex.Lock()
	defer fake.unbindMutex.Unlock()
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 domain.UnbindSpec
			result2 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 domain.UnbindSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) Update(arg1 context.Context, arg2 string, arg3 domain.UpdateDetails, arg4 bool) (domain.UpdateServiceSpec, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.UpdateDetails
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBroker) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceBroker) UpdateCalls(stub func(context.Context, string, domain.UpdateDetails, bool) (domain.UpdateServiceSpec, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeServiceBroker) UpdateArgsForCall(i int) (context.Context, string, domain.UpdateDetails, bool) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeServiceBroker) UpdateReturns(result1 domain.UpdateServiceSpec, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) UpdateReturnsOnCall(i int, result1 domain.UpdateServiceSpec, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 domain.UpdateServiceSpec
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBroker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceBroker) recordInvocation(key string, args []interface{}) {
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

var _ domain.ServiceBroker = new(FakeServiceBroker)
