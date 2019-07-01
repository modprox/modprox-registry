package zips

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"

	"oss.indeed.com/go/modprox/pkg/repository"
	"oss.indeed.com/go/modprox/pkg/upstream"
)

// UpstreamClientMock implements UpstreamClient
type UpstreamClientMock struct {
	t minimock.Tester

	funcGet          func(rp1 *upstream.Request) (b1 repository.Blob, err error)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUpstreamClientMockGet

	funcProtocols          func() (sa1 []string)
	afterProtocolsCounter  uint64
	beforeProtocolsCounter uint64
	ProtocolsMock          mUpstreamClientMockProtocols
}

// NewUpstreamClientMock returns a mock for UpstreamClient
func NewUpstreamClientMock(t minimock.Tester) *UpstreamClientMock {
	m := &UpstreamClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetMock = mUpstreamClientMockGet{mock: m}
	m.GetMock.callArgs = []*UpstreamClientMockGetParams{}

	m.ProtocolsMock = mUpstreamClientMockProtocols{mock: m}

	return m
}

type mUpstreamClientMockGet struct {
	mock               *UpstreamClientMock
	defaultExpectation *UpstreamClientMockGetExpectation
	expectations       []*UpstreamClientMockGetExpectation

	callArgs []*UpstreamClientMockGetParams
	mutex    sync.RWMutex
}

// UpstreamClientMockGetExpectation specifies expectation struct of the UpstreamClient.Get
type UpstreamClientMockGetExpectation struct {
	mock    *UpstreamClientMock
	params  *UpstreamClientMockGetParams
	results *UpstreamClientMockGetResults
	Counter uint64
}

// UpstreamClientMockGetParams contains parameters of the UpstreamClient.Get
type UpstreamClientMockGetParams struct {
	rp1 *upstream.Request
}

// UpstreamClientMockGetResults contains results of the UpstreamClient.Get
type UpstreamClientMockGetResults struct {
	b1  repository.Blob
	err error
}

// Expect sets up expected params for UpstreamClient.Get
func (mmGet *mUpstreamClientMockGet) Expect(rp1 *upstream.Request) *mUpstreamClientMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UpstreamClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UpstreamClientMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &UpstreamClientMockGetParams{rp1}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Return sets up results that will be returned by UpstreamClient.Get
func (mmGet *mUpstreamClientMockGet) Return(b1 repository.Blob, err error) *UpstreamClientMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UpstreamClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UpstreamClientMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UpstreamClientMockGetResults{b1, err}
	return mmGet.mock
}

