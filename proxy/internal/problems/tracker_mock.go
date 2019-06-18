package problems

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"oss.indeed.com/go/modprox/pkg/coordinates"
)

// TrackerMock implements Tracker
type TrackerMock struct {
	t minimock.Tester

	funcProblem          func(module coordinates.Module) (p1 Problem, b1 bool)
	afterProblemCounter  uint64
	beforeProblemCounter uint64
	ProblemMock          mTrackerMockProblem

	funcProblems          func() (pa1 []Problem)
	afterProblemsCounter  uint64
	beforeProblemsCounter uint64
	ProblemsMock          mTrackerMockProblems

	funcSet          func(p1 Problem)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mTrackerMockSet
}

// NewTrackerMock returns a mock for Tracker
func NewTrackerMock(t minimock.Tester) *TrackerMock {
	m := &TrackerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}
	m.ProblemMock = mTrackerMockProblem{mock: m}
	m.ProblemsMock = mTrackerMockProblems{mock: m}
	m.SetMock = mTrackerMockSet{mock: m}

	return m
}

type mTrackerMockProblem struct {
	mock               *TrackerMock
	defaultExpectation *TrackerMockProblemExpectation
	expectations       []*TrackerMockProblemExpectation
}

// TrackerMockProblemExpectation specifies expectation struct of the Tracker.Problem
type TrackerMockProblemExpectation struct {
	mock    *TrackerMock
	params  *TrackerMockProblemParams
	results *TrackerMockProblemResults
	Counter uint64
}

// TrackerMockProblemParams contains parameters of the Tracker.Problem
type TrackerMockProblemParams struct {
	module coordinates.Module
}

// TrackerMockProblemResults contains results of the Tracker.Problem
type TrackerMockProblemResults struct {
	p1 Problem
	b1 bool
}

// Expect sets up expected params for Tracker.Problem
func (m *mTrackerMockProblem) Expect(module coordinates.Module) *mTrackerMockProblem {
	if m.mock.funcProblem != nil {
		m.mock.t.Fatalf("TrackerMock.Problem mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &TrackerMockProblemExpectation{}
	}

	m.defaultExpectation.params = &TrackerMockProblemParams{module}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Tracker.Problem
func (m *mTrackerMockProblem) Return(p1 Problem, b1 bool) *TrackerMock {
	if m.mock.funcProblem != nil {
		m.mock.t.Fatalf("TrackerMock.Problem mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &TrackerMockProblemExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &TrackerMockProblemResults{p1, b1}
	return m.mock
}

//Set uses given function f to mock the Tracker.Problem method
func (m *mTrackerMockProblem) Set(f func(module coordinates.Module) (p1 Problem, b1 bool)) *TrackerMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Tracker.Problem method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Tracker.Problem method")
	}

	m.mock.funcProblem = f
	return m.mock
}

// When sets expectation for the Tracker.Problem which will trigger the result defined by the following
// Then helper
func (m *mTrackerMockProblem) When(module coordinates.Module) *TrackerMockProblemExpectation {
	if m.mock.funcProblem != nil {
		m.mock.t.Fatalf("TrackerMock.Problem mock is already set by Set")
	}

	expectation := &TrackerMockProblemExpectation{
		mock:   m.mock,
		params: &TrackerMockProblemParams{module},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Tracker.Problem return parameters for the expectation previously defined by the When method
func (e *TrackerMockProblemExpectation) Then(p1 Problem, b1 bool) *TrackerMock {
	e.results = &TrackerMockProblemResults{p1, b1}
	return e.mock
}

// Problem implements Tracker
func (m *TrackerMock) Problem(module coordinates.Module) (p1 Problem, b1 bool) {
	mm_atomic.AddUint64(&m.beforeProblemCounter, 1)
	defer mm_atomic.AddUint64(&m.afterProblemCounter, 1)

	for _, e := range m.ProblemMock.expectations {
		if minimock.Equal(*e.params, TrackerMockProblemParams{module}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.b1
		}
	}

	if m.ProblemMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.ProblemMock.defaultExpectation.Counter, 1)
		want := m.ProblemMock.defaultExpectation.params
		got := TrackerMockProblemParams{module}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("TrackerMock.Problem got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.ProblemMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the TrackerMock.Problem")
		}
		return (*results).p1, (*results).b1
	}
	if m.funcProblem != nil {
		return m.funcProblem(module)
	}
	m.t.Fatalf("Unexpected call to TrackerMock.Problem. %v", module)
	return
}

// ProblemAfterCounter returns a count of finished TrackerMock.Problem invocations
func (m *TrackerMock) ProblemAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterProblemCounter)
}

