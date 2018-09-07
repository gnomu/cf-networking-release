// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type EgressDestinationRepo struct {
	AllStub        func(tx db.Transaction) ([]store.EgressDestination, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
		tx db.Transaction
	}
	allReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	CreateIPRangeStub        func(tx db.Transaction, destinationTerminalGUID, startIP, endIP, protocol string, startPort, endPort, icmpType, icmpCode int64) (int64, error)
	createIPRangeMutex       sync.RWMutex
	createIPRangeArgsForCall []struct {
		tx                      db.Transaction
		destinationTerminalGUID string
		startIP                 string
		endIP                   string
		protocol                string
		startPort               int64
		endPort                 int64
		icmpType                int64
		icmpCode                int64
	}
	createIPRangeReturns struct {
		result1 int64
		result2 error
	}
	createIPRangeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressDestinationRepo) All(tx db.Transaction) ([]store.EgressDestination, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct {
		tx db.Transaction
	}{tx})
	fake.recordInvocation("All", []interface{}{tx})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub(tx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *EgressDestinationRepo) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressDestinationRepo) AllArgsForCall(i int) db.Transaction {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return fake.allArgsForCall[i].tx
}

func (fake *EgressDestinationRepo) AllReturns(result1 []store.EgressDestination, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) AllReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) CreateIPRange(tx db.Transaction, destinationTerminalGUID string, startIP string, endIP string, protocol string, startPort int64, endPort int64, icmpType int64, icmpCode int64) (int64, error) {
	fake.createIPRangeMutex.Lock()
	ret, specificReturn := fake.createIPRangeReturnsOnCall[len(fake.createIPRangeArgsForCall)]
	fake.createIPRangeArgsForCall = append(fake.createIPRangeArgsForCall, struct {
		tx                      db.Transaction
		destinationTerminalGUID string
		startIP                 string
		endIP                   string
		protocol                string
		startPort               int64
		endPort                 int64
		icmpType                int64
		icmpCode                int64
	}{tx, destinationTerminalGUID, startIP, endIP, protocol, startPort, endPort, icmpType, icmpCode})
	fake.recordInvocation("CreateIPRange", []interface{}{tx, destinationTerminalGUID, startIP, endIP, protocol, startPort, endPort, icmpType, icmpCode})
	fake.createIPRangeMutex.Unlock()
	if fake.CreateIPRangeStub != nil {
		return fake.CreateIPRangeStub(tx, destinationTerminalGUID, startIP, endIP, protocol, startPort, endPort, icmpType, icmpCode)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIPRangeReturns.result1, fake.createIPRangeReturns.result2
}

func (fake *EgressDestinationRepo) CreateIPRangeCallCount() int {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return len(fake.createIPRangeArgsForCall)
}

func (fake *EgressDestinationRepo) CreateIPRangeArgsForCall(i int) (db.Transaction, string, string, string, string, int64, int64, int64, int64) {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return fake.createIPRangeArgsForCall[i].tx, fake.createIPRangeArgsForCall[i].destinationTerminalGUID, fake.createIPRangeArgsForCall[i].startIP, fake.createIPRangeArgsForCall[i].endIP, fake.createIPRangeArgsForCall[i].protocol, fake.createIPRangeArgsForCall[i].startPort, fake.createIPRangeArgsForCall[i].endPort, fake.createIPRangeArgsForCall[i].icmpType, fake.createIPRangeArgsForCall[i].icmpCode
}

func (fake *EgressDestinationRepo) CreateIPRangeReturns(result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	fake.createIPRangeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) CreateIPRangeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	if fake.createIPRangeReturnsOnCall == nil {
		fake.createIPRangeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createIPRangeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressDestinationRepo) recordInvocation(key string, args []interface{}) {
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
