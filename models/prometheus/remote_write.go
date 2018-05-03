package prometheus

import (
	"github.com/prometheus/common/model"
)

// RemoteWrite represents remote_write section
type RemoteWrite struct {
	// The URL of the endpoint to send samples to.
	URL string `json:"url" yaml:"url"`

	// Timeout for requests to the remote write endpoint.
	RemoteTimeout model.Duration `json:"remote_timeout,omitempty" yaml:"remote_timeout,omitempty"` // default = 30s

	// List of remote write relabel configurations.
	WriteRelabelConfigs []*RelabelConfig `json:"write_relabel_configs,omitempty" yaml:"write_relabel_configs,omitempty"`

	// Sets the `Authorization` header on every scrape request with the
	// configured username and password.
	BasicAuth *BasicAuth `json:"basic_auth,omitempty" yaml:"basic_auth,omitempty"`

	// Sets the `Authorization` header on every scrape request with
	// the configured bearer token. It is mutually exclusive with `bearer_token_file`.
	BearerToken Secret `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`

	// Sets the `Authorization` header on every scrape request with the bearer token
	// read from the configured file. It is mutually exclusive with `bearer_token`.
	BearerTokenFile string `json:"bearer_token_file,omitempty" yaml:"bearer_token_file,omitempty"`

	// Configures the scrape request's TLS settings.
	TLSConfig *TLSConfig `json:"tls_config,omitempty" yaml:"tls_config,omitempty"`

	// Optional proxy URL.
	ProxyURL string `json:"proxy_url,omitempty" yaml:"proxy_url,omitempty"`
}
