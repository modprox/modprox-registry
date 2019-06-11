// Code autogenerated by mockery v3
//
// Do not manually edit the content of this file.

// Package zipstest contains autogenerated mocks.
package zipstest

import "github.com/stretchr/testify/mock"
import "oss.indeed.com/go/modprox/pkg/repository"
import "oss.indeed.com/go/modprox/pkg/upstream"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Get provides a mock function with given fields: mockeryArg0
func (mockerySelf *Client) Get(mockeryArg0 *upstream.Request) (repository.Blob, error) {
	ret := mockerySelf.Called(mockeryArg0)

	var r0 repository.Blob
	if rf, ok := ret.Get(0).(func(*upstream.Request) repository.Blob); ok {
		r0 = rf(mockeryArg0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*upstream.Request) error); ok {
		r1 = rf(mockeryArg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Protocols provides a mock function with given fields:
func (mockerySelf *Client) Protocols() []string {
	ret := mockerySelf.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
