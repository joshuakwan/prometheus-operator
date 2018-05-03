package prometheus

import (
	"github.com/prometheus/common/model"
	"regexp"
	"github.com/getlantern/deepcopy"
)

// Secret represents a Secret string
type Secret string

// Regexp represents a Regexp
type Regexp struct {
	*regexp.Regexp
}

// BasicAuth defines Basic Auth information
type BasicAuth struct {
	Username string `json:"username" yaml:"username"`
	Password Secret `json:"password" yaml:"username"`
}

type Message struct {
	Text string `json:"message"`
}

// TLSConfig defines TLS configuration items
type TLSConfig struct {
	// CA certificate to validate the server certificate with.
	CAFile string `json:"ca_file" yaml:"ca_file"`
	// Certificate and key files for client cert authentication to the server.
	CertFile string `json:"cert_file" yaml:"cert_file"`
	KeyFile  string `json:"key_file" yaml:"key_file"`
	// ServerName extension to indicate the name of the server.
	ServerName string `json:"server_name" yaml:"server_name"`
	// Disable validation of the server certificate.
	InsecureSkipVerity bool `json:"insecure_skip_verity" yaml:"insecure_skip_verity"`
}

func (re *Regexp) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	regex, err := regexp.Compile("^(?:" + s + ")$")
	if err != nil {
		return err
	}
	re.Regexp = regex
	return nil
}

func (re Regexp) MarshalYAML() (interface{}, error) {
	if re.Regexp != nil {
		return re.String(), nil
	}
	return nil, nil
}

func Update(dst interface{}, src interface{}) {
	deepcopy.Copy(dst, src)
}

// Alerting represents alerting section
type Alerting struct {
	AlertRelabelConfigs []*RelabelConfig      `json:"alert_relabel_configs,omitempty" yaml:"alert_relabel_configs,omitempty"`
	AlertManagers       []*AlertManagerConfig `json:"alertmanagers,omitempty" yaml:"alertmanagers,omitempty"`
}

// AlertManagerConfig represents alerting.alertmanagers section
type AlertManagerConfig struct {
	// Per-target Alertmanager timeout when pushing alerts.
	Timeout model.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"` //default = 10s

	// Prefix for the HTTP path alerts are pushed to.
	PathPrefix string `json:"path_prefix,omitempty" yaml:"path_prefix,omitempty"` //default = /

	// Configures the protocol scheme used for requests.
	// [ scheme: <scheme> | default = http ]
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`

	// Sets the `Authorization` header on every request with the
	// configured username and password.
	BasicAuth *BasicAuth `json:"basic_auth,omitempty" yaml:"basic_auth,omitempty"`

	// Sets the `Authorization` header on every request with
	// the configured bearer token. It is mutually exclusive with `bearer_token_file`.
	BearerToken Secret `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`

	// Sets the `Authorization` header on every request with the bearer token
	// read from the configured file. It is mutually exclusive with `bearer_token`.
	BearerTokenFile string `json:"bearer_token_file,omitempty" yaml:"bearer_token_file,omitempty"`

	// Configures the scrape request's TLS settings.
	TLSConfig *TLSConfig `json:"tls_config,omitempty" yaml:"tls_config,omitempty"`

	// Optional proxy URL.
	ProxyURL string `json:"proxy_url,omitempty" yaml:"proxy_url,omitempty"`

	// List of Azure service discovery configurations.
	AzureSdConfigs []*AzureSdConfig `json:"azure_sd_configs,omitempty" yaml:"azure_sd_configs,omitempty"`

	// List of Consul service discovery configurations.
	ConsulSdConfigs []*ConsulSdConfig `json:"consul_sd_configs,omitempty" yaml:"consul_sd_configs,omitempty"`

	// List of DNS service discovery configurations.
	DNSSdConfigs []*DNSSdConfig `json:"dns_sd_configs,omitempty" yaml:"dns_sd_configs,omitempty"`

	// List of EC2 service discovery configurations.
	EC2SdConfigs []*EC2SdConfig `json:"ec2_sd_configs,omitempty" yaml:"ec2_sd_configs,omitempty"`

	// List of OpenStack service discovery configurations.
	OpenStackSdConfigs []*OpenStackSdConfig `json:"openstack_sd_configs,omitempty" yaml:"openstack_sd_configs,omitempty"`

	// List of file service discovery configurations.
	FileSdConfigs []*FileSdConfig `json:"file_sd_configs,omitempty" yaml:"file_sd_configs,omitempty"`

	// List of GCE service discovery configurations.
	GCESdConfigs []*GCESdConfig `json:"gce_sd_configs,omitempty" yaml:"gce_sd_configs,omitempty"`

	// List of Kubernetes service discovery configurations.
	KubernetesSdConfigs []*KubernetesSdConfig `json:"kubernetes_sd_configs,omitempty" yaml:"kubernetes_sd_configs,omitempty"`

	// List of Marathon service discovery configurations.
	MarathonSdConfigs []*MarathonSdConfig `json:"marathon_sd_configs,omitempty" yaml:"marathon_sd_configs,omitempty"`

	// List of AirBnB's Nerve service discovery configurations.
	NerveSdConfigs []*NerveSdConfig `json:"nerve_sd_configs,omitempty" yaml:"nerve_sd_configs,omitempty"`

	// List of Zookeeper Serverset service discovery configurations.
	ServersetSdConfigs []*ServersetSdConfig `json:"serverset_sd_configs,omitempty" yaml:"serverset_sd_configs,omitempty"`

	// List of Triton service discovery configurations.
	TritonSdConfigs []*TritonSdConfig `json:"triton_sd_configs,omitempty" yaml:"triton_sd_configs,omitempty"`

	// List of labeled statically configured targets for this job.
	StaticConfigs []*StaticConfig `json:"static_configs,omitempty" yaml:"static_configs,omitempty"`

	// List of target relabel configurations.
	RelabelConfigs []*RelabelConfig `json:"relabel_configs,omitempty" yaml:"relabel_configs,omitempty"`
}
