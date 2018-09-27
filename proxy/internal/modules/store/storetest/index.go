// Code autogenerated by mockery v3
//
// Do not manually edit the content of this file.

// Package storetest contains autogenerated mocks.
package storetest

import "github.com/modprox/mp/pkg/coordinates"
import "github.com/stretchr/testify/mock"
import "github.com/modprox/mp/pkg/repository"
import "github.com/modprox/mp/proxy/internal/modules/store"

// Index is an autogenerated mock type for the Index type
type Index struct {
	mock.Mock
}

// Contains provides a mock function with given fields: mockeryArg0
func (mockerySelf *Index) Contains(mockeryArg0 coordinates.Module) (bool, error) {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(coordinates.Module) bool); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(coordinates.Module) error); ok {
		r1 = rf(mockeryArg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IDs provides a mock function with given fields:
func (mockerySelf *Index) IDs() (coordinates.RangeIDs, error) {
	ret := mockerySelf.Called()

	var r0 coordinates.RangeIDs
	if rf, ok := ret.Get(0).(func() coordinates.RangeIDs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coordinates.RangeIDs)
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

// Info provides a mock function with given fields: mockeryArg0
func (mockerySelf *Index) Info(mockeryArg0 coordinates.Module) (repository.RevInfo, error) {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 repository.RevInfo
	if rf, ok := ret.Get(0).(func(coordinates.Module) repository.RevInfo); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Get(0).(repository.RevInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(coordinates.Module) error); ok {
		r1 = rf(mockeryArg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mod provides a mock function with given fields: mockeryArg0
func (mockerySelf *Index) Mod(mockeryArg0 coordinates.Module) (string, error) {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 string
	if rf, ok := ret.Get(0).(func(coordinates.Module) string); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(coordinates.Module) error); ok {
		r1 = rf(mockeryArg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: mockeryArg0
func (mockerySelf *Index) Put(mockeryArg0 store.ModuleAddition) error {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 error
	if rf, ok := ret.Get(0).(func(store.ModuleAddition) error); ok {
		r0 = rf(mockeryArg0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Summary provides a mock function with given fields:
func (mockerySelf *Index) Summary() (int, int, error) {
	ret := mockerySelf.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Versions provides a mock function with given fields: module
func (mockerySelf *Index) Versions(module string) ([]string, error) {
	ret := mockerySelf.Called(module)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(module)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
