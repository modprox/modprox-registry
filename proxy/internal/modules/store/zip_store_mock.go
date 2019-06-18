package store

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"oss.indeed.com/go/modprox/pkg/coordinates"
	"oss.indeed.com/go/modprox/pkg/repository"
)

// ZipStoreMock implements ZipStore
type ZipStoreMock struct {
	t minimock.Tester

	funcDelZip          func(m1 coordinates.Module) (err error)
	afterDelZipCounter  uint64
	beforeDelZipCounter uint64
	DelZipMock          mZipStoreMockDelZip

	funcGetZip          func(m1 coordinates.Module) (b1 repository.Blob, err error)
	afterGetZipCounter  uint64
	beforeGetZipCounter uint64
	GetZipMock          mZipStoreMockGetZip

	funcPutZip          func(m1 coordinates.Module, b1 repository.Blob) (err error)
	afterPutZipCounter  uint64
	beforePutZipCounter uint64
	PutZipMock          mZipStoreMockPutZip
}

// NewZipStoreMock returns a mock for ZipStore
func NewZipStoreMock(t minimock.Tester) *ZipStoreMock {
	m := &ZipStoreMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}
	m.DelZipMock = mZipStoreMockDelZip{mock: m}
	m.GetZipMock = mZipStoreMockGetZip{mock: m}
	m.PutZipMock = mZipStoreMockPutZip{mock: m}

	return m
}

type mZipStoreMockDelZip struct {
	mock               *ZipStoreMock
	defaultExpectation *ZipStoreMockDelZipExpectation
	expectations       []*ZipStoreMockDelZipExpectation
}

// ZipStoreMockDelZipExpectation specifies expectation struct of the ZipStore.DelZip
type ZipStoreMockDelZipExpectation struct {
	mock    *ZipStoreMock
	params  *ZipStoreMockDelZipParams
	results *ZipStoreMockDelZipResults
	Counter uint64
}

// ZipStoreMockDelZipParams contains parameters of the ZipStore.DelZip
type ZipStoreMockDelZipParams struct {
	m1 coordinates.Module
}

// ZipStoreMockDelZipResults contains results of the ZipStore.DelZip
type ZipStoreMockDelZipResults struct {
	err error
}

// Expect sets up expected params for ZipStore.DelZip
func (m *mZipStoreMockDelZip) Expect(m1 coordinates.Module) *mZipStoreMockDelZip {
	if m.mock.funcDelZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.DelZip mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &ZipStoreMockDelZipExpectation{}
	}

	m.defaultExpectation.params = &ZipStoreMockDelZipParams{m1}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by ZipStore.DelZip
func (m *mZipStoreMockDelZip) Return(err error) *ZipStoreMock {
	if m.mock.funcDelZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.DelZip mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &ZipStoreMockDelZipExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &ZipStoreMockDelZipResults{err}
	return m.mock
}

//Set uses given function f to mock the ZipStore.DelZip method
func (m *mZipStoreMockDelZip) Set(f func(m1 coordinates.Module) (err error)) *ZipStoreMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the ZipStore.DelZip method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the ZipStore.DelZip method")
	}

	m.mock.funcDelZip = f
	return m.mock
}

// When sets expectation for the ZipStore.DelZip which will trigger the result defined by the following
// Then helper
func (m *mZipStoreMockDelZip) When(m1 coordinates.Module) *ZipStoreMockDelZipExpectation {
	if m.mock.funcDelZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.DelZip mock is already set by Set")
	}

	expectation := &ZipStoreMockDelZipExpectation{
		mock:   m.mock,
		params: &ZipStoreMockDelZipParams{m1},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up ZipStore.DelZip return parameters for the expectation previously defined by the When method
func (e *ZipStoreMockDelZipExpectation) Then(err error) *ZipStoreMock {
	e.results = &ZipStoreMockDelZipResults{err}
	return e.mock
}

// DelZip implements ZipStore
func (m *ZipStoreMock) DelZip(m1 coordinates.Module) (err error) {
	mm_atomic.AddUint64(&m.beforeDelZipCounter, 1)
	defer mm_atomic.AddUint64(&m.afterDelZipCounter, 1)

	for _, e := range m.DelZipMock.expectations {
		if minimock.Equal(*e.params, ZipStoreMockDelZipParams{m1}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.DelZipMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.DelZipMock.defaultExpectation.Counter, 1)
		want := m.DelZipMock.defaultExpectation.params
		got := ZipStoreMockDelZipParams{m1}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("ZipStoreMock.DelZip got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.DelZipMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the ZipStoreMock.DelZip")
		}
		return (*results).err
	}
	if m.funcDelZip != nil {
		return m.funcDelZip(m1)
	}
	m.t.Fatalf("Unexpected call to ZipStoreMock.DelZip. %v", m1)
	return
}

