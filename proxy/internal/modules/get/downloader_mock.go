package get

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"oss.indeed.com/go/modprox/pkg/coordinates"
)

// DownloaderMock implements Downloader
type DownloaderMock struct {
	t minimock.Tester

	funcDownload          func(module coordinates.SerialModule) (err error)
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

	return m
}

type mDownloaderMockDownload struct {
	mock               *DownloaderMock
	defaultExpectation *DownloaderMockDownloadExpectation
	expectations       []*DownloaderMockDownloadExpectation
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
func (m *mDownloaderMockDownload) Expect(module coordinates.SerialModule) *mDownloaderMockDownload {
	if m.mock.funcDownload != nil {
		m.mock.t.Fatalf("DownloaderMock.Download mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &DownloaderMockDownloadExpectation{}
	}

	m.defaultExpectation.params = &DownloaderMockDownloadParams{module}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Downloader.Download
func (m *mDownloaderMockDownload) Return(err error) *DownloaderMock {
	if m.mock.funcDownload != nil {
		m.mock.t.Fatalf("DownloaderMock.Download mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &DownloaderMockDownloadExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &DownloaderMockDownloadResults{err}
	return m.mock
}

//Set uses given function f to mock the Downloader.Download method
func (m *mDownloaderMockDownload) Set(f func(module coordinates.SerialModule) (err error)) *DownloaderMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Downloader.Download method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Downloader.Download method")
	}

	m.mock.funcDownload = f
	return m.mock
}

// When sets expectation for the Downloader.Download which will trigger the result defined by the following
// Then helper
func (m *mDownloaderMockDownload) When(module coordinates.SerialModule) *DownloaderMockDownloadExpectation {
	if m.mock.funcDownload != nil {
		m.mock.t.Fatalf("DownloaderMock.Download mock is already set by Set")
	}

	expectation := &DownloaderMockDownloadExpectation{
		mock:   m.mock,
		params: &DownloaderMockDownloadParams{module},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Downloader.Download return parameters for the expectation previously defined by the When method
func (e *DownloaderMockDownloadExpectation) Then(err error) *DownloaderMock {
	e.results = &DownloaderMockDownloadResults{err}
	return e.mock
}

// Download implements Downloader
func (m *DownloaderMock) Download(module coordinates.SerialModule) (err error) {
	mm_atomic.AddUint64(&m.beforeDownloadCounter, 1)
	defer mm_atomic.AddUint64(&m.afterDownloadCounter, 1)

	for _, e := range m.DownloadMock.expectations {
		if minimock.Equal(*e.params, DownloaderMockDownloadParams{module}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.DownloadMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.DownloadMock.defaultExpectation.Counter, 1)
		want := m.DownloadMock.defaultExpectation.params
		got := DownloaderMockDownloadParams{module}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("DownloaderMock.Download got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.DownloadMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the DownloaderMock.Download")
		}
		return (*results).err
	}
	if m.funcDownload != nil {
		return m.funcDownload(module)
	}
	m.t.Fatalf("Unexpected call to DownloaderMock.Download. %v", module)
	return
}

// DownloadAfterCounter returns a count of finished DownloaderMock.Download invocations
func (m *DownloaderMock) DownloadAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterDownloadCounter)
}

// DownloadBeforeCounter returns a count of DownloaderMock.Download invocations
func (m *DownloaderMock) DownloadBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeDownloadCounter)
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
		m.t.Errorf("Expected call to DownloaderMock.Download with params: %#v", *m.DownloadMock.defaultExpectation.params)
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