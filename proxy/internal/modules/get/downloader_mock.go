package get

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"oss.indeed.com/go/modprox/pkg/coordinates"
)

// DownloaderMock implements Downloader
type DownloaderMock struct {
	t minimock.Tester

	funcDownload          func(module coordinates.SerialModule) (err error)
	inspectFuncDownload   func(module coordinates.SerialModule)
	afterDownloadCounter  uint64
	beforeDownloadCounter uint64
	DownloadMock          mDownloaderMockDownload
}

// NewDownloaderMock returns a mock for Downloader
func NewDownloaderMock(t minimock.Tester) *DownloaderMock {
	m := &DownloaderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DownloadMock = mDownloaderMockDownload{mock: m}
	m.DownloadMock.callArgs = []*DownloaderMockDownloadParams{}

	return m
}

type mDownloaderMockDownload struct {
	mock               *DownloaderMock
	defaultExpectation *DownloaderMockDownloadExpectation
	expectations       []*DownloaderMockDownloadExpectation

	callArgs []*DownloaderMockDownloadParams
	mutex    sync.RWMutex
}

// DownloaderMockDownloadExpectation specifies expectation struct of the Downloader.Download
type DownloaderMockDownloadExpectation struct {
	mock    *DownloaderMock
	params  *DownloaderMockDownloadParams
	results *DownloaderMockDownloadResults
	Counter uint64
}

// DownloaderMockDownloadParams contains parameters of the Downloader.Download
type DownloaderMockDownloadParams struct {
	module coordinates.SerialModule
}

// DownloaderMockDownloadResults contains results of the Downloader.Download
type DownloaderMockDownloadResults struct {
	err error
}

// Expect sets up expected params for Downloader.Download
func (mmDownload *mDownloaderMockDownload) Expect(module coordinates.SerialModule) *mDownloaderMockDownload {
	if mmDownload.mock.funcDownload != nil {
		mmDownload.mock.t.Fatalf("DownloaderMock.Download mock is already set by Set")
	}

	if mmDownload.defaultExpectation == nil {
		mmDownload.defaultExpectation = &DownloaderMockDownloadExpectation{}
	}

	mmDownload.defaultExpectation.params = &DownloaderMockDownloadParams{module}
	for _, e := range mmDownload.expectations {
		if minimock.Equal(e.params, mmDownload.defaultExpectation.params) {
			mmDownload.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDownload.defaultExpectation.params)
		}
	}

	return mmDownload
}

// Inspect accepts an inspector function that has same arguments as the Downloader.Download
func (mmDownload *mDownloaderMockDownload) Inspect(f func(module coordinates.SerialModule)) *mDownloaderMockDownload {
	if mmDownload.mock.inspectFuncDownload != nil {
		mmDownload.mock.t.Fatalf("Inspect function is already set for DownloaderMock.Download")
	}

	mmDownload.mock.inspectFuncDownload = f

	return mmDownload
}

// Return sets up results that will be returned by Downloader.Download
func (mmDownload *mDownloaderMockDownload) Return(err error) *DownloaderMock {
	if mmDownload.mock.funcDownload != nil {
		mmDownload.mock.t.Fatalf("DownloaderMock.Download mock is already set by Set")
	}

	if mmDownload.defaultExpectation == nil {
		mmDownload.defaultExpectation = &DownloaderMockDownloadExpectation{mock: mmDownload.mock}
	}
	mmDownload.defaultExpectation.results = &DownloaderMockDownloadResults{err}
	return mmDownload.mock
}

//Set uses given function f to mock the Downloader.Download method
func (mmDownload *mDownloaderMockDownload) Set(f func(module coordinates.SerialModule) (err error)) *DownloaderMock {
	if mmDownload.defaultExpectation != nil {
		mmDownload.mock.t.Fatalf("Default expectation is already set for the Downloader.Download method")
	}

	if len(mmDownload.expectations) > 0 {
		mmDownload.mock.t.Fatalf("Some expectations are already set for the Downloader.Download method")
	}

	mmDownload.mock.funcDownload = f
	return mmDownload.mock
}

