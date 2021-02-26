package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			// stolen from https://github.com/hashicorp/terraform-provider-kubernetes/blob/master/kubernetes/provider.go
			Schema: map[string]*schema.Schema{
				"host": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_HOST", ""),
					Description: "The hostname (in form of URI) of Kubernetes master.",
				},
				"username": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_USER", ""),
					Description: "The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				},
				"password": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_PASSWORD", ""),
					Description: "The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				},
				"insecure": {
					Type:        schema.TypeBool,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_INSECURE", false),
					Description: "Whether server should be accessed without verifying the TLS certificate.",
				},
				"client_certificate": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_CLIENT_CERT_DATA", ""),
					Description: "PEM-encoded client certificate for TLS authentication.",
				},
				"client_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_CLIENT_KEY_DATA", ""),
					Description: "PEM-encoded client certificate key for TLS authentication.",
				},
				"cluster_ca_certificate": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_CLUSTER_CA_CERT_DATA", ""),
					Description: "PEM-encoded root certificates bundle for TLS authentication.",
				},
				"config_paths": {
					Type:        schema.TypeList,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Optional:    true,
					Description: "A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.",
				},
				"config_path": {
					Type:          schema.TypeString,
					Optional:      true,
					DefaultFunc:   schema.EnvDefaultFunc("KUBE_CONFIG_PATH", nil),
					Description:   "Path to the kube config file. Can be set with KUBE_CONFIG_PATH.",
					ConflictsWith: []string{"config_paths"},
				},
				"config_context": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_CTX", ""),
				},
				"config_context_auth_info": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_CTX_AUTH_INFO", ""),
					Description: "",
				},
				"config_context_cluster": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_CTX_CLUSTER", ""),
					Description: "",
				},
				"token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KUBE_TOKEN", ""),
					Description: "Token to authenticate an service account",
				},
				"exec": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"api_version": {
								Type:     schema.TypeString,
								Required: true,
							},
							"command": {
								Type:     schema.TypeString,
								Required: true,
							},
							"env": {
								Type:     schema.TypeMap,
								Optional: true,
								Elem:     &schema.Schema{Type: schema.TypeString},
							},
							"args": {
								Type:     schema.TypeList,
								Optional: true,
								Elem:     &schema.Schema{Type: schema.TypeString},
							},
						},
					},
					Description: "",
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"backend_config": resourceBackendConfig(),
				//"frontend_config": resourceFrontendConfig(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent
		//userAgent := p.UserAgent("terraform-provider-febeconfig", version)
		return &apiClient{}, nil
	}
}
