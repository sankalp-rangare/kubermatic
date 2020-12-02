// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/addon"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/admin"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/alibaba"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/aws"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/azure"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/credentials"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/datacenter"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/digitalocean"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/gcp"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/hetzner"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/metric"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/openstack"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/operations"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/packet"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/project"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/serviceaccounts"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/settings"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/tokens"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/users"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/versions"
	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/client/vsphere"
)

// Default kubermatic HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "dev.kubermatic.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new kubermatic HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Kubermatic {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new kubermatic HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Kubermatic {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new kubermatic client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Kubermatic {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Kubermatic)
	cli.Transport = transport
	cli.Addon = addon.New(transport, formats)
	cli.Admin = admin.New(transport, formats)
	cli.Alibaba = alibaba.New(transport, formats)
	cli.Aws = aws.New(transport, formats)
	cli.Azure = azure.New(transport, formats)
	cli.Credentials = credentials.New(transport, formats)
	cli.Datacenter = datacenter.New(transport, formats)
	cli.Digitalocean = digitalocean.New(transport, formats)
	cli.Gcp = gcp.New(transport, formats)
	cli.Hetzner = hetzner.New(transport, formats)
	cli.Metric = metric.New(transport, formats)
	cli.Openstack = openstack.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Packet = packet.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Serviceaccounts = serviceaccounts.New(transport, formats)
	cli.Settings = settings.New(transport, formats)
	cli.Tokens = tokens.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Versions = versions.New(transport, formats)
	cli.Vsphere = vsphere.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Kubermatic is a client for kubermatic
type Kubermatic struct {
	Addon addon.ClientService

	Admin admin.ClientService

	Alibaba alibaba.ClientService

	Aws aws.ClientService

	Azure azure.ClientService

	Credentials credentials.ClientService

	Datacenter datacenter.ClientService

	Digitalocean digitalocean.ClientService

	Gcp gcp.ClientService

	Hetzner hetzner.ClientService

	Metric metric.ClientService

	Openstack openstack.ClientService

	Operations operations.ClientService

	Packet packet.ClientService

	Project project.ClientService

	Serviceaccounts serviceaccounts.ClientService

	Settings settings.ClientService

	Tokens tokens.ClientService

	Users users.ClientService

	Versions versions.ClientService

	Vsphere vsphere.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Kubermatic) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Addon.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.Alibaba.SetTransport(transport)
	c.Aws.SetTransport(transport)
	c.Azure.SetTransport(transport)
	c.Credentials.SetTransport(transport)
	c.Datacenter.SetTransport(transport)
	c.Digitalocean.SetTransport(transport)
	c.Gcp.SetTransport(transport)
	c.Hetzner.SetTransport(transport)
	c.Metric.SetTransport(transport)
	c.Openstack.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Packet.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Serviceaccounts.SetTransport(transport)
	c.Settings.SetTransport(transport)
	c.Tokens.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Versions.SetTransport(transport)
	c.Vsphere.SetTransport(transport)
}