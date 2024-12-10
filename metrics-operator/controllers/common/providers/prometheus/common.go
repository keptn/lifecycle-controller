package prometheus

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1"
	promapi "github.com/prometheus/client_golang/api"
	"github.com/prometheus/common/config"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const secretKeyUserName = "user"
const secretKeyPassword = "password"

var ErrSecretKeyRefNotDefined = errors.New("the SecretKeyRef property with the Prometheus API Key is missing")
var ErrInvalidSecretFormat = errors.New("secret key does not contain user and password")

type SecretData struct {
	User     string        `json:"user"`
	Password config.Secret `json:"password"`
}

//go:generate moq -pkg fake -skip-ensure -out ./fake/roundtripper_mock.go . IRoundTripper
type IRoundTripper interface {
	GetRoundTripper(context.Context, metricsapi.KeptnMetricsProvider, client.Client) (http.RoundTripper, error)
}

type RoundTripperRetriever struct {
}

func (r RoundTripperRetriever) GetRoundTripper(ctx context.Context, provider metricsapi.KeptnMetricsProvider, k8sClient client.Client) (http.RoundTripper, error) {
	secret, err := getPrometheusSecret(ctx, provider, k8sClient)
	if err != nil {
		if errors.Is(err, ErrSecretKeyRefNotDefined) {
			return promapi.DefaultRoundTripper, nil
		}
		return nil, err
	}

	transport := promapi.DefaultRoundTripper.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: provider.Spec.InsecureSkipTlsVerify,
	}

	return config.NewBasicAuthRoundTripper(secret.User, secret.Password, "", "", transport), nil
}

func getPrometheusSecret(ctx context.Context, provider metricsapi.KeptnMetricsProvider, k8sClient client.Client) (*SecretData, error) {
	if !provider.HasSecretDefined() {
		return nil, ErrSecretKeyRefNotDefined
	}
	secret := &corev1.Secret{}
	if err := k8sClient.Get(ctx, types.NamespacedName{Name: provider.Spec.SecretKeyRef.Name, Namespace: provider.Namespace}, secret); err != nil {
		return nil, err
	}
	var secretData SecretData
	user, ok := secret.Data[secretKeyUserName]
	pw, yes := secret.Data[secretKeyPassword]
	if !ok || !yes {
		return nil, ErrInvalidSecretFormat
	}
	secretData.User = string(user)
	secretData.Password = config.Secret(pw)
	return &secretData, nil
}
