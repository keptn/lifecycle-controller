package converter

import (
	"testing"

	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1alpha3"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const sliContent = `
spec_version: "1.0"
indicators:
  throughput: "builtin:service.requestCount.total:merge(0):count?scope=tag(keptn_project:$PROJECT),tag(keptn_stage:$STAGE),tag(keptn_service:$SERVICE),tag(keptn_deployment:$DEPLOYMENT)"
  response_time_p95: "builtin:service.response.time:merge(0):percentile(95)?scope=tag(keptn_project:$PROJECT),tag(keptn_stage:$STAGE),tag(keptn_service:$SERVICE),tag(keptn_deployment:$DEPLOYMENT)"`

const expectedOutput = `---
kind: AnalysisValueTemplate
apiversion: metrics.keptn.sh/v1alpha3
metadata:
  name: throughput
  generatename: ""
  namespace: ""
  selflink: ""
  uid: ""
  resourceversion: ""
  generation: 0
  creationtimestamp: "0001-01-01T00:00:00Z"
  deletiontimestamp: null
  deletiongraceperiodseconds: null
  labels: {}
  annotations: {}
  ownerreferences: []
  finalizers: []
  managedfields: []
spec:
  provider:
    name: dynatrace
    namespace: keptn
  query: builtin:service.requestCount.total:merge(0):count?scope=tag(keptn_project:$PROJECT),tag(keptn_stage:$STAGE),tag(keptn_service:$SERVICE),tag(keptn_deployment:$DEPLOYMENT)
---
kind: AnalysisValueTemplate
apiversion: metrics.keptn.sh/v1alpha3
metadata:
  name: response_time_p95
  generatename: ""
  namespace: ""
  selflink: ""
  uid: ""
  resourceversion: ""
  generation: 0
  creationtimestamp: "0001-01-01T00:00:00Z"
  deletiontimestamp: null
  deletiongraceperiodseconds: null
  labels: {}
  annotations: {}
  ownerreferences: []
  finalizers: []
  managedfields: []
spec:
  provider:
    name: dynatrace
    namespace: keptn
  query: builtin:service.response.time:merge(0):percentile(95)?scope=tag(keptn_project:$PROJECT),tag(keptn_stage:$STAGE),tag(keptn_service:$SERVICE),tag(keptn_deployment:$DEPLOYMENT)
`

func TestConvertMapToAnalysisValueTemplate(t *testing.T) {
	converter := NewSLIConverter()

	// map of slis is nil
	res := converter.convertMapToAnalysisValueTemplate(nil, "provider", "default")
	require.Equal(t, 0, len(res))

	// map of slis is empty
	res = converter.convertMapToAnalysisValueTemplate(map[string]string{}, "provider", "default")
	require.Equal(t, 0, len(res))

	// valid input
	in := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	out := []*metricsapi.AnalysisValueTemplate{
		{
			TypeMeta: v1.TypeMeta{
				Kind:       "AnalysisValueTemplate",
				APIVersion: "metrics.keptn.sh/v1alpha3",
			},
			ObjectMeta: v1.ObjectMeta{
				Name: "key1",
			},
			Spec: metricsapi.AnalysisValueTemplateSpec{
				Query: "val1",
				Provider: metricsapi.ObjectReference{
					Name:      "provider",
					Namespace: "default",
				},
			},
		},
		{
			TypeMeta: v1.TypeMeta{
				Kind:       "AnalysisValueTemplate",
				APIVersion: "metrics.keptn.sh/v1alpha3",
			},
			ObjectMeta: v1.ObjectMeta{
				Name: "key2",
			},
			Spec: metricsapi.AnalysisValueTemplateSpec{
				Query: "val2",
				Provider: metricsapi.ObjectReference{
					Name:      "provider",
					Namespace: "default",
				},
			},
		},
	}
	res = converter.convertMapToAnalysisValueTemplate(in, "provider", "default")
	require.Equal(t, 2, len(res))
	require.Equal(t, out, res)
}

func TestConvertSLI(t *testing.T) {
	c := NewSLIConverter()
	// no provider nor namespace
	res, err := c.Convert([]byte(sliContent), "", "")
	require.NotNil(t, err)
	require.Equal(t, "", res)

	// invalid file content
	res, err = c.Convert([]byte("invalid"), "dynatrace", "keptn")
	require.NotNil(t, err)
	require.Equal(t, "", res)

	// happy path
	res, err = c.Convert([]byte(sliContent), "dynatrace", "keptn")
	require.Nil(t, err)
	require.Equal(t, expectedOutput, res)
}