// ProblemBeforeCounter returns a count of TrackerMock.Problem invocations
func (m *TrackerMock) ProblemBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeProblemCounter)
}

// MinimockProblemDone returns true if the count of the Problem invocations corresponds
// the number of defined expectations
func (m *TrackerMock) MinimockProblemDone() bool {
	for _, e := range m.ProblemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProblemMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProblemCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProblem != nil && mm_atomic.LoadUint64(&m.afterProblemCounter) < 1 {
		return false
	}
	return true
}

// MinimockProblemInspect logs each unmet expectation
func (m *TrackerMock) MinimockProblemInspect() {
	for _, e := range m.ProblemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TrackerMock.Problem with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProblemMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProblemCounter) < 1 {
		if m.ProblemMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TrackerMock.Problem")
		} else {
			m.t.Errorf("Expected call to TrackerMock.Problem with params: %#v", *m.ProblemMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProblem != nil && mm_atomic.LoadUint64(&m.afterProblemCounter) < 1 {
		m.t.Error("Expected call to TrackerMock.Problem")
	}
}

type mTrackerMockProblems struct {
	mock               *TrackerMock
	defaultExpectation *TrackerMockProblemsExpectation
	expectations       []*TrackerMockProblemsExpectation
}

// TrackerMockProblemsExpectation specifies expectation struct of the Tracker.Problems
type TrackerMockProblemsExpectation struct {
	mock *TrackerMock

	results *TrackerMockProblemsResults
	Counter uint64
}

// TrackerMockProblemsResults contains results of the Tracker.Problems
type TrackerMockProblemsResults struct {
	pa1 []Problem
}

// Expect sets up expected params for Tracker.Problems
func (m *mTrackerMockProblems) Expect() *mTrackerMockProblems {
	if m.mock.funcProblems != nil {
		m.mock.t.Fatalf("TrackerMock.Problems mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &TrackerMockProblemsExpectation{}
	}

	return m
}

// Return sets up results that will be returned by Tracker.Problems
func (m *mTrackerMockProblems) Return(pa1 []Problem) *TrackerMock {
	if m.mock.funcProblems != nil {
		m.mock.t.Fatalf("TrackerMock.Problems mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &TrackerMockProblemsExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &TrackerMockProblemsResults{pa1}
	return m.mock
}

//Set uses given function f to mock the Tracker.Problems method
func (m *mTrackerMockProblems) Set(f func() (pa1 []Problem)) *TrackerMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Tracker.Problems method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Tracker.Problems method")
	}

	m.mock.funcProblems = f
	return m.mock
}

// Problems implements Tracker
func (m *TrackerMock) Problems() (pa1 []Problem) {
	mm_atomic.AddUint64(&m.beforeProblemsCounter, 1)
	defer mm_atomic.AddUint64(&m.afterProblemsCounter, 1)

	if m.ProblemsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.ProblemsMock.defaultExpectation.Counter, 1)

		results := m.ProblemsMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the TrackerMock.Problems")
		}
		return (*results).pa1
	}
	if m.funcProblems != nil {
		return m.funcProblems()
	}
	m.t.Fatalf("Unexpected call to TrackerMock.Problems.")
	return
}