// When sets expectation for the Downloader.Download which will trigger the result defined by the following
// Then helper
func (mmDownload *mDownloaderMockDownload) When(module coordinates.SerialModule) *DownloaderMockDownloadExpectation {
	if mmDownload.mock.funcDownload != nil {
		mmDownload.mock.t.Fatalf("DownloaderMock.Download mock is already set by Set")
	}

	expectation := &DownloaderMockDownloadExpectation{
		mock:   mmDownload.mock,
		params: &DownloaderMockDownloadParams{module},
	}
	mmDownload.expectations = append(mmDownload.expectations, expectation)
	return expectation
}

// Then sets up Downloader.Download return parameters for the expectation previously defined by the When method
func (e *DownloaderMockDownloadExpectation) Then(err error) *DownloaderMock {
	e.results = &DownloaderMockDownloadResults{err}
	return e.mock
}

// Download implements Downloader
func (mmDownload *DownloaderMock) Download(module coordinates.SerialModule) (err error) {
	mm_atomic.AddUint64(&mmDownload.beforeDownloadCounter, 1)
	defer mm_atomic.AddUint64(&mmDownload.afterDownloadCounter, 1)

	if mmDownload.inspectFuncDownload != nil {
		mmDownload.inspectFuncDownload(module)
	}

	params := &DownloaderMockDownloadParams{module}

	// Record call args
	mmDownload.DownloadMock.mutex.Lock()
	mmDownload.DownloadMock.callArgs = append(mmDownload.DownloadMock.callArgs, params)
	mmDownload.DownloadMock.mutex.Unlock()

	for _, e := range mmDownload.DownloadMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDownload.DownloadMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDownload.DownloadMock.defaultExpectation.Counter, 1)
		want := mmDownload.DownloadMock.defaultExpectation.params
		got := DownloaderMockDownloadParams{module}
		if want != nil && !minimock.Equal(*want, got) {
			mmDownload.t.Errorf("DownloaderMock.Download got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmDownload.DownloadMock.defaultExpectation.results
		if results == nil {
			mmDownload.t.Fatal("No results are set for the DownloaderMock.Download")
		}
		return (*results).err
	}
	if mmDownload.funcDownload != nil {
		return mmDownload.funcDownload(module)
	}
	mmDownload.t.Fatalf("Unexpected call to DownloaderMock.Download. %v", module)
	return
}

// DownloadAfterCounter returns a count of finished DownloaderMock.Download invocations
func (mmDownload *DownloaderMock) DownloadAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDownload.afterDownloadCounter)
}

// DownloadBeforeCounter returns a count of DownloaderMock.Download invocations
func (mmDownload *DownloaderMock) DownloadBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDownload.beforeDownloadCounter)
}

// Calls returns a list of arguments used in each call to DownloaderMock.Download.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDownload *mDownloaderMockDownload) Calls() []*DownloaderMockDownloadParams {
	mmDownload.mutex.RLock()

	argCopy := make([]*DownloaderMockDownloadParams, len(mmDownload.callArgs))
	copy(argCopy, mmDownload.callArgs)

	mmDownload.mutex.RUnlock()

	return argCopy
}

// MinimockDownloadDone returns true if the count of the Download invocations corresponds
// the number of defined expectations
func (m *DownloaderMock) MinimockDownloadDone() bool {
	for _, e := range m.DownloadMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DownloadMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDownloadCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDownload != nil && mm_atomic.LoadUint64(&m.afterDownloadCounter) < 1 {
		return false
	}
	return true
}

// MinimockDownloadInspect logs each unmet expectation
func (m *DownloaderMock) MinimockDownloadInspect() {
	for _, e := range m.DownloadMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DownloaderMock.Download with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DownloadMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDownloadCounter) < 1 {
		if m.DownloadMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DownloaderMock.Download")
		} else {
			m.t.Errorf("Expected call to DownloaderMock.Download with params: %#v", *m.DownloadMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDownload != nil && mm_atomic.LoadUint64(&m.afterDownloadCounter) < 1 {
		m.t.Error("Expected call to DownloaderMock.Download")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DownloaderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDownloadInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DownloaderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DownloaderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDownloadDone()
}
