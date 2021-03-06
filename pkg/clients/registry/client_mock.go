package registry

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"io"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ClientMock implements Client
type ClientMock struct {
	t minimock.Tester

	funcGet          func(path string, rw io.Writer) (err error)
	inspectFuncGet   func(path string, rw io.Writer)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mClientMockGet

	funcPost          func(path string, body io.Reader, rw io.Writer) (err error)
	inspectFuncPost   func(path string, body io.Reader, rw io.Writer)
	afterPostCounter  uint64
	beforePostCounter uint64
	PostMock          mClientMockPost
}

// NewClientMock returns a mock for Client
func NewClientMock(t minimock.Tester) *ClientMock {
	m := &ClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetMock = mClientMockGet{mock: m}
	m.GetMock.callArgs = []*ClientMockGetParams{}

	m.PostMock = mClientMockPost{mock: m}
	m.PostMock.callArgs = []*ClientMockPostParams{}

	return m
}

type mClientMockGet struct {
	mock               *ClientMock
	defaultExpectation *ClientMockGetExpectation
	expectations       []*ClientMockGetExpectation

	callArgs []*ClientMockGetParams
	mutex    sync.RWMutex
}

// ClientMockGetExpectation specifies expectation struct of the Client.Get
type ClientMockGetExpectation struct {
	mock    *ClientMock
	params  *ClientMockGetParams
	results *ClientMockGetResults
	Counter uint64
}

// ClientMockGetParams contains parameters of the Client.Get
type ClientMockGetParams struct {
	path string
	rw   io.Writer
}

// ClientMockGetResults contains results of the Client.Get
type ClientMockGetResults struct {
	err error
}

// Expect sets up expected params for Client.Get
func (mmGet *mClientMockGet) Expect(path string, rw io.Writer) *mClientMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &ClientMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &ClientMockGetParams{path, rw}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Client.Get
func (mmGet *mClientMockGet) Inspect(f func(path string, rw io.Writer)) *mClientMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for ClientMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Client.Get
func (mmGet *mClientMockGet) Return(err error) *ClientMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &ClientMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &ClientMockGetResults{err}
	return mmGet.mock
}

//Set uses given function f to mock the Client.Get method
func (mmGet *mClientMockGet) Set(f func(path string, rw io.Writer) (err error)) *ClientMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Client.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Client.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Client.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mClientMockGet) When(path string, rw io.Writer) *ClientMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ClientMock.Get mock is already set by Set")
	}

	expectation := &ClientMockGetExpectation{
		mock:   mmGet.mock,
		params: &ClientMockGetParams{path, rw},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Client.Get return parameters for the expectation previously defined by the When method
func (e *ClientMockGetExpectation) Then(err error) *ClientMock {
	e.results = &ClientMockGetResults{err}
	return e.mock
}

// Get implements Client
func (mmGet *ClientMock) Get(path string, rw io.Writer) (err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(path, rw)
	}

	mm_params := &ClientMockGetParams{path, rw}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := ClientMockGetParams{path, rw}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("ClientMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the ClientMock.Get")
		}
		return (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(path, rw)
	}
	mmGet.t.Fatalf("Unexpected call to ClientMock.Get. %v %v", path, rw)
	return
}

// GetAfterCounter returns a count of finished ClientMock.Get invocations
func (mmGet *ClientMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of ClientMock.Get invocations
func (mmGet *ClientMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to ClientMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mClientMockGet) Calls() []*ClientMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*ClientMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockGetDone() bool {
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
func (m *ClientMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMock.Get")
		} else {
			m.t.Errorf("Expected call to ClientMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to ClientMock.Get")
	}
}

type mClientMockPost struct {
	mock               *ClientMock
	defaultExpectation *ClientMockPostExpectation
	expectations       []*ClientMockPostExpectation

	callArgs []*ClientMockPostParams
	mutex    sync.RWMutex
}

// ClientMockPostExpectation specifies expectation struct of the Client.Post
type ClientMockPostExpectation struct {
	mock    *ClientMock
	params  *ClientMockPostParams
	results *ClientMockPostResults
	Counter uint64
}

// ClientMockPostParams contains parameters of the Client.Post
type ClientMockPostParams struct {
	path string
	body io.Reader
	rw   io.Writer
}

// ClientMockPostResults contains results of the Client.Post
type ClientMockPostResults struct {
	err error
}

// Expect sets up expected params for Client.Post
func (mmPost *mClientMockPost) Expect(path string, body io.Reader, rw io.Writer) *mClientMockPost {
	if mmPost.mock.funcPost != nil {
		mmPost.mock.t.Fatalf("ClientMock.Post mock is already set by Set")
	}

	if mmPost.defaultExpectation == nil {
		mmPost.defaultExpectation = &ClientMockPostExpectation{}
	}

	mmPost.defaultExpectation.params = &ClientMockPostParams{path, body, rw}
	for _, e := range mmPost.expectations {
		if minimock.Equal(e.params, mmPost.defaultExpectation.params) {
			mmPost.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPost.defaultExpectation.params)
		}
	}

	return mmPost
}

