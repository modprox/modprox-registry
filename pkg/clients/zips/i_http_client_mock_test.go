package zips

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"net/http"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// IHTTPClientMock implements iHTTPClient
type IHTTPClientMock struct {
	t minimock.Tester

	funcDo          func(req *http.Request) (rp1 *http.Response, err error)
	inspectFuncDo   func(req *http.Request)
	afterDoCounter  uint64
	beforeDoCounter uint64
	DoMock          mIHTTPClientMockDo
}

// NewIHTTPClientMock returns a mock for iHTTPClient
func NewIHTTPClientMock(t minimock.Tester) *IHTTPClientMock {
	m := &IHTTPClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DoMock = mIHTTPClientMockDo{mock: m}
	m.DoMock.callArgs = []*IHTTPClientMockDoParams{}

	return m
}

type mIHTTPClientMockDo struct {
	mock               *IHTTPClientMock
	defaultExpectation *IHTTPClientMockDoExpectation
	expectations       []*IHTTPClientMockDoExpectation

	callArgs []*IHTTPClientMockDoParams
	mutex    sync.RWMutex
}

// IHTTPClientMockDoExpectation specifies expectation struct of the iHTTPClient.Do
type IHTTPClientMockDoExpectation struct {
	mock    *IHTTPClientMock
	params  *IHTTPClientMockDoParams
	results *IHTTPClientMockDoResults
	Counter uint64
}

// IHTTPClientMockDoParams contains parameters of the iHTTPClient.Do
type IHTTPClientMockDoParams struct {
	req *http.Request
}

// IHTTPClientMockDoResults contains results of the iHTTPClient.Do
type IHTTPClientMockDoResults struct {
	rp1 *http.Response
	err error
}

// Expect sets up expected params for iHTTPClient.Do
func (mmDo *mIHTTPClientMockDo) Expect(req *http.Request) *mIHTTPClientMockDo {
	if mmDo.mock.funcDo != nil {
		mmDo.mock.t.Fatalf("IHTTPClientMock.Do mock is already set by Set")
	}

	if mmDo.defaultExpectation == nil {
		mmDo.defaultExpectation = &IHTTPClientMockDoExpectation{}
	}

	mmDo.defaultExpectation.params = &IHTTPClientMockDoParams{req}
	for _, e := range mmDo.expectations {
		if minimock.Equal(e.params, mmDo.defaultExpectation.params) {
			mmDo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDo.defaultExpectation.params)
		}
	}

	return mmDo
}

// Inspect accepts an inspector function that has same arguments as the iHTTPClient.Do
func (mmDo *mIHTTPClientMockDo) Inspect(f func(req *http.Request)) *mIHTTPClientMockDo {
	if mmDo.mock.inspectFuncDo != nil {
		mmDo.mock.t.Fatalf("Inspect function is already set for IHTTPClientMock.Do")
	}

	mmDo.mock.inspectFuncDo = f

	return mmDo
}

// Return sets up results that will be returned by iHTTPClient.Do
func (mmDo *mIHTTPClientMockDo) Return(rp1 *http.Response, err error) *IHTTPClientMock {
	if mmDo.mock.funcDo != nil {
		mmDo.mock.t.Fatalf("IHTTPClientMock.Do mock is already set by Set")
	}

	if mmDo.defaultExpectation == nil {
		mmDo.defaultExpectation = &IHTTPClientMockDoExpectation{mock: mmDo.mock}
	}
	mmDo.defaultExpectation.results = &IHTTPClientMockDoResults{rp1, err}
	return mmDo.mock
}

//Set uses given function f to mock the iHTTPClient.Do method
func (mmDo *mIHTTPClientMockDo) Set(f func(req *http.Request) (rp1 *http.Response, err error)) *IHTTPClientMock {
	if mmDo.defaultExpectation != nil {
		mmDo.mock.t.Fatalf("Default expectation is already set for the iHTTPClient.Do method")
	}

	if len(mmDo.expectations) > 0 {
		mmDo.mock.t.Fatalf("Some expectations are already set for the iHTTPClient.Do method")
	}

	mmDo.mock.funcDo = f
	return mmDo.mock
}

