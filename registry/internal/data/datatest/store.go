// Code autogenerated by mockery v3
//
// Do not manually edit the content of this file.

// Package datatest contains autogenerated mocks.
package datatest

import "oss.indeed.com/go/modprox/pkg/coordinates"

import "github.com/stretchr/testify/mock"
import "oss.indeed.com/go/modprox/pkg/netservice"
import "oss.indeed.com/go/modprox/pkg/clients/payloads"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// DeleteModuleByID provides a mock function with given fields: id
func (mockerySelf *Store) DeleteModuleByID(id int) error {
	ret := mockerySelf.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertModules provides a mock function with given fields: mockeryArg0
func (mockerySelf *Store) InsertModules(mockeryArg0 []coordinates.Module) (int, error) {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 int
	if rf, ok := ret.Get(0).(func([]coordinates.Module) int); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]coordinates.Module) error); ok {
		r1 = rf(mockeryArg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHeartbeats provides a mock function with given fields:
func (mockerySelf *Store) ListHeartbeats() ([]payloads.Heartbeat, error) {
	ret := mockerySelf.Called()

	var r0 []payloads.Heartbeat
	if rf, ok := ret.Get(0).(func() []payloads.Heartbeat); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payloads.Heartbeat)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModuleIDs provides a mock function with given fields:
func (mockerySelf *Store) ListModuleIDs() ([]int64, error) {
	ret := mockerySelf.Called()

	var r0 []int64
	if rf, ok := ret.Get(0).(func() []int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModules provides a mock function with given fields:
func (mockerySelf *Store) ListModules() ([]coordinates.SerialModule, error) {
	ret := mockerySelf.Called()

	var r0 []coordinates.SerialModule
	if rf, ok := ret.Get(0).(func() []coordinates.SerialModule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coordinates.SerialModule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModulesByIDs provides a mock function with given fields: ids
func (mockerySelf *Store) ListModulesByIDs(ids []int64) ([]coordinates.SerialModule, error) {
	ret := mockerySelf.Called(ids)

	var r0 []coordinates.SerialModule
	if rf, ok := ret.Get(0).(func([]int64) []coordinates.SerialModule); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coordinates.SerialModule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModulesBySource provides a mock function with given fields: source
func (mockerySelf *Store) ListModulesBySource(source string) ([]coordinates.SerialModule, error) {
	ret := mockerySelf.Called(source)

	var r0 []coordinates.SerialModule
	if rf, ok := ret.Get(0).(func(string) []coordinates.SerialModule); ok {
		r0 = rf(source)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coordinates.SerialModule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(source)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStartConfigs provides a mock function with given fields:
func (mockerySelf *Store) ListStartConfigs() ([]payloads.Configuration, error) {
	ret := mockerySelf.Called()

	var r0 []payloads.Configuration
	if rf, ok := ret.Get(0).(func() []payloads.Configuration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payloads.Configuration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurgeProxy provides a mock function with given fields: instance
func (mockerySelf *Store) PurgeProxy(instance netservice.Instance) error {
	ret := mockerySelf.Called(instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(netservice.Instance) error); ok {
		r0 = rf(instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetHeartbeat provides a mock function with given fields: mockeryArg0
func (mockerySelf *Store) SetHeartbeat(mockeryArg0 payloads.Heartbeat) error {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 error
	if rf, ok := ret.Get(0).(func(payloads.Heartbeat) error); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetStartConfig provides a mock function with given fields: mockeryArg0
func (mockerySelf *Store) SetStartConfig(mockeryArg0 payloads.Configuration) error {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 error
	if rf, ok := ret.Get(0).(func(payloads.Configuration) error); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
