package keptnevaluation

import (
	"context"
	"testing"

	"github.com/go-logr/logr/testr"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	klcv1alpha3 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha3"
	"github.com/keptn/lifecycle-toolkit/operator/controllers/common/fake"
	"github.com/keptn/lifecycle-toolkit/operator/controllers/common/providers"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const KltNamespace = "klt-namespace"

func TestKeptnEvaluationReconciler_fetchDefinition(t *testing.T) {

	metricEvalDef, DTEvalDef, PromEvalDef, EvalDef := setupEvalDefinitions()
	DTProv, PromProv := setupProviders()
	client := fake.NewClient(metricEvalDef, DTEvalDef, PromEvalDef, EvalDef, DTProv, PromProv)

	r := &KeptnEvaluationReconciler{
		Client: client,
		Scheme: client.Scheme(),
		Log:    testr.New(t),
	}

	tests := []struct {
		name                 string
		namespacedDefinition types.NamespacedName
		wantDef              *klcv1alpha3.KeptnEvaluationDefinition
		wantErr              bool
	}{
		{
			name: "keptn metrics",
			namespacedDefinition: types.NamespacedName{
				Namespace: KltNamespace,
				Name:      "myKeptn",
			},
			wantDef: metricEvalDef,
		},
		{
			name: "DT metrics",
			namespacedDefinition: types.NamespacedName{
				Namespace: KltNamespace,
				Name:      "myDT",
			},
			wantDef: DTEvalDef,
		},

		{
			name: "Prometheus metrics",
			namespacedDefinition: types.NamespacedName{
				Namespace: KltNamespace,
				Name:      "myProm",
			},
			wantDef: PromEvalDef,
		},

		{
			name: "Unexisting Evaluation Def",
			namespacedDefinition: types.NamespacedName{
				Namespace: KltNamespace,
				Name:      "whatever",
			},
			wantDef: nil,
			wantErr: true,
		},
		{
			name: "Unexisting Provider",
			namespacedDefinition: types.NamespacedName{
				Namespace: KltNamespace,
				Name:      "mydef",
			},
			wantDef: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := r.fetchDefinition(context.TODO(), tt.namespacedDefinition)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchDefinitionAndProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantDef != nil {
				require.Equal(t, got.Name, tt.wantDef.Name)
			} else {
				require.Nil(t, got)
			}
		})
	}
}

func setupEvalDefinitions() (*klcv1alpha3.KeptnEvaluationDefinition, *klcv1alpha3.KeptnEvaluationDefinition, *klcv1alpha3.KeptnEvaluationDefinition, *klcv1alpha3.KeptnEvaluationDefinition) {
	metricEvalDef := &klcv1alpha3.KeptnEvaluationDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: KltNamespace,
			Name:      "myKeptn",
		},
		Spec: klcv1alpha3.KeptnEvaluationDefinitionSpec{
			Source:     providers.KeptnMetricProviderName,
			Objectives: nil,
		},
		Status: klcv1alpha3.KeptnEvaluationDefinitionStatus{},
	}

	DTEvalDef := &klcv1alpha3.KeptnEvaluationDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: KltNamespace,
			Name:      "myDT",
		},
		Spec: klcv1alpha3.KeptnEvaluationDefinitionSpec{
			Source:     providers.DynatraceProviderName,
			Objectives: nil,
		},
		Status: klcv1alpha3.KeptnEvaluationDefinitionStatus{},
	}

	PromEvalDef := &klcv1alpha3.KeptnEvaluationDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: KltNamespace,
			Name:      "myProm",
		},
		Spec: klcv1alpha3.KeptnEvaluationDefinitionSpec{
			Source:     providers.PrometheusProviderName,
			Objectives: nil,
		},
		Status: klcv1alpha3.KeptnEvaluationDefinitionStatus{},
	}

	EvalDef := &klcv1alpha3.KeptnEvaluationDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: KltNamespace,
			Name:      "mdef",
		},
		Spec: klcv1alpha3.KeptnEvaluationDefinitionSpec{
			Source:     "dunno",
			Objectives: nil,
		},
		Status: klcv1alpha3.KeptnEvaluationDefinitionStatus{},
	}

	return metricEvalDef, DTEvalDef, PromEvalDef, EvalDef
}

func setupProviders() (*klcv1alpha2.KeptnEvaluationProvider, *klcv1alpha2.KeptnEvaluationProvider) {
	DTProv := &klcv1alpha2.KeptnEvaluationProvider{
		ObjectMeta: metav1.ObjectMeta{
			Name:      providers.DynatraceProviderName,
			Namespace: KltNamespace,
		},
	}

	PromProv := &klcv1alpha2.KeptnEvaluationProvider{
		ObjectMeta: metav1.ObjectMeta{
			Name:      providers.PrometheusProviderName,
			Namespace: KltNamespace,
		},
	}

	return DTProv, PromProv
}
