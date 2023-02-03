package dynatrace

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-logr/logr"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	dtclient "github.com/keptn/lifecycle-toolkit/operator/controllers/common/providers/dynatrace/client"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type keptnDynatraceDQLProvider struct {
	log       logr.Logger
	k8sClient client.Client

	dtClient dtclient.DTAPIClient
}

type DynatraceDQLHandler struct {
	RequestToken string `json:"requestToken"`
}

type DynatraceDQLResult struct {
	State  string    `json:"state"`
	Result DQLResult `json:"result,omitempty"`
}

type DQLResult struct {
	Records []DQLRecord `json:"records"`
}

type DQLRecord struct {
	Value DQLMetric `json:"value"`
}

type DQLMetric struct {
	count int64   `json:"count"`
	sum   float64 `json:"sum"`
	min   float64 `json:"min"`
	avg   float64 `json:"avg"`
	max   float64 `json:"max"`
}

type KeptnDynatraceDQLProviderOption func(provider *keptnDynatraceDQLProvider)

func WithDTAPIClient(dtApiClient dtclient.DTAPIClient) KeptnDynatraceDQLProviderOption {
	return func(provider *keptnDynatraceDQLProvider) {
		provider.dtClient = dtApiClient
	}
}

func WithLogger(logger logr.Logger) KeptnDynatraceDQLProviderOption {
	return func(provider *keptnDynatraceDQLProvider) {
		provider.log = logger
	}
}

func NewKeptnDynatraceDQLProvider(k8sClient client.Client, opts ...KeptnDynatraceDQLProviderOption) (*keptnDynatraceDQLProvider, error) {
	provider := &keptnDynatraceDQLProvider{
		log:       logr.Logger{},
		k8sClient: k8sClient,
	}

	for _, o := range opts {
		o(provider)
	}

	return provider, nil
}

// EvaluateQuery fetches the SLI values from dynatrace provider
func (d *keptnDynatraceDQLProvider) EvaluateQuery(ctx context.Context, objective klcv1alpha2.Objective, provider klcv1alpha2.KeptnEvaluationProvider) (string, []byte, error) {
	if err := d.ensureDTClientIsSetUp(ctx, provider); err != nil {
		return "", nil, err
	}
	// submit DQL
	dqlHandler, err := d.postDQL(ctx, objective.Query)
	if err != nil {
		d.log.Error(err, "Error while posting the DQL query: %s", objective.Query)
		return "", nil, err
	}
	// attend result
	results, err := d.getDQL(ctx, *dqlHandler)
	if err != nil {
		d.log.Error(err, "Error while waiting for DQL query: %s", dqlHandler)
		return "", nil, err
	}
	// parse result
	if len(results.Records) > 1 {
		d.log.Info("More than a single result, the first one will be used")
	}
	if len(results.Records) == 0 {
		return "", nil, ErrInvalidResult
	}
	r := fmt.Sprintf("%f", results.Records[0].Value.avg)
	b, err := json.Marshal(results)
	if err != nil {
		d.log.Error(err, "Error marshaling DQL results")
	}
	return r, b, nil
}

func (d *keptnDynatraceDQLProvider) ensureDTClientIsSetUp(ctx context.Context, provider klcv1alpha2.KeptnEvaluationProvider) error {
	// try to initialize the DT API Client if it has not been set in the options
	if d.dtClient == nil {
		secret, err := getDTSecret(ctx, provider, d.k8sClient)
		if err != nil {
			return err
		}
		config, err := dtclient.NewAPIConfig(
			provider.Spec.TargetServer,
			secret,
			dtclient.WithScopes(d.getScopes()),
		)
		if err != nil {
			return err
		}
		d.dtClient = dtclient.NewAPIClient(*config, dtclient.WithLogger(d.log))
	}
	return nil
}

func (d *keptnDynatraceDQLProvider) getScopes() string {
	return "storage:metrics:read environment:roles:viewer"
}

func (d *keptnDynatraceDQLProvider) postDQL(ctx context.Context, query string) (*DynatraceDQLHandler, error) {
	d.log.V(10).Info("posting DQL")
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	values := url.Values{}
	values.Add("query", query)

	path := fmt.Sprintf("/platform/storage/query/v0.7/query:execute?%s", values.Encode())

	b, err := d.dtClient.Do(ctx, path, http.MethodPost, []byte(`{}`))
	if err != nil {
		return nil, err
	}

	dqlHandler := &DynatraceDQLHandler{}
	err = json.Unmarshal(b, &dqlHandler)
	if err != nil {
		return nil, err
	}
	return dqlHandler, nil
}

func (d *keptnDynatraceDQLProvider) getDQL(ctx context.Context, handler DynatraceDQLHandler) (*DQLResult, error) {
	d.log.V(10).Info("posting DQL")
	for true {
		r, err := d.retrieveDQLResults(ctx, handler)
		if err != nil {
			return &DQLResult{}, err
		}
		if r.State == "SUCCEEDED" {
			return &r.Result, nil
		}
		d.log.V(10).Info("DQL not finished, got: %s", r.State)
		time.Sleep(5 * time.Second)
	}
	return nil, errors.New("something went wrong while waiting for the DQL to be finished")
}

func (d *keptnDynatraceDQLProvider) retrieveDQLResults(ctx context.Context, handler DynatraceDQLHandler) (*DynatraceDQLResult, error) {
	d.log.V(10).Info("Getting DQL")
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	values := url.Values{}
	values.Add("request-token", handler.RequestToken)

	path := fmt.Sprintf("/platform/storage/query/v0.7/query:poll?%s", values.Encode())

	b, err := d.dtClient.Do(ctx, path, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	result := &DynatraceDQLResult{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		d.log.Error(err, "Error while parsing response")
		return result, err
	}
	return result, nil
}
