package proxies

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// PrunerMock implements Pruner
type PrunerMock struct {
	t minimock.Tester

	funcPrune          func(t1 time.Time) (err error)
	inspectFuncPrune   func(t1 time.Time)
	afterPruneCounter  uint64
	beforePruneCounter uint64
	PruneMock          mPrunerMockPrune
}

// NewPrunerMock returns a mock for Pruner
func NewPrunerMock(t minimock.Tester) *PrunerMock {
	m := &PrunerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.PruneMock = mPrunerMockPrune{mock: m}
	m.PruneMock.callArgs = []*PrunerMockPruneParams{}

	return m
}

type mPrunerMockPrune struct {
	mock               *PrunerMock
	defaultExpectation *PrunerMockPruneExpectation
	expectations       []*PrunerMockPruneExpectation

	callArgs []*PrunerMockPruneParams
	mutex    sync.RWMutex
}

// PrunerMockPruneExpectation specifies expectation struct of the Pruner.Prune
type PrunerMockPruneExpectation struct {
	mock    *PrunerMock
	params  *PrunerMockPruneParams
	results *PrunerMockPruneResults
	Counter uint64
}

// PrunerMockPruneParams contains parameters of the Pruner.Prune
type PrunerMockPruneParams struct {
	t1 time.Time
}

// PrunerMockPruneResults contains results of the Pruner.Prune
type PrunerMockPruneResults struct {
	err error
}

// Expect sets up expected params for Pruner.Prune
func (mmPrune *mPrunerMockPrune) Expect(t1 time.Time) *mPrunerMockPrune {
	if mmPrune.mock.funcPrune != nil {
		mmPrune.mock.t.Fatalf("PrunerMock.Prune mock is already set by Set")
	}

	if mmPrune.defaultExpectation == nil {
		mmPrune.defaultExpectation = &PrunerMockPruneExpectation{}
	}

	mmPrune.defaultExpectation.params = &PrunerMockPruneParams{t1}
	for _, e := range mmPrune.expectations {
		if minimock.Equal(e.params, mmPrune.defaultExpectation.params) {
			mmPrune.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPrune.defaultExpectation.params)
		}
	}

	return mmPrune
}

// Inspect accepts an inspector function that has same arguments as the Pruner.Prune
func (mmPrune *mPrunerMockPrune) Inspect(f func(t1 time.Time)) *mPrunerMockPrune {
	if mmPrune.mock.inspectFuncPrune != nil {
		mmPrune.mock.t.Fatalf("Inspect function is already set for PrunerMock.Prune")
	}

	mmPrune.mock.inspectFuncPrune = f

	return mmPrune
}

// Return sets up results that will be returned by Pruner.Prune
func (mmPrune *mPrunerMockPrune) Return(err error) *PrunerMock {
	if mmPrune.mock.funcPrune != nil {
		mmPrune.mock.t.Fatalf("PrunerMock.Prune mock is already set by Set")
	}

	if mmPrune.defaultExpectation == nil {
		mmPrune.defaultExpectation = &PrunerMockPruneExpectation{mock: mmPrune.mock}
	}
	mmPrune.defaultExpectation.results = &PrunerMockPruneResults{err}
	return mmPrune.mock
}

//Set uses given function f to mock the Pruner.Prune method
func (mmPrune *mPrunerMockPrune) Set(f func(t1 time.Time) (err error)) *PrunerMock {
	if mmPrune.defaultExpectation != nil {
		mmPrune.mock.t.Fatalf("Default expectation is already set for the Pruner.Prune method")
	}

	if len(mmPrune.expectations) > 0 {
		mmPrune.mock.t.Fatalf("Some expectations are already set for the Pruner.Prune method")
	}

	mmPrune.mock.funcPrune = f
	return mmPrune.mock
}