// When sets expectation for the iHTTPClient.Do which will trigger the result defined by the following
// Then helper
func (mmDo *mIHTTPClientMockDo) When(req *http.Request) *IHTTPClientMockDoExpectation {
	if mmDo.mock.funcDo != nil {
		mmDo.mock.t.Fatalf("IHTTPClientMock.Do mock is already set by Set")
	}

	expectation := &IHTTPClientMockDoExpectation{
		mock:   mmDo.mock,
		params: &IHTTPClientMockDoParams{req},
	}
	mmDo.expectations = append(mmDo.expectations, expectation)
	return expectation
}

// Then sets up iHTTPClient.Do return parameters for the expectation previously defined by the When method
func (e *IHTTPClientMockDoExpectation) Then(rp1 *http.Response, err error) *IHTTPClientMock {
	e.results = &IHTTPClientMockDoResults{rp1, err}
	return e.mock
}

// Do implements iHTTPClient
func (mmDo *IHTTPClientMock) Do(req *http.Request) (rp1 *http.Response, err error) {
	mm_atomic.AddUint64(&mmDo.beforeDoCounter, 1)
	defer mm_atomic.AddUint64(&mmDo.afterDoCounter, 1)

	if mmDo.inspectFuncDo != nil {
		mmDo.inspectFuncDo(req)
	}

	mm_params := &IHTTPClientMockDoParams{req}

	// Record call args
	mmDo.DoMock.mutex.Lock()
	mmDo.DoMock.callArgs = append(mmDo.DoMock.callArgs, mm_params)
	mmDo.DoMock.mutex.Unlock()

	for _, e := range mmDo.DoMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.rp1, e.results.err
		}
	}

	if mmDo.DoMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDo.DoMock.defaultExpectation.Counter, 1)
		mm_want := mmDo.DoMock.defaultExpectation.params
		mm_got := IHTTPClientMockDoParams{req}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDo.t.Errorf("IHTTPClientMock.Do got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDo.DoMock.defaultExpectation.results
		if mm_results == nil {
			mmDo.t.Fatal("No results are set for the IHTTPClientMock.Do")
		}
		return (*mm_results).rp1, (*mm_results).err
	}
	if mmDo.funcDo != nil {
		return mmDo.funcDo(req)
	}
	mmDo.t.Fatalf("Unexpected call to IHTTPClientMock.Do. %v", req)
	return
}

// DoAfterCounter returns a count of finished IHTTPClientMock.Do invocations
func (mmDo *IHTTPClientMock) DoAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDo.afterDoCounter)
}

// DoBeforeCounter returns a count of IHTTPClientMock.Do invocations
func (mmDo *IHTTPClientMock) DoBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDo.beforeDoCounter)
}

// Calls returns a list of arguments used in each call to IHTTPClientMock.Do.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDo *mIHTTPClientMockDo) Calls() []*IHTTPClientMockDoParams {
	mmDo.mutex.RLock()

	argCopy := make([]*IHTTPClientMockDoParams, len(mmDo.callArgs))
	copy(argCopy, mmDo.callArgs)

	mmDo.mutex.RUnlock()

	return argCopy
}

// MinimockDoDone returns true if the count of the Do invocations corresponds
// the number of defined expectations
func (m *IHTTPClientMock) MinimockDoDone() bool {
	for _, e := range m.DoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDoCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDo != nil && mm_atomic.LoadUint64(&m.afterDoCounter) < 1 {
		return false
	}
	return true
}

// MinimockDoInspect logs each unmet expectation
func (m *IHTTPClientMock) MinimockDoInspect() {
	for _, e := range m.DoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IHTTPClientMock.Do with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDoCounter) < 1 {
		if m.DoMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IHTTPClientMock.Do")
		} else {
			m.t.Errorf("Expected call to IHTTPClientMock.Do with params: %#v", *m.DoMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDo != nil && mm_atomic.LoadUint64(&m.afterDoCounter) < 1 {
		m.t.Error("Expected call to IHTTPClientMock.Do")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *IHTTPClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDoInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *IHTTPClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *IHTTPClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDoDone()
}