// DelZipAfterCounter returns a count of finished ZipStoreMock.DelZip invocations
func (m *ZipStoreMock) DelZipAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterDelZipCounter)
}

// DelZipBeforeCounter returns a count of ZipStoreMock.DelZip invocations
func (m *ZipStoreMock) DelZipBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeDelZipCounter)
}

// MinimockDelZipDone returns true if the count of the DelZip invocations corresponds
// the number of defined expectations
func (m *ZipStoreMock) MinimockDelZipDone() bool {
	for _, e := range m.DelZipMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DelZipMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDelZipCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelZip != nil && mm_atomic.LoadUint64(&m.afterDelZipCounter) < 1 {
		return false
	}
	return true
}

// MinimockDelZipInspect logs each unmet expectation
func (m *ZipStoreMock) MinimockDelZipInspect() {
	for _, e := range m.DelZipMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ZipStoreMock.DelZip with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DelZipMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDelZipCounter) < 1 {
		if m.DelZipMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ZipStoreMock.DelZip")
		} else {
			m.t.Errorf("Expected call to ZipStoreMock.DelZip with params: %#v", *m.DelZipMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelZip != nil && mm_atomic.LoadUint64(&m.afterDelZipCounter) < 1 {
		m.t.Error("Expected call to ZipStoreMock.DelZip")
	}
}

type mZipStoreMockGetZip struct {
	mock               *ZipStoreMock
	defaultExpectation *ZipStoreMockGetZipExpectation
	expectations       []*ZipStoreMockGetZipExpectation
}

// ZipStoreMockGetZipExpectation specifies expectation struct of the ZipStore.GetZip
type ZipStoreMockGetZipExpectation struct {
	mock    *ZipStoreMock
	params  *ZipStoreMockGetZipParams
	results *ZipStoreMockGetZipResults
	Counter uint64
}

// ZipStoreMockGetZipParams contains parameters of the ZipStore.GetZip
type ZipStoreMockGetZipParams struct {
	m1 coordinates.Module
}

// ZipStoreMockGetZipResults contains results of the ZipStore.GetZip
type ZipStoreMockGetZipResults struct {
	b1  repository.Blob
	err error
}

// Expect sets up expected params for ZipStore.GetZip
func (m *mZipStoreMockGetZip) Expect(m1 coordinates.Module) *mZipStoreMockGetZip {
	if m.mock.funcGetZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.GetZip mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &ZipStoreMockGetZipExpectation{}
	}

	m.defaultExpectation.params = &ZipStoreMockGetZipParams{m1}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by ZipStore.GetZip
func (m *mZipStoreMockGetZip) Return(b1 repository.Blob, err error) *ZipStoreMock {
	if m.mock.funcGetZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.GetZip mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &ZipStoreMockGetZipExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &ZipStoreMockGetZipResults{b1, err}
	return m.mock
}

//Set uses given function f to mock the ZipStore.GetZip method
func (m *mZipStoreMockGetZip) Set(f func(m1 coordinates.Module) (b1 repository.Blob, err error)) *ZipStoreMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the ZipStore.GetZip method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the ZipStore.GetZip method")
	}

	m.mock.funcGetZip = f
	return m.mock
}

// When sets expectation for the ZipStore.GetZip which will trigger the result defined by the following
// Then helper
func (m *mZipStoreMockGetZip) When(m1 coordinates.Module) *ZipStoreMockGetZipExpectation {
	if m.mock.funcGetZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.GetZip mock is already set by Set")
	}

	expectation := &ZipStoreMockGetZipExpectation{
		mock:   m.mock,
		params: &ZipStoreMockGetZipParams{m1},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up ZipStore.GetZip return parameters for the expectation previously defined by the When method
func (e *ZipStoreMockGetZipExpectation) Then(b1 repository.Blob, err error) *ZipStoreMock {
	e.results = &ZipStoreMockGetZipResults{b1, err}
	return e.mock
}

// GetZip implements ZipStore
func (m *ZipStoreMock) GetZip(m1 coordinates.Module) (b1 repository.Blob, err error) {
	mm_atomic.AddUint64(&m.beforeGetZipCounter, 1)
	defer mm_atomic.AddUint64(&m.afterGetZipCounter, 1)

	for _, e := range m.GetZipMock.expectations {
		if minimock.Equal(*e.params, ZipStoreMockGetZipParams{m1}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if m.GetZipMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.GetZipMock.defaultExpectation.Counter, 1)
		want := m.GetZipMock.defaultExpectation.params
		got := ZipStoreMockGetZipParams{m1}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("ZipStoreMock.GetZip got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.GetZipMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the ZipStoreMock.GetZip")
		}
		return (*results).b1, (*results).err
	}
	if m.funcGetZip != nil {
		return m.funcGetZip(m1)
	}
	m.t.Fatalf("Unexpected call to ZipStoreMock.GetZip. %v", m1)
	return
}

// GetZipAfterCounter returns a count of finished ZipStoreMock.GetZip invocations
func (m *ZipStoreMock) GetZipAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterGetZipCounter)
}

// GetZipBeforeCounter returns a count of ZipStoreMock.GetZip invocations
func (m *ZipStoreMock) GetZipBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeGetZipCounter)
}

// MinimockGetZipDone returns true if the count of the GetZip invocations corresponds
// the number of defined expectations
func (m *ZipStoreMock) MinimockGetZipDone() bool {
	for _, e := range m.GetZipMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetZipMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetZipCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetZip != nil && mm_atomic.LoadUint64(&m.afterGetZipCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetZipInspect logs each unmet expectation
func (m *ZipStoreMock) MinimockGetZipInspect() {
	for _, e := range m.GetZipMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ZipStoreMock.GetZip with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetZipMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetZipCounter) < 1 {
		if m.GetZipMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ZipStoreMock.GetZip")
		} else {
			m.t.Errorf("Expected call to ZipStoreMock.GetZip with params: %#v", *m.GetZipMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetZip != nil && mm_atomic.LoadUint64(&m.afterGetZipCounter) < 1 {
		m.t.Error("Expected call to ZipStoreMock.GetZip")
	}
}

type mZipStoreMockPutZip struct {
	mock               *ZipStoreMock
	defaultExpectation *ZipStoreMockPutZipExpectation
	expectations       []*ZipStoreMockPutZipExpectation
}

// ZipStoreMockPutZipExpectation specifies expectation struct of the ZipStore.PutZip
type ZipStoreMockPutZipExpectation struct {
	mock    *ZipStoreMock
	params  *ZipStoreMockPutZipParams
	results *ZipStoreMockPutZipResults
	Counter uint64
}

// ZipStoreMockPutZipParams contains parameters of the ZipStore.PutZip
type ZipStoreMockPutZipParams struct {
	m1 coordinates.Module
	b1 repository.Blob
}

// ZipStoreMockPutZipResults contains results of the ZipStore.PutZip
type ZipStoreMockPutZipResults struct {
	err error
}

// Expect sets up expected params for ZipStore.PutZip
func (m *mZipStoreMockPutZip) Expect(m1 coordinates.Module, b1 repository.Blob) *mZipStoreMockPutZip {
	if m.mock.funcPutZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.PutZip mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &ZipStoreMockPutZipExpectation{}
	}

	m.defaultExpectation.params = &ZipStoreMockPutZipParams{m1, b1}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by ZipStore.PutZip
func (m *mZipStoreMockPutZip) Return(err error) *ZipStoreMock {
	if m.mock.funcPutZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.PutZip mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &ZipStoreMockPutZipExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &ZipStoreMockPutZipResults{err}
	return m.mock
}

//Set uses given function f to mock the ZipStore.PutZip method
func (m *mZipStoreMockPutZip) Set(f func(m1 coordinates.Module, b1 repository.Blob) (err error)) *ZipStoreMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the ZipStore.PutZip method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the ZipStore.PutZip method")
	}

	m.mock.funcPutZip = f
	return m.mock
}

// When sets expectation for the ZipStore.PutZip which will trigger the result defined by the following
// Then helper
func (m *mZipStoreMockPutZip) When(m1 coordinates.Module, b1 repository.Blob) *ZipStoreMockPutZipExpectation {
	if m.mock.funcPutZip != nil {
		m.mock.t.Fatalf("ZipStoreMock.PutZip mock is already set by Set")
	}

	expectation := &ZipStoreMockPutZipExpectation{
		mock:   m.mock,
		params: &ZipStoreMockPutZipParams{m1, b1},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up ZipStore.PutZip return parameters for the expectation previously defined by the When method
func (e *ZipStoreMockPutZipExpectation) Then(err error) *ZipStoreMock {
	e.results = &ZipStoreMockPutZipResults{err}
	return e.mock
}

// PutZip implements ZipStore
func (m *ZipStoreMock) PutZip(m1 coordinates.Module, b1 repository.Blob) (err error) {
	mm_atomic.AddUint64(&m.beforePutZipCounter, 1)
	defer mm_atomic.AddUint64(&m.afterPutZipCounter, 1)

	for _, e := range m.PutZipMock.expectations {
		if minimock.Equal(*e.params, ZipStoreMockPutZipParams{m1, b1}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.PutZipMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.PutZipMock.defaultExpectation.Counter, 1)
		want := m.PutZipMock.defaultExpectation.params
		got := ZipStoreMockPutZipParams{m1, b1}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("ZipStoreMock.PutZip got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.PutZipMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the ZipStoreMock.PutZip")
		}
		return (*results).err
	}
	if m.funcPutZip != nil {
		return m.funcPutZip(m1, b1)
	}
	m.t.Fatalf("Unexpected call to ZipStoreMock.PutZip. %v %v", m1, b1)
	return
}

// PutZipAfterCounter returns a count of finished ZipStoreMock.PutZip invocations
func (m *ZipStoreMock) PutZipAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterPutZipCounter)
}

// PutZipBeforeCounter returns a count of ZipStoreMock.PutZip invocations
func (m *ZipStoreMock) PutZipBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforePutZipCounter)
}

// MinimockPutZipDone returns true if the count of the PutZip invocations corresponds
// the number of defined expectations
func (m *ZipStoreMock) MinimockPutZipDone() bool {
	for _, e := range m.PutZipMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PutZipMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPutZipCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPutZip != nil && mm_atomic.LoadUint64(&m.afterPutZipCounter) < 1 {
		return false
	}
	return true
}

// MinimockPutZipInspect logs each unmet expectation
func (m *ZipStoreMock) MinimockPutZipInspect() {
	for _, e := range m.PutZipMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ZipStoreMock.PutZip with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PutZipMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPutZipCounter) < 1 {
		if m.PutZipMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ZipStoreMock.PutZip")
		} else {
			m.t.Errorf("Expected call to ZipStoreMock.PutZip with params: %#v", *m.PutZipMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPutZip != nil && mm_atomic.LoadUint64(&m.afterPutZipCounter) < 1 {
		m.t.Error("Expected call to ZipStoreMock.PutZip")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ZipStoreMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDelZipInspect()

		m.MinimockGetZipInspect()

		m.MinimockPutZipInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ZipStoreMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ZipStoreMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDelZipDone() &&
		m.MinimockGetZipDone() &&
		m.MinimockPutZipDone()
}
