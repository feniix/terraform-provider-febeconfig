package provider

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackendConfig() *schema.Resource {
	return &schema.Resource{
		//CreateContext: resourceBackendConfigCreate,
		//ReadContext: resourceBackendConfigRead,
		//UpdateContext: resourceBackendConfigUpdate,
		//DeleteContext: resourceBackendConfigDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: resourceBackendConfigSchemaV1(),
	}
}

func resourceBackendConfigSchemaV1() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": namespacedMetadataSchema("backendconfig", false),
		"spec": {
			Type:        schema.TypeList,
			Description: "Spec defines the specification of the dessired behavior of the backendconfig. More info: https://cloud.google.com/kubernetes-engine/docs/how-to/ingress-features#configuring_ingress_features_through_backendconfig_parameters",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"timeout_sec": {
						Type:        schema.TypeInt,
						Description: "Set a backend service timeout period in seconds. If you do not specify a value, the default value is 30 seconds.",
						Optional:    true,
						Default:     30,
					},
					"connection_draining": {
						Type: schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"draining_timeout_sec": {
									Type: schema.TypeInt,

								},
							},
						},
					},
					"cdn": {
						Type:        schema.TypeMap,
						Description: "TODO: CDN configuration",
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:        schema.TypeBool,
									Description: "TODO: If set to true, Cloud CDN is enabled for this Ingress backend.",
									Optional:    true,
									Default:     false,
								},
								"cache_policy": {
									Type:        schema.TypeMap,
									Description: "TODO: configure cache policy",
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"include_host": {
												Type:        schema.TypeBool,
												Description: "TODO: If set to true, requests to different hosts are cached separately.",
												Optional:    true,
											},
											"include_protocol": {
												Type:        schema.TypeBool,
												Description: "TODO: If set to true, HTTP and HTTPS requests are cached separately.",
												Optional:    true,
											},
											"include_query_string": {
												Type:        schema.TypeBool,
												Description: "TODO: If set to true, query string parameters are included in the cache key according to query_string_blacklist or query_string_whitelist. If neither is set, the entire query string is included. If set to false, the entire query string is excluded from the cache key.",
												Optional:    true,
											},
											"query_string_blacklist": {
												Type:          schema.TypeSet,
												Description:   "TODO: Specify a string array with the names of query string parameters to exclude from cache keys. All other parameters are included. You can specify query_string_blacklist or query_string_whitelist, but not both.",
												Optional:      true,
												ConflictsWith: []string{"query_string_whitelist"},
												Elem:          &schema.Schema{Type: schema.TypeString},
											},
											"query_string_whitelist": {
												Type:          schema.TypeSet,
												Description:   "TODO: Specify a string array with the names of query string parameters to include in cache keys. All other parameters are excluded. You can query_string_blacklist or query_string_whitelist, but not both.",
												Optional:      true,
												ConflictsWith: []string{"query_string_blacklist"},
												Elem:          &schema.Schema{Type: schema.TypeString},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
