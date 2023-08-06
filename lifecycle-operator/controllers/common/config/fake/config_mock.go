// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"sync"
	"time"
)

// MockConfig is a mock implementation of config.IConfig.
//
//	func TestSomethingThatUsesIConfig(t *testing.T) {
//
//		// make and configure a mocked config.IConfig
//		mockedIConfig := &MockConfig{
//			GetCloudEventsEndpointFunc: func() string {
//				panic("mock out the GetCloudEventsEndpoint method")
//			},
//			GetCreationRequestTimeoutFunc: func() time.Duration {
//				panic("mock out the GetCreationRequestTimeout method")
//			},
//			SetCloudEventsEndpointFunc: func(endpoint string)  {
//				panic("mock out the SetCloudEventsEndpoint method")
//			},
//			SetCreationRequestTimeoutFunc: func(value time.Duration)  {
//				panic("mock out the SetCreationRequestTimeout method")
//			},
//		}
//
//		// use mockedIConfig in code that requires config.IConfig
//		// and then make assertions.
//
//	}
type MockConfig struct {
	// GetCloudEventsEndpointFunc mocks the GetCloudEventsEndpoint method.
	GetCloudEventsEndpointFunc func() string

	// GetCreationRequestTimeoutFunc mocks the GetCreationRequestTimeout method.
	GetCreationRequestTimeoutFunc func() time.Duration

	// SetCloudEventsEndpointFunc mocks the SetCloudEventsEndpoint method.
	SetCloudEventsEndpointFunc func(endpoint string)

	// SetCreationRequestTimeoutFunc mocks the SetCreationRequestTimeout method.
	SetCreationRequestTimeoutFunc func(value time.Duration)

	// calls tracks calls to the methods.
	calls struct {
		// GetCloudEventsEndpoint holds details about calls to the GetCloudEventsEndpoint method.
		GetCloudEventsEndpoint []struct {
		}
		// GetCreationRequestTimeout holds details about calls to the GetCreationRequestTimeout method.
		GetCreationRequestTimeout []struct {
		}
		// SetCloudEventsEndpoint holds details about calls to the SetCloudEventsEndpoint method.
		SetCloudEventsEndpoint []struct {
			// Endpoint is the endpoint argument value.
			Endpoint string
		}
		// SetCreationRequestTimeout holds details about calls to the SetCreationRequestTimeout method.
		SetCreationRequestTimeout []struct {
			// Value is the value argument value.
			Value time.Duration
		}
	}
	lockGetCloudEventsEndpoint    sync.RWMutex
	lockGetCreationRequestTimeout sync.RWMutex
	lockSetCloudEventsEndpoint    sync.RWMutex
	lockSetCreationRequestTimeout sync.RWMutex
}

// GetCloudEventsEndpoint calls GetCloudEventsEndpointFunc.
func (mock *MockConfig) GetCloudEventsEndpoint() string {
	if mock.GetCloudEventsEndpointFunc == nil {
		panic("MockConfig.GetCloudEventsEndpointFunc: method is nil but IConfig.GetCloudEventsEndpoint was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCloudEventsEndpoint.Lock()
	mock.calls.GetCloudEventsEndpoint = append(mock.calls.GetCloudEventsEndpoint, callInfo)
	mock.lockGetCloudEventsEndpoint.Unlock()
	return mock.GetCloudEventsEndpointFunc()
}

// GetCloudEventsEndpointCalls gets all the calls that were made to GetCloudEventsEndpoint.
// Check the length with:
//
//	len(mockedIConfig.GetCloudEventsEndpointCalls())
func (mock *MockConfig) GetCloudEventsEndpointCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCloudEventsEndpoint.RLock()
	calls = mock.calls.GetCloudEventsEndpoint
	mock.lockGetCloudEventsEndpoint.RUnlock()
	return calls
}

// GetCreationRequestTimeout calls GetCreationRequestTimeoutFunc.
func (mock *MockConfig) GetCreationRequestTimeout() time.Duration {
	if mock.GetCreationRequestTimeoutFunc == nil {
		panic("MockConfig.GetCreationRequestTimeoutFunc: method is nil but IConfig.GetCreationRequestTimeout was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCreationRequestTimeout.Lock()
	mock.calls.GetCreationRequestTimeout = append(mock.calls.GetCreationRequestTimeout, callInfo)
	mock.lockGetCreationRequestTimeout.Unlock()
	return mock.GetCreationRequestTimeoutFunc()
}

// GetCreationRequestTimeoutCalls gets all the calls that were made to GetCreationRequestTimeout.
// Check the length with:
//
//	len(mockedIConfig.GetCreationRequestTimeoutCalls())
func (mock *MockConfig) GetCreationRequestTimeoutCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCreationRequestTimeout.RLock()
	calls = mock.calls.GetCreationRequestTimeout
	mock.lockGetCreationRequestTimeout.RUnlock()
	return calls
}

// SetCloudEventsEndpoint calls SetCloudEventsEndpointFunc.
func (mock *MockConfig) SetCloudEventsEndpoint(endpoint string) {
	if mock.SetCloudEventsEndpointFunc == nil {
		panic("MockConfig.SetCloudEventsEndpointFunc: method is nil but IConfig.SetCloudEventsEndpoint was just called")
	}
	callInfo := struct {
		Endpoint string
	}{
		Endpoint: endpoint,
	}
	mock.lockSetCloudEventsEndpoint.Lock()
	mock.calls.SetCloudEventsEndpoint = append(mock.calls.SetCloudEventsEndpoint, callInfo)
	mock.lockSetCloudEventsEndpoint.Unlock()
	mock.SetCloudEventsEndpointFunc(endpoint)
}

// SetCloudEventsEndpointCalls gets all the calls that were made to SetCloudEventsEndpoint.
// Check the length with:
//
//	len(mockedIConfig.SetCloudEventsEndpointCalls())
func (mock *MockConfig) SetCloudEventsEndpointCalls() []struct {
	Endpoint string
} {
	var calls []struct {
		Endpoint string
	}
	mock.lockSetCloudEventsEndpoint.RLock()
	calls = mock.calls.SetCloudEventsEndpoint
	mock.lockSetCloudEventsEndpoint.RUnlock()
	return calls
}

// SetCreationRequestTimeout calls SetCreationRequestTimeoutFunc.
func (mock *MockConfig) SetCreationRequestTimeout(value time.Duration) {
	if mock.SetCreationRequestTimeoutFunc == nil {
		panic("MockConfig.SetCreationRequestTimeoutFunc: method is nil but IConfig.SetCreationRequestTimeout was just called")
	}
	callInfo := struct {
		Value time.Duration
	}{
		Value: value,
	}
	mock.lockSetCreationRequestTimeout.Lock()
	mock.calls.SetCreationRequestTimeout = append(mock.calls.SetCreationRequestTimeout, callInfo)
	mock.lockSetCreationRequestTimeout.Unlock()
	mock.SetCreationRequestTimeoutFunc(value)
}

// SetCreationRequestTimeoutCalls gets all the calls that were made to SetCreationRequestTimeout.
// Check the length with:
//
//	len(mockedIConfig.SetCreationRequestTimeoutCalls())
func (mock *MockConfig) SetCreationRequestTimeoutCalls() []struct {
	Value time.Duration
} {
	var calls []struct {
		Value time.Duration
	}
	mock.lockSetCreationRequestTimeout.RLock()
	calls = mock.calls.SetCreationRequestTimeout
	mock.lockSetCreationRequestTimeout.RUnlock()
	return calls
}
