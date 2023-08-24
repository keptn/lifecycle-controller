package converter

import (
	"fmt"

	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1alpha3"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SLIConverter struct {
}

type SLI struct {
	Indicators map[string]string `yaml:"indicators"`
}

func NewSLIConverter() *SLIConverter {
	return &SLIConverter{}
}

func (c *SLIConverter) Convert(fileContent []byte, provider string, namespace string) (string, error) {
	//check that provider and namespace is set
	if provider == "" || namespace == "" {
		return "", fmt.Errorf("--sli-provider and --sli-namespace needs to be set for conversion")
	}

	// unmarshall content
	content := &SLI{}
	err := yaml.Unmarshal(fileContent, content)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling file content: %s", err.Error())
	}

	// convert
	analysisValueTemplates := c.convertMapToAnalysisValueTemplate(content.Indicators, provider, namespace)

	result := ""
	for _, v := range analysisValueTemplates {
		// marshal AnalysisValueTemplate to Yaml
		yamlData, err := yaml.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("error marshalling data: %s", err.Error())
		}

		// store output
		result += "---\n"
		result += string(yamlData)
	}

	return result, nil
}

func (c *SLIConverter) convertMapToAnalysisValueTemplate(slis map[string]string, provider string, namespace string) []*metricsapi.AnalysisValueTemplate {
	result := make([]*metricsapi.AnalysisValueTemplate, len(slis))
	i := 0
	for key, query := range slis {
		template := &metricsapi.AnalysisValueTemplate{
			TypeMeta: v1.TypeMeta{
				Kind:       "AnalysisValueTemplate",
				APIVersion: "metrics.keptn.sh/v1alpha3",
			},
			ObjectMeta: v1.ObjectMeta{
				Name: key,
			},
			Spec: metricsapi.AnalysisValueTemplateSpec{
				Query: query,
				Provider: metricsapi.ObjectReference{
					Name:      provider,
					Namespace: namespace,
				},
			},
		}
		result[i] = template
		i++
	}
	return result
}
