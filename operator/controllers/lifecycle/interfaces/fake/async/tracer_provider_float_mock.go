// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/asyncfloat64"
	"sync"
)

// ITracerProviderAsyncFloat64Mock is a mock implementation of interfaces.ITracerProviderAsyncFloat64.
//
//	func TestSomethingThatUsesITracerProviderAsyncFloat64(t *testing.T) {
//
//		// make and configure a mocked interfaces.ITracerProviderAsyncFloat64
//		mockedITracerProviderAsyncFloat64 := &ITracerProviderAsyncFloat64Mock{
//			CounterFunc: func(name string, opts ...instrument.Option) (asyncfloat64.Counter, error) {
//				panic("mock out the Counter method")
//			},
//			GaugeFunc: func(name string, opts ...instrument.Option) (asyncfloat64.Gauge, error) {
//				panic("mock out the Gauge method")
//			},
//			UpDownCounterFunc: func(name string, opts ...instrument.Option) (asyncfloat64.UpDownCounter, error) {
//				panic("mock out the UpDownCounter method")
//			},
//		}
//
//		// use mockedITracerProviderAsyncFloat64 in code that requires interfaces.ITracerProviderAsyncFloat64
//		// and then make assertions.
//
//	}
type ITracerProviderAsyncFloat64Mock struct {
	// CounterFunc mocks the Counter method.
	CounterFunc func(name string, opts ...instrument.Option) (asyncfloat64.Counter, error)

	// GaugeFunc mocks the Gauge method.
	GaugeFunc func(name string, opts ...instrument.Option) (asyncfloat64.Gauge, error)

	// UpDownCounterFunc mocks the UpDownCounter method.
	UpDownCounterFunc func(name string, opts ...instrument.Option) (asyncfloat64.UpDownCounter, error)

	// calls tracks calls to the methods.
	calls struct {
		// Counter holds details about calls to the Counter method.
		Counter []struct {
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts []instrument.Option
		}
		// Gauge holds details about calls to the Gauge method.
		Gauge []struct {
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts []instrument.Option
		}
		// UpDownCounter holds details about calls to the UpDownCounter method.
		UpDownCounter []struct {
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts []instrument.Option
		}
	}
	lockCounter       sync.RWMutex
	lockGauge         sync.RWMutex
	lockUpDownCounter sync.RWMutex
}

// Counter calls CounterFunc.
func (mock *ITracerProviderAsyncFloat64Mock) Counter(name string, opts ...instrument.Option) (asyncfloat64.Counter, error) {
	if mock.CounterFunc == nil {
		panic("ITracerProviderAsyncFloat64Mock.CounterFunc: method is nil but ITracerProviderAsyncFloat64.Counter was just called")
	}
	callInfo := struct {
		Name string
		Opts []instrument.Option
	}{
		Name: name,
		Opts: opts,
	}
	mock.lockCounter.Lock()
	mock.calls.Counter = append(mock.calls.Counter, callInfo)
	mock.lockCounter.Unlock()
	return mock.CounterFunc(name, opts...)
}

// CounterCalls gets all the calls that were made to Counter.
// Check the length with:
//
//	len(mockedITracerProviderAsyncFloat64.CounterCalls())
func (mock *ITracerProviderAsyncFloat64Mock) CounterCalls() []struct {
	Name string
	Opts []instrument.Option
} {
	var calls []struct {
		Name string
		Opts []instrument.Option
	}
	mock.lockCounter.RLock()
	calls = mock.calls.Counter
	mock.lockCounter.RUnlock()
	return calls
}

// Gauge calls GaugeFunc.
func (mock *ITracerProviderAsyncFloat64Mock) Gauge(name string, opts ...instrument.Option) (asyncfloat64.Gauge, error) {
	if mock.GaugeFunc == nil {
		panic("ITracerProviderAsyncFloat64Mock.GaugeFunc: method is nil but ITracerProviderAsyncFloat64.Gauge was just called")
	}
	callInfo := struct {
		Name string
		Opts []instrument.Option
	}{
		Name: name,
		Opts: opts,
	}
	mock.lockGauge.Lock()
	mock.calls.Gauge = append(mock.calls.Gauge, callInfo)
	mock.lockGauge.Unlock()
	return mock.GaugeFunc(name, opts...)
}

// GaugeCalls gets all the calls that were made to Gauge.
// Check the length with:
//
//	len(mockedITracerProviderAsyncFloat64.GaugeCalls())
func (mock *ITracerProviderAsyncFloat64Mock) GaugeCalls() []struct {
	Name string
	Opts []instrument.Option
} {
	var calls []struct {
		Name string
		Opts []instrument.Option
	}
	mock.lockGauge.RLock()
	calls = mock.calls.Gauge
	mock.lockGauge.RUnlock()
	return calls
}

// UpDownCounter calls UpDownCounterFunc.
func (mock *ITracerProviderAsyncFloat64Mock) UpDownCounter(name string, opts ...instrument.Option) (asyncfloat64.UpDownCounter, error) {
	if mock.UpDownCounterFunc == nil {
		panic("ITracerProviderAsyncFloat64Mock.UpDownCounterFunc: method is nil but ITracerProviderAsyncFloat64.UpDownCounter was just called")
	}
	callInfo := struct {
		Name string
		Opts []instrument.Option
	}{
		Name: name,
		Opts: opts,
	}
	mock.lockUpDownCounter.Lock()
	mock.calls.UpDownCounter = append(mock.calls.UpDownCounter, callInfo)
	mock.lockUpDownCounter.Unlock()
	return mock.UpDownCounterFunc(name, opts...)
}

// UpDownCounterCalls gets all the calls that were made to UpDownCounter.
// Check the length with:
//
//	len(mockedITracerProviderAsyncFloat64.UpDownCounterCalls())
func (mock *ITracerProviderAsyncFloat64Mock) UpDownCounterCalls() []struct {
	Name string
	Opts []instrument.Option
} {
	var calls []struct {
		Name string
		Opts []instrument.Option
	}
	mock.lockUpDownCounter.RLock()
	calls = mock.calls.UpDownCounter
	mock.lockUpDownCounter.RUnlock()
	return calls
}