// When sets expectation for the Pruner.Prune which will trigger the result defined by the following
// Then helper
func (mmPrune *mPrunerMockPrune) When(t1 time.Time) *PrunerMockPruneExpectation {
	if mmPrune.mock.funcPrune != nil {
		mmPrune.mock.t.Fatalf("PrunerMock.Prune mock is already set by Set")
	}

	expectation := &PrunerMockPruneExpectation{
		mock:   mmPrune.mock,
		params: &PrunerMockPruneParams{t1},
	}
	mmPrune.expectations = append(mmPrune.expectations, expectation)
	return expectation
}

// Then sets up Pruner.Prune return parameters for the expectation previously defined by the When method
func (e *PrunerMockPruneExpectation) Then(err error) *PrunerMock {
	e.results = &PrunerMockPruneResults{err}
	return e.mock
}

// Prune implements Pruner
func (mmPrune *PrunerMock) Prune(t1 time.Time) (err error) {
	mm_atomic.AddUint64(&mmPrune.beforePruneCounter, 1)
	defer mm_atomic.AddUint64(&mmPrune.afterPruneCounter, 1)

	if mmPrune.inspectFuncPrune != nil {
		mmPrune.inspectFuncPrune(t1)
	}

	mm_params := &PrunerMockPruneParams{t1}

	// Record call args
	mmPrune.PruneMock.mutex.Lock()
	mmPrune.PruneMock.callArgs = append(mmPrune.PruneMock.callArgs, mm_params)
	mmPrune.PruneMock.mutex.Unlock()

	for _, e := range mmPrune.PruneMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPrune.PruneMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPrune.PruneMock.defaultExpectation.Counter, 1)
		mm_want := mmPrune.PruneMock.defaultExpectation.params
		mm_got := PrunerMockPruneParams{t1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPrune.t.Errorf("PrunerMock.Prune got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPrune.PruneMock.defaultExpectation.results
		if mm_results == nil {
			mmPrune.t.Fatal("No results are set for the PrunerMock.Prune")
		}
		return (*mm_results).err
	}
	if mmPrune.funcPrune != nil {
		return mmPrune.funcPrune(t1)
	}
	mmPrune.t.Fatalf("Unexpected call to PrunerMock.Prune. %v", t1)
	return
}

// PruneAfterCounter returns a count of finished PrunerMock.Prune invocations
func (mmPrune *PrunerMock) PruneAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPrune.afterPruneCounter)
}

// PruneBeforeCounter returns a count of PrunerMock.Prune invocations
func (mmPrune *PrunerMock) PruneBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPrune.beforePruneCounter)
}

// Calls returns a list of arguments used in each call to PrunerMock.Prune.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPrune *mPrunerMockPrune) Calls() []*PrunerMockPruneParams {
	mmPrune.mutex.RLock()

	argCopy := make([]*PrunerMockPruneParams, len(mmPrune.callArgs))
	copy(argCopy, mmPrune.callArgs)

	mmPrune.mutex.RUnlock()

	return argCopy
}

// MinimockPruneDone returns true if the count of the Prune invocations corresponds
// the number of defined expectations
func (m *PrunerMock) MinimockPruneDone() bool {
	for _, e := range m.PruneMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PruneMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPruneCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPrune != nil && mm_atomic.LoadUint64(&m.afterPruneCounter) < 1 {
		return false
	}
	return true
}

// MinimockPruneInspect logs each unmet expectation
func (m *PrunerMock) MinimockPruneInspect() {
	for _, e := range m.PruneMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PrunerMock.Prune with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PruneMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPruneCounter) < 1 {
		if m.PruneMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PrunerMock.Prune")
		} else {
			m.t.Errorf("Expected call to PrunerMock.Prune with params: %#v", *m.PruneMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPrune != nil && mm_atomic.LoadUint64(&m.afterPruneCounter) < 1 {
		m.t.Error("Expected call to PrunerMock.Prune")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PrunerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockPruneInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PrunerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *PrunerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockPruneDone()
}
