// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"sync"
)

// ITracerMock is a mock implementation of common.ITracer.
//
// 	func TestSomethingThatUsesITracer(t *testing.T) {
//
// 		// make and configure a mocked common.ITracer
// 		mockedITracer := &ITracerMock{
// 			StartFunc: func(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
// 				panic("mock out the Start method")
// 			},
// 		}
//
// 		// use mockedITracer in code that requires common.ITracer
// 		// and then make assertions.
//
// 	}
type ITracerMock struct {
	// StartFunc mocks the Start method.
	StartFunc func(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span)

	// calls tracks calls to the methods.
	calls struct {
		// Start holds details about calls to the Start method.
		Start []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SpanName is the spanName argument value.
			SpanName string
			// Opts is the opts argument value.
			Opts []trace.SpanStartOption
		}
	}
	lockStart *sync.RWMutex
}

// Start calls StartFunc.
func (mock ITracerMock) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if mock.StartFunc == nil {
		panic("ITracerMock.StartFunc: method is nil but ITracer.Start was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		SpanName string
		Opts     []trace.SpanStartOption
	}{
		Ctx:      ctx,
		SpanName: spanName,
		Opts:     opts,
	}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	return mock.StartFunc(ctx, spanName, opts...)
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//     len(mockedITracer.StartCalls())
func (mock ITracerMock) StartCalls() []struct {
	Ctx      context.Context
	SpanName string
	Opts     []trace.SpanStartOption
} {
	var calls []struct {
		Ctx      context.Context
		SpanName string
		Opts     []trace.SpanStartOption
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}