// Inspect accepts an inspector function that has same arguments as the Client.Post
func (mmPost *mClientMockPost) Inspect(f func(path string, body io.Reader, rw io.Writer)) *mClientMockPost {
	if mmPost.mock.inspectFuncPost != nil {
		mmPost.mock.t.Fatalf("Inspect function is already set for ClientMock.Post")
	}

	mmPost.mock.inspectFuncPost = f

	return mmPost
}

// Return sets up results that will be returned by Client.Post
func (mmPost *mClientMockPost) Return(err error) *ClientMock {
	if mmPost.mock.funcPost != nil {
		mmPost.mock.t.Fatalf("ClientMock.Post mock is already set by Set")
	}

	if mmPost.defaultExpectation == nil {
		mmPost.defaultExpectation = &ClientMockPostExpectation{mock: mmPost.mock}
	}
	mmPost.defaultExpectation.results = &ClientMockPostResults{err}
	return mmPost.mock
}

//Set uses given function f to mock the Client.Post method
func (mmPost *mClientMockPost) Set(f func(path string, body io.Reader, rw io.Writer) (err error)) *ClientMock {
	if mmPost.defaultExpectation != nil {
		mmPost.mock.t.Fatalf("Default expectation is already set for the Client.Post method")
	}

	if len(mmPost.expectations) > 0 {
		mmPost.mock.t.Fatalf("Some expectations are already set for the Client.Post method")
	}

	mmPost.mock.funcPost = f
	return mmPost.mock
}

// When sets expectation for the Client.Post which will trigger the result defined by the following
// Then helper
func (mmPost *mClientMockPost) When(path string, body io.Reader, rw io.Writer) *ClientMockPostExpectation {
	if mmPost.mock.funcPost != nil {
		mmPost.mock.t.Fatalf("ClientMock.Post mock is already set by Set")
	}

	expectation := &ClientMockPostExpectation{
		mock:   mmPost.mock,
		params: &ClientMockPostParams{path, body, rw},
	}
	mmPost.expectations = append(mmPost.expectations, expectation)
	return expectation
}

// Then sets up Client.Post return parameters for the expectation previously defined by the When method
func (e *ClientMockPostExpectation) Then(err error) *ClientMock {
	e.results = &ClientMockPostResults{err}
	return e.mock
}

// Post implements Client
func (mmPost *ClientMock) Post(path string, body io.Reader, rw io.Writer) (err error) {
	mm_atomic.AddUint64(&mmPost.beforePostCounter, 1)
	defer mm_atomic.AddUint64(&mmPost.afterPostCounter, 1)

	if mmPost.inspectFuncPost != nil {
		mmPost.inspectFuncPost(path, body, rw)
	}

	mm_params := &ClientMockPostParams{path, body, rw}

	// Record call args
	mmPost.PostMock.mutex.Lock()
	mmPost.PostMock.callArgs = append(mmPost.PostMock.callArgs, mm_params)
	mmPost.PostMock.mutex.Unlock()

	for _, e := range mmPost.PostMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPost.PostMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPost.PostMock.defaultExpectation.Counter, 1)
		mm_want := mmPost.PostMock.defaultExpectation.params
		mm_got := ClientMockPostParams{path, body, rw}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPost.t.Errorf("ClientMock.Post got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPost.PostMock.defaultExpectation.results
		if mm_results == nil {
			mmPost.t.Fatal("No results are set for the ClientMock.Post")
		}
		return (*mm_results).err
	}
	if mmPost.funcPost != nil {
		return mmPost.funcPost(path, body, rw)
	}
	mmPost.t.Fatalf("Unexpected call to ClientMock.Post. %v %v %v", path, body, rw)
	return
}

// PostAfterCounter returns a count of finished ClientMock.Post invocations
func (mmPost *ClientMock) PostAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPost.afterPostCounter)
}

// PostBeforeCounter returns a count of ClientMock.Post invocations
func (mmPost *ClientMock) PostBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPost.beforePostCounter)
}

// Calls returns a list of arguments used in each call to ClientMock.Post.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPost *mClientMockPost) Calls() []*ClientMockPostParams {
	mmPost.mutex.RLock()

	argCopy := make([]*ClientMockPostParams, len(mmPost.callArgs))
	copy(argCopy, mmPost.callArgs)

	mmPost.mutex.RUnlock()

	return argCopy
}

// MinimockPostDone returns true if the count of the Post invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockPostDone() bool {
	for _, e := range m.PostMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PostMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPost != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		return false
	}
	return true
}

// MinimockPostInspect logs each unmet expectation
func (m *ClientMock) MinimockPostInspect() {
	for _, e := range m.PostMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMock.Post with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PostMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		if m.PostMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMock.Post")
		} else {
			m.t.Errorf("Expected call to ClientMock.Post with params: %#v", *m.PostMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPost != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		m.t.Error("Expected call to ClientMock.Post")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetInspect()

		m.MinimockPostInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetDone() &&
		m.MinimockPostDone()
}
