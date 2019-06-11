// Code autogenerated by mockery v3
//
// Do not manually edit the content of this file.

// Package gettest contains autogenerated mocks.
package gettest

import "oss.indeed.com/go/modprox/pkg/coordinates"

import "github.com/stretchr/testify/mock"

// Downloader is an autogenerated mock type for the Downloader type
type Downloader struct {
	mock.Mock
}

// Download provides a mock function with given fields: module
func (mockerySelf *Downloader) Download(module coordinates.SerialModule) error {
	ret := mockerySelf.Called(module)

	var r0 error
	if rf, ok := ret.Get(0).(func(coordinates.SerialModule) error); ok {
		r0 = rf(module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
