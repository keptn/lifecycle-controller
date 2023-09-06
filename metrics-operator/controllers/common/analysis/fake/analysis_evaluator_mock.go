// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1alpha3"
	"github.com/keptn/lifecycle-toolkit/metrics-operator/controllers/common/analysis/types"
	"sync"
)

// IAnalysisEvaluatorMock is a mock implementation of analysis.IAnalysisEvaluator.
//
//	func TestSomethingThatUsesIAnalysisEvaluator(t *testing.T) {
//
//		// make and configure a mocked analysis.IAnalysisEvaluator
//		mockedIAnalysisEvaluator := &IAnalysisEvaluatorMock{
//			EvaluateFunc: func(values map[string]v1alpha3.ProviderResult, ad *v1alpha3.AnalysisDefinition) types.AnalysisResult {
//				panic("mock out the Evaluate method")
//			},
//		}
//
//		// use mockedIAnalysisEvaluator in code that requires analysis.IAnalysisEvaluator
//		// and then make assertions.
//
//	}
type IAnalysisEvaluatorMock struct {
	// EvaluateFunc mocks the Evaluate method.
	EvaluateFunc func(values map[string]v1alpha3.ProviderResult, ad *v1alpha3.AnalysisDefinition) types.AnalysisResult

	// calls tracks calls to the methods.
	calls struct {
		// Evaluate holds details about calls to the Evaluate method.
		Evaluate []struct {
			// Values is the values argument value.
			Values map[string]v1alpha3.ProviderResult
			// Ad is the ad argument value.
			Ad *v1alpha3.AnalysisDefinition
		}
	}
	lockEvaluate sync.RWMutex
}

// Evaluate calls EvaluateFunc.
func (mock *IAnalysisEvaluatorMock) Evaluate(values map[string]v1alpha3.ProviderResult, ad *v1alpha3.AnalysisDefinition) types.AnalysisResult {
	if mock.EvaluateFunc == nil {
		panic("IAnalysisEvaluatorMock.EvaluateFunc: method is nil but IAnalysisEvaluator.Evaluate was just called")
	}
	callInfo := struct {
		Values map[string]v1alpha3.ProviderResult
		Ad     *v1alpha3.AnalysisDefinition
	}{
		Values: values,
		Ad:     ad,
	}
	mock.lockEvaluate.Lock()
	mock.calls.Evaluate = append(mock.calls.Evaluate, callInfo)
	mock.lockEvaluate.Unlock()
	return mock.EvaluateFunc(values, ad)
}

// EvaluateCalls gets all the calls that were made to Evaluate.
// Check the length with:
//
//	len(mockedIAnalysisEvaluator.EvaluateCalls())
func (mock *IAnalysisEvaluatorMock) EvaluateCalls() []struct {
	Values map[string]v1alpha3.ProviderResult
	Ad     *v1alpha3.AnalysisDefinition
} {
	var calls []struct {
		Values map[string]v1alpha3.ProviderResult
		Ad     *v1alpha3.AnalysisDefinition
	}
	mock.lockEvaluate.RLock()
	calls = mock.calls.Evaluate
	mock.lockEvaluate.RUnlock()
	return calls
}