//Set uses given function f to mock the UpstreamClient.Get method
func (mmGet *mUpstreamClientMockGet) Set(f func(rp1 *upstream.Request) (b1 repository.Blob, err error)) *UpstreamClientMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UpstreamClient.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UpstreamClient.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the UpstreamClient.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUpstreamClientMockGet) When(rp1 *upstream.Request) *UpstreamClientMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UpstreamClientMock.Get mock is already set by Set")
	}

	expectation := &UpstreamClientMockGetExpectation{
		mock:   mmGet.mock,
		params: &UpstreamClientMockGetParams{rp1},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UpstreamClient.Get return parameters for the expectation previously defined by the When method
func (e *UpstreamClientMockGetExpectation) Then(b1 repository.Blob, err error) *UpstreamClientMock {
	e.results = &UpstreamClientMockGetResults{b1, err}
	return e.mock
}

// Get implements UpstreamClient
func (mmGet *UpstreamClientMock) Get(rp1 *upstream.Request) (b1 repository.Blob, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	params := &UpstreamClientMockGetParams{rp1}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		want := mmGet.GetMock.defaultExpectation.params
		got := UpstreamClientMockGetParams{rp1}
		if want != nil && !minimock.Equal(*want, got) {
			mmGet.t.Errorf("UpstreamClientMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmGet.GetMock.defaultExpectation.results
		if results == nil {
			mmGet.t.Fatal("No results are set for the UpstreamClientMock.Get")
		}
		return (*results).b1, (*results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(rp1)
	}
	mmGet.t.Fatalf("Unexpected call to UpstreamClientMock.Get. %v", rp1)
	return
}

// GetAfterCounter returns a count of finished UpstreamClientMock.Get invocations
func (mmGet *UpstreamClientMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UpstreamClientMock.Get invocations
func (mmGet *UpstreamClientMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UpstreamClientMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUpstreamClientMockGet) Calls() []*UpstreamClientMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UpstreamClientMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UpstreamClientMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *UpstreamClientMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UpstreamClientMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UpstreamClientMock.Get")
		} else {
			m.t.Errorf("Expected call to UpstreamClientMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to UpstreamClientMock.Get")
	}
}

type mUpstreamClientMockProtocols struct {
	mock               *UpstreamClientMock
	defaultExpectation *UpstreamClientMockProtocolsExpectation
	expectations       []*UpstreamClientMockProtocolsExpectation
}

// UpstreamClientMockProtocolsExpectation specifies expectation struct of the UpstreamClient.Protocols
type UpstreamClientMockProtocolsExpectation struct {
	mock *UpstreamClientMock

	results *UpstreamClientMockProtocolsResults
	Counter uint64
}

// UpstreamClientMockProtocolsResults contains results of the UpstreamClient.Protocols
type UpstreamClientMockProtocolsResults struct {
	sa1 []string
}

// Expect sets up expected params for UpstreamClient.Protocols
func (mmProtocols *mUpstreamClientMockProtocols) Expect() *mUpstreamClientMockProtocols {
	if mmProtocols.mock.funcProtocols != nil {
		mmProtocols.mock.t.Fatalf("UpstreamClientMock.Protocols mock is already set by Set")
	}

	if mmProtocols.defaultExpectation == nil {
		mmProtocols.defaultExpectation = &UpstreamClientMockProtocolsExpectation{}
	}

	return mmProtocols
}

// Return sets up results that will be returned by UpstreamClient.Protocols
func (mmProtocols *mUpstreamClientMockProtocols) Return(sa1 []string) *UpstreamClientMock {
	if mmProtocols.mock.funcProtocols != nil {
		mmProtocols.mock.t.Fatalf("UpstreamClientMock.Protocols mock is already set by Set")
	}

	if mmProtocols.defaultExpectation == nil {
		mmProtocols.defaultExpectation = &UpstreamClientMockProtocolsExpectation{mock: mmProtocols.mock}
	}
	mmProtocols.defaultExpectation.results = &UpstreamClientMockProtocolsResults{sa1}
	return mmProtocols.mock
}

//Set uses given function f to mock the UpstreamClient.Protocols method
func (mmProtocols *mUpstreamClientMockProtocols) Set(f func() (sa1 []string)) *UpstreamClientMock {
	if mmProtocols.defaultExpectation != nil {
		mmProtocols.mock.t.Fatalf("Default expectation is already set for the UpstreamClient.Protocols method")
	}

	if len(mmProtocols.expectations) > 0 {
		mmProtocols.mock.t.Fatalf("Some expectations are already set for the UpstreamClient.Protocols method")
	}

	mmProtocols.mock.funcProtocols = f
	return mmProtocols.mock
}

// Protocols implements UpstreamClient
func (mmProtocols *UpstreamClientMock) Protocols() (sa1 []string) {
	mm_atomic.AddUint64(&mmProtocols.beforeProtocolsCounter, 1)
	defer mm_atomic.AddUint64(&mmProtocols.afterProtocolsCounter, 1)

	if mmProtocols.ProtocolsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmProtocols.ProtocolsMock.defaultExpectation.Counter, 1)

		results := mmProtocols.ProtocolsMock.defaultExpectation.results
		if results == nil {
			mmProtocols.t.Fatal("No results are set for the UpstreamClientMock.Protocols")
		}
		return (*results).sa1
	}
	if mmProtocols.funcProtocols != nil {
		return mmProtocols.funcProtocols()
	}
	mmProtocols.t.Fatalf("Unexpected call to UpstreamClientMock.Protocols.")
	return
}

// ProtocolsAfterCounter returns a count of finished UpstreamClientMock.Protocols invocations
func (mmProtocols *UpstreamClientMock) ProtocolsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProtocols.afterProtocolsCounter)
}

// ProtocolsBeforeCounter returns a count of UpstreamClientMock.Protocols invocations
func (mmProtocols *UpstreamClientMock) ProtocolsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProtocols.beforeProtocolsCounter)
}

// MinimockProtocolsDone returns true if the count of the Protocols invocations corresponds
// the number of defined expectations
func (m *UpstreamClientMock) MinimockProtocolsDone() bool {
	for _, e := range m.ProtocolsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProtocolsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProtocolsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProtocols != nil && mm_atomic.LoadUint64(&m.afterProtocolsCounter) < 1 {
		return false
	}
	return true
}

// MinimockProtocolsInspect logs each unmet expectation
func (m *UpstreamClientMock) MinimockProtocolsInspect() {
	for _, e := range m.ProtocolsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to UpstreamClientMock.Protocols")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProtocolsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProtocolsCounter) < 1 {
		m.t.Error("Expected call to UpstreamClientMock.Protocols")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProtocols != nil && mm_atomic.LoadUint64(&m.afterProtocolsCounter) < 1 {
		m.t.Error("Expected call to UpstreamClientMock.Protocols")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UpstreamClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetInspect()

		m.MinimockProtocolsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UpstreamClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UpstreamClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetDone() &&
		m.MinimockProtocolsDone()
}