// ProblemsAfterCounter returns a count of finished TrackerMock.Problems invocations
func (m *TrackerMock) ProblemsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterProblemsCounter)
}

// ProblemsBeforeCounter returns a count of TrackerMock.Problems invocations
func (m *TrackerMock) ProblemsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeProblemsCounter)
}

// MinimockProblemsDone returns true if the count of the Problems invocations corresponds
// the number of defined expectations
func (m *TrackerMock) MinimockProblemsDone() bool {
	for _, e := range m.ProblemsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProblemsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProblemsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProblems != nil && mm_atomic.LoadUint64(&m.afterProblemsCounter) < 1 {
		return false
	}
	return true
}

// MinimockProblemsInspect logs each unmet expectation
func (m *TrackerMock) MinimockProblemsInspect() {
	for _, e := range m.ProblemsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to TrackerMock.Problems")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProblemsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProblemsCounter) < 1 {
		m.t.Error("Expected call to TrackerMock.Problems")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProblems != nil && mm_atomic.LoadUint64(&m.afterProblemsCounter) < 1 {
		m.t.Error("Expected call to TrackerMock.Problems")
	}
}

type mTrackerMockSet struct {
	mock               *TrackerMock
	defaultExpectation *TrackerMockSetExpectation
	expectations       []*TrackerMockSetExpectation
}

// TrackerMockSetExpectation specifies expectation struct of the Tracker.Set
type TrackerMockSetExpectation struct {
	mock   *TrackerMock
	params *TrackerMockSetParams

	Counter uint64
}

// TrackerMockSetParams contains parameters of the Tracker.Set
type TrackerMockSetParams struct {
	p1 Problem
}

// Expect sets up expected params for Tracker.Set
func (m *mTrackerMockSet) Expect(p1 Problem) *mTrackerMockSet {
	if m.mock.funcSet != nil {
		m.mock.t.Fatalf("TrackerMock.Set mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &TrackerMockSetExpectation{}
	}

	m.defaultExpectation.params = &TrackerMockSetParams{p1}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Tracker.Set
func (m *mTrackerMockSet) Return() *TrackerMock {
	if m.mock.funcSet != nil {
		m.mock.t.Fatalf("TrackerMock.Set mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &TrackerMockSetExpectation{mock: m.mock}
	}

	return m.mock
}

//Set uses given function f to mock the Tracker.Set method
func (m *mTrackerMockSet) Set(f func(p1 Problem)) *TrackerMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Tracker.Set method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Tracker.Set method")
	}

	m.mock.funcSet = f
	return m.mock
}

// Set implements Tracker
func (m *TrackerMock) Set(p1 Problem) {
	mm_atomic.AddUint64(&m.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&m.afterSetCounter, 1)

	for _, e := range m.SetMock.expectations {
		if minimock.Equal(*e.params, TrackerMockSetParams{p1}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if m.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.SetMock.defaultExpectation.Counter, 1)
		want := m.SetMock.defaultExpectation.params
		got := TrackerMockSetParams{p1}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("TrackerMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if m.funcSet != nil {
		m.funcSet(p1)
		return
	}
	m.t.Fatalf("Unexpected call to TrackerMock.Set. %v", p1)

}

// SetAfterCounter returns a count of finished TrackerMock.Set invocations
func (m *TrackerMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterSetCounter)
}

// SetBeforeCounter returns a count of TrackerMock.Set invocations
func (m *TrackerMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeSetCounter)
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *TrackerMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *TrackerMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TrackerMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TrackerMock.Set")
		} else {
			m.t.Errorf("Expected call to TrackerMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to TrackerMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TrackerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockProblemInspect()

		m.MinimockProblemsInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TrackerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TrackerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockProblemDone() &&
		m.MinimockProblemsDone() &&
		m.MinimockSetDone()
}
