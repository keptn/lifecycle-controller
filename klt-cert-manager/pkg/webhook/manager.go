package webhook

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

const (
	metricsBindAddress = ":8383"
	port               = 8443
)

//go:generate moq -pkg fake -skip-ensure -out ../fake/manager_mock.go . IManager:MockManager
type IManager manager.Manager

//go:generate moq -pkg fake -skip-ensure -out ../fake/webhookmanager_mock.go . Provider:MockWebhookManager
type Provider interface {
	SetupWebhookServer(mgr manager.Manager)
}

type WebhookProvider struct {
	certificateDirectory string
	certificateFileName  string
	keyFileName          string
}

func NewWebhookManagerProvider(certificateDirectory string, keyFileName string, certificateFileName string) WebhookProvider {
	return WebhookProvider{
		certificateDirectory: certificateDirectory,
		certificateFileName:  certificateFileName,
		keyFileName:          keyFileName,
	}
}

func (provider WebhookProvider) createOptions(scheme *runtime.Scheme, namespace string) ctrl.Options {
	return ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsBindAddress,
		WebhookServer:      webhook.NewServer(webhook.Options{Port: port}),
		Namespace:          namespace,
	}
}

func (provider WebhookProvider) SetupWebhookServer(mgr manager.Manager) {
	webhookServer := mgr.GetWebhookServer()
	webhookServer.(*webhook.DefaultServer).Options.CertDir = provider.certificateDirectory
	webhookServer.(*webhook.DefaultServer).Options.KeyName = provider.keyFileName
	webhookServer.(*webhook.DefaultServer).Options.CertName = provider.certificateFileName
}
