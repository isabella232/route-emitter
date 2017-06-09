// This file was generated by counterfeiter
package fakeroutingtable

import (
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/route-emitter/routingtable"
	"code.cloudfoundry.org/route-emitter/routingtable/schema/endpoint"
)

type FakeNATSRoutingTable struct {
	AddEndpointStub        func(key endpoint.RoutingKey, routingEndpoint routingtable.Endpoint) routingtable.MessagesToEmit
	addEndpointMutex       sync.RWMutex
	addEndpointArgsForCall []struct {
		key             endpoint.RoutingKey
		routingEndpoint routingtable.Endpoint
	}
	addEndpointReturns struct {
		result1 routingtable.MessagesToEmit
	}
	SetRoutesStub        func(key endpoint.RoutingKey, routes []routingtable.Route, modTag *models.ModificationTag) routingtable.MessagesToEmit
	setRoutesMutex       sync.RWMutex
	setRoutesArgsForCall []struct {
		key    endpoint.RoutingKey
		routes []routingtable.Route
		modTag *models.ModificationTag
	}
	setRoutesReturns struct {
		result1 routingtable.MessagesToEmit
	}
	EndpointsForIndexStub        func(key endpoint.RoutingKey, index int32) []routingtable.Endpoint
	endpointsForIndexMutex       sync.RWMutex
	endpointsForIndexArgsForCall []struct {
		key   endpoint.RoutingKey
		index int32
	}
	endpointsForIndexReturns struct {
		result1 []routingtable.Endpoint
	}
	GetRoutesStub        func(key endpoint.RoutingKey) []routingtable.Route
	getRoutesMutex       sync.RWMutex
	getRoutesArgsForCall []struct {
		key endpoint.RoutingKey
	}
	getRoutesReturns struct {
		result1 []routingtable.Route
	}
	MessagesToEmitStub        func() routingtable.MessagesToEmit
	messagesToEmitMutex       sync.RWMutex
	messagesToEmitArgsForCall []struct{}
	messagesToEmitReturns     struct {
		result1 routingtable.MessagesToEmit
	}
	RemoveEndpointStub        func(key endpoint.RoutingKey, routingEndpoint routingtable.Endpoint) routingtable.MessagesToEmit
	removeEndpointMutex       sync.RWMutex
	removeEndpointArgsForCall []struct {
		key             endpoint.RoutingKey
		routingEndpoint routingtable.Endpoint
	}
	removeEndpointReturns struct {
		result1 routingtable.MessagesToEmit
	}
	RemoveRoutesStub        func(key endpoint.RoutingKey, modTag *models.ModificationTag) routingtable.MessagesToEmit
	removeRoutesMutex       sync.RWMutex
	removeRoutesArgsForCall []struct {
		key    endpoint.RoutingKey
		modTag *models.ModificationTag
	}
	removeRoutesReturns struct {
		result1 routingtable.MessagesToEmit
	}
	RouteCountStub        func() int
	routeCountMutex       sync.RWMutex
	routeCountArgsForCall []struct{}
	routeCountReturns     struct {
		result1 int
	}
	SwapStub        func(newTable routingtable.NATSRoutingTable, domains models.DomainSet) routingtable.MessagesToEmit
	swapMutex       sync.RWMutex
	swapArgsForCall []struct {
		newTable routingtable.NATSRoutingTable
		domains  models.DomainSet
	}
	swapReturns struct {
		result1 routingtable.MessagesToEmit
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNATSRoutingTable) AddEndpoint(key endpoint.RoutingKey, routingEndpoint routingtable.Endpoint) routingtable.MessagesToEmit {
	fake.addEndpointMutex.Lock()
	fake.addEndpointArgsForCall = append(fake.addEndpointArgsForCall, struct {
		key             endpoint.RoutingKey
		routingEndpoint routingtable.Endpoint
	}{key, routingEndpoint})
	fake.recordInvocation("AddEndpoint", []interface{}{key, routingEndpoint})
	fake.addEndpointMutex.Unlock()
	if fake.AddEndpointStub != nil {
		return fake.AddEndpointStub(key, routingEndpoint)
	} else {
		return fake.addEndpointReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) AddEndpointCallCount() int {
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	return len(fake.addEndpointArgsForCall)
}

func (fake *FakeNATSRoutingTable) AddEndpointArgsForCall(i int) (endpoint.RoutingKey, routingtable.Endpoint) {
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	return fake.addEndpointArgsForCall[i].key, fake.addEndpointArgsForCall[i].routingEndpoint
}

func (fake *FakeNATSRoutingTable) AddEndpointReturns(result1 routingtable.MessagesToEmit) {
	fake.AddEndpointStub = nil
	fake.addEndpointReturns = struct {
		result1 routingtable.MessagesToEmit
	}{result1}
}

func (fake *FakeNATSRoutingTable) SetRoutes(key endpoint.RoutingKey, routes []routingtable.Route, modTag *models.ModificationTag) routingtable.MessagesToEmit {
	var routesCopy []routingtable.Route
	if routes != nil {
		routesCopy = make([]routingtable.Route, len(routes))
		copy(routesCopy, routes)
	}
	fake.setRoutesMutex.Lock()
	fake.setRoutesArgsForCall = append(fake.setRoutesArgsForCall, struct {
		key    endpoint.RoutingKey
		routes []routingtable.Route
		modTag *models.ModificationTag
	}{key, routesCopy, modTag})
	fake.recordInvocation("SetRoutes", []interface{}{key, routesCopy, modTag})
	fake.setRoutesMutex.Unlock()
	if fake.SetRoutesStub != nil {
		return fake.SetRoutesStub(key, routes, modTag)
	} else {
		return fake.setRoutesReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) SetRoutesCallCount() int {
	fake.setRoutesMutex.RLock()
	defer fake.setRoutesMutex.RUnlock()
	return len(fake.setRoutesArgsForCall)
}

func (fake *FakeNATSRoutingTable) SetRoutesArgsForCall(i int) (endpoint.RoutingKey, []routingtable.Route, *models.ModificationTag) {
	fake.setRoutesMutex.RLock()
	defer fake.setRoutesMutex.RUnlock()
	return fake.setRoutesArgsForCall[i].key, fake.setRoutesArgsForCall[i].routes, fake.setRoutesArgsForCall[i].modTag
}

func (fake *FakeNATSRoutingTable) SetRoutesReturns(result1 routingtable.MessagesToEmit) {
	fake.SetRoutesStub = nil
	fake.setRoutesReturns = struct {
		result1 routingtable.MessagesToEmit
	}{result1}
}

func (fake *FakeNATSRoutingTable) EndpointsForIndex(key endpoint.RoutingKey, index int32) []routingtable.Endpoint {
	fake.endpointsForIndexMutex.Lock()
	fake.endpointsForIndexArgsForCall = append(fake.endpointsForIndexArgsForCall, struct {
		key   endpoint.RoutingKey
		index int32
	}{key, index})
	fake.recordInvocation("EndpointsForIndex", []interface{}{key, index})
	fake.endpointsForIndexMutex.Unlock()
	if fake.EndpointsForIndexStub != nil {
		return fake.EndpointsForIndexStub(key, index)
	} else {
		return fake.endpointsForIndexReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) EndpointsForIndexCallCount() int {
	fake.endpointsForIndexMutex.RLock()
	defer fake.endpointsForIndexMutex.RUnlock()
	return len(fake.endpointsForIndexArgsForCall)
}

func (fake *FakeNATSRoutingTable) EndpointsForIndexArgsForCall(i int) (endpoint.RoutingKey, int32) {
	fake.endpointsForIndexMutex.RLock()
	defer fake.endpointsForIndexMutex.RUnlock()
	return fake.endpointsForIndexArgsForCall[i].key, fake.endpointsForIndexArgsForCall[i].index
}

func (fake *FakeNATSRoutingTable) EndpointsForIndexReturns(result1 []routingtable.Endpoint) {
	fake.EndpointsForIndexStub = nil
	fake.endpointsForIndexReturns = struct {
		result1 []routingtable.Endpoint
	}{result1}
}

func (fake *FakeNATSRoutingTable) GetRoutes(key endpoint.RoutingKey) []routingtable.Route {
	fake.getRoutesMutex.Lock()
	fake.getRoutesArgsForCall = append(fake.getRoutesArgsForCall, struct {
		key endpoint.RoutingKey
	}{key})
	fake.recordInvocation("GetRoutes", []interface{}{key})
	fake.getRoutesMutex.Unlock()
	if fake.GetRoutesStub != nil {
		return fake.GetRoutesStub(key)
	} else {
		return fake.getRoutesReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) GetRoutesCallCount() int {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return len(fake.getRoutesArgsForCall)
}

func (fake *FakeNATSRoutingTable) GetRoutesArgsForCall(i int) endpoint.RoutingKey {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return fake.getRoutesArgsForCall[i].key
}

func (fake *FakeNATSRoutingTable) GetRoutesReturns(result1 []routingtable.Route) {
	fake.GetRoutesStub = nil
	fake.getRoutesReturns = struct {
		result1 []routingtable.Route
	}{result1}
}

func (fake *FakeNATSRoutingTable) MessagesToEmit() routingtable.MessagesToEmit {
	fake.messagesToEmitMutex.Lock()
	fake.messagesToEmitArgsForCall = append(fake.messagesToEmitArgsForCall, struct{}{})
	fake.recordInvocation("MessagesToEmit", []interface{}{})
	fake.messagesToEmitMutex.Unlock()
	if fake.MessagesToEmitStub != nil {
		return fake.MessagesToEmitStub()
	} else {
		return fake.messagesToEmitReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) MessagesToEmitCallCount() int {
	fake.messagesToEmitMutex.RLock()
	defer fake.messagesToEmitMutex.RUnlock()
	return len(fake.messagesToEmitArgsForCall)
}

func (fake *FakeNATSRoutingTable) MessagesToEmitReturns(result1 routingtable.MessagesToEmit) {
	fake.MessagesToEmitStub = nil
	fake.messagesToEmitReturns = struct {
		result1 routingtable.MessagesToEmit
	}{result1}
}

func (fake *FakeNATSRoutingTable) RemoveEndpoint(key endpoint.RoutingKey, routingEndpoint routingtable.Endpoint) routingtable.MessagesToEmit {
	fake.removeEndpointMutex.Lock()
	fake.removeEndpointArgsForCall = append(fake.removeEndpointArgsForCall, struct {
		key             endpoint.RoutingKey
		routingEndpoint routingtable.Endpoint
	}{key, routingEndpoint})
	fake.recordInvocation("RemoveEndpoint", []interface{}{key, routingEndpoint})
	fake.removeEndpointMutex.Unlock()
	if fake.RemoveEndpointStub != nil {
		return fake.RemoveEndpointStub(key, routingEndpoint)
	} else {
		return fake.removeEndpointReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) RemoveEndpointCallCount() int {
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	return len(fake.removeEndpointArgsForCall)
}

func (fake *FakeNATSRoutingTable) RemoveEndpointArgsForCall(i int) (endpoint.RoutingKey, routingtable.Endpoint) {
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	return fake.removeEndpointArgsForCall[i].key, fake.removeEndpointArgsForCall[i].routingEndpoint
}

func (fake *FakeNATSRoutingTable) RemoveEndpointReturns(result1 routingtable.MessagesToEmit) {
	fake.RemoveEndpointStub = nil
	fake.removeEndpointReturns = struct {
		result1 routingtable.MessagesToEmit
	}{result1}
}

func (fake *FakeNATSRoutingTable) RemoveRoutes(key endpoint.RoutingKey, modTag *models.ModificationTag) routingtable.MessagesToEmit {
	fake.removeRoutesMutex.Lock()
	fake.removeRoutesArgsForCall = append(fake.removeRoutesArgsForCall, struct {
		key    endpoint.RoutingKey
		modTag *models.ModificationTag
	}{key, modTag})
	fake.recordInvocation("RemoveRoutes", []interface{}{key, modTag})
	fake.removeRoutesMutex.Unlock()
	if fake.RemoveRoutesStub != nil {
		return fake.RemoveRoutesStub(key, modTag)
	} else {
		return fake.removeRoutesReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) RemoveRoutesCallCount() int {
	fake.removeRoutesMutex.RLock()
	defer fake.removeRoutesMutex.RUnlock()
	return len(fake.removeRoutesArgsForCall)
}

func (fake *FakeNATSRoutingTable) RemoveRoutesArgsForCall(i int) (endpoint.RoutingKey, *models.ModificationTag) {
	fake.removeRoutesMutex.RLock()
	defer fake.removeRoutesMutex.RUnlock()
	return fake.removeRoutesArgsForCall[i].key, fake.removeRoutesArgsForCall[i].modTag
}

func (fake *FakeNATSRoutingTable) RemoveRoutesReturns(result1 routingtable.MessagesToEmit) {
	fake.RemoveRoutesStub = nil
	fake.removeRoutesReturns = struct {
		result1 routingtable.MessagesToEmit
	}{result1}
}

func (fake *FakeNATSRoutingTable) RouteCount() int {
	fake.routeCountMutex.Lock()
	fake.routeCountArgsForCall = append(fake.routeCountArgsForCall, struct{}{})
	fake.recordInvocation("RouteCount", []interface{}{})
	fake.routeCountMutex.Unlock()
	if fake.RouteCountStub != nil {
		return fake.RouteCountStub()
	} else {
		return fake.routeCountReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) RouteCountCallCount() int {
	fake.routeCountMutex.RLock()
	defer fake.routeCountMutex.RUnlock()
	return len(fake.routeCountArgsForCall)
}

func (fake *FakeNATSRoutingTable) RouteCountReturns(result1 int) {
	fake.RouteCountStub = nil
	fake.routeCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeNATSRoutingTable) Swap(newTable routingtable.NATSRoutingTable, domains models.DomainSet) routingtable.MessagesToEmit {
	fake.swapMutex.Lock()
	fake.swapArgsForCall = append(fake.swapArgsForCall, struct {
		newTable routingtable.NATSRoutingTable
		domains  models.DomainSet
	}{newTable, domains})
	fake.recordInvocation("Swap", []interface{}{newTable, domains})
	fake.swapMutex.Unlock()
	if fake.SwapStub != nil {
		return fake.SwapStub(newTable, domains)
	} else {
		return fake.swapReturns.result1
	}
}

func (fake *FakeNATSRoutingTable) SwapCallCount() int {
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return len(fake.swapArgsForCall)
}

func (fake *FakeNATSRoutingTable) SwapArgsForCall(i int) (routingtable.NATSRoutingTable, models.DomainSet) {
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return fake.swapArgsForCall[i].newTable, fake.swapArgsForCall[i].domains
}

func (fake *FakeNATSRoutingTable) SwapReturns(result1 routingtable.MessagesToEmit) {
	fake.SwapStub = nil
	fake.swapReturns = struct {
		result1 routingtable.MessagesToEmit
	}{result1}
}

func (fake *FakeNATSRoutingTable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	fake.setRoutesMutex.RLock()
	defer fake.setRoutesMutex.RUnlock()
	fake.endpointsForIndexMutex.RLock()
	defer fake.endpointsForIndexMutex.RUnlock()
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	fake.messagesToEmitMutex.RLock()
	defer fake.messagesToEmitMutex.RUnlock()
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	fake.removeRoutesMutex.RLock()
	defer fake.removeRoutesMutex.RUnlock()
	fake.routeCountMutex.RLock()
	defer fake.routeCountMutex.RUnlock()
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNATSRoutingTable) recordInvocation(key string, args []interface{}) {
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

var _ routingtable.NATSRoutingTable = new(FakeNATSRoutingTable)
