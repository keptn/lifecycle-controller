package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

type apiConfig struct {
	serverURL        string
	authURL          string
	oAuthCredentials oAuthCredentials
}
type secretValues struct {
	Token   string `json:"token"`
	AuthUrl string `json:"authurl"`
}

type APIConfigOption func(config *apiConfig)

func WithAuthURL(authURL string) APIConfigOption {
	return func(config *apiConfig) {
		config.authURL = authURL
	}
}

// WithScopes passes the given scopes to the client config
func WithScopes(scopes []OAuthScope) APIConfigOption {
	return func(config *apiConfig) {
		config.oAuthCredentials.scopes = scopes
	}
}

// NewAPIConfig returns a new apiConfig that can be used for initializing a DTAPIClient with the NewAPIClient function
func NewAPIConfig(serverURL string, secret []byte, opts ...APIConfigOption) (*apiConfig, error) {
	var secValue secretValues
	if err := json.Unmarshal(secret, &secValue); err != nil {
		return nil, err
	}

	if err := validateOAuthSecret(secValue.Token); err != nil {
		return nil, err
	}

	secretParts := strings.Split(secValue.Token, ".")
	clientId := fmt.Sprintf("%s.%s", secretParts[0], secretParts[1])
	clientSecret := fmt.Sprintf("%s.%s", clientId, secretParts[2])

	cfg := &apiConfig{
		serverURL: serverURL,
		authURL:   secValue.AuthUrl,
		oAuthCredentials: oAuthCredentials{
			clientID:     clientId,
			clientSecret: clientSecret,
			scopes:       []OAuthScope{OAuthScopeStorageMetricsRead, OAuthScopeEnvironmentRoleViewer},
		},
	}

	for _, o := range opts {
		o(cfg)
	}

	return cfg, nil
}
