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
			Create: schema.DefaultTimeout(10 * time.Minute), //nolint:gomnd
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: resourceBackendConfigSchemaV1(),
	}
}

//nolint:funlen
func resourceBackendConfigSchemaV1() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": namespacedMetadataSchema("backendconfig", false),
		"spec": {
			Type:        schema.TypeList,
			Description: "Spec defines the specification of the desired behavior of the backendconfig. More info: https://cloud.google.com/kubernetes-engine/docs/how-to/ingress-features#configuring_ingress_features_through_backendconfig_parameters",
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
					"cdn": {
						Type:        schema.TypeList,
						Description: "TODO: CDN configuration",
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:        schema.TypeBool,
									Description: "TODO: If set to true, Cloud CDN is enabled for this Ingress backend.",
									Optional:    true,
								},
								"cache_policy": {
									Type:        schema.TypeList,
									Description: "TODO: configure cache policy",
									Optional:    true,
									MaxItems:    1,
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
												Elem:          &schema.Schema{Type: schema.TypeString, Required: true},
												Set:           schema.HashString,
											},
											"query_string_whitelist": {
												Type:          schema.TypeSet,
												Description:   "TODO: Specify a string array with the names of query string parameters to include in cache keys. All other parameters are excluded. You can query_string_blacklist or query_string_whitelist, but not both.",
												Optional:      true,
												ConflictsWith: []string{"query_string_blacklist"},
												Elem:          &schema.Schema{Type: schema.TypeString, Required: true},
												Set:           schema.HashString,
											},
										},
									},
								},
							},
						},
					},
					"connection_draining": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"draining_timeout_sec": {
									Type:        schema.TypeInt,
									Description: "TODO: Connection draining timeout is the time, in seconds, to wait for connections to drain. For the specified duration of the timeout, existing requests to the removed backend are given time to complete. The load balancer does not send new requests to the removed backend. After the timeout duration is reached, all remaining connections to the backend are closed. The timeout duration can be from 0 to 3600 seconds. The default value is 0, which also disables connection draining.",
									Required:    true,
								},
							},
						},
					},
					"health_check": {
						Type:        schema.TypeList,
						Description: "TODO: If set, these parameters supersede the Kubernetes readiness probe settings and health check default values.",
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"check_interval_sec": {
									Type:        schema.TypeInt,
									Description: "TODO: Specify the check-interval, in seconds, for each health check prober. This is the time from the start of one prober's check to the start of its next check. If you omit this parameter, the Google Cloud default of 5 seconds is used.",
									Optional:    true,
								},
								"timeout_sec": {
									Type:        schema.TypeInt,
									Description: "TODO: Specify the amount of time that Google Cloud waits for a response to a probe. The value of timeout must be less than or equal to the interval. Units are seconds. Each probe requires an HTTP 200 (OK) response code to be delivered before the probe timeout.",
									Optional:    true,
								},
								"healthy_threshold": {
									Type:        schema.TypeInt,
									Description: "TODO: Specify the number of sequential connection attempts that must succeed or fail, for at least one prober, in order to change the health state from healthy to unhealthy or vice versa. If you omit one of these parameters, Google Cloud uses the default value of 2.",
									Optional:    true,
								},
								"unhealthy_threshold": {
									Type:        schema.TypeInt,
									Description: "TODO: Specify the number of sequential connection attempts that must succeed or fail, for at least one prober, in order to change the health state from healthy to unhealthy or vice versa. If you omit one of these parameters, Google Cloud uses the default value of 2.",
									Optional:    true,
								},
								"type": {
									Type:        schema.TypeString,
									Description: "TODO: Specify a protocol used by probe systems for health checking. The BackendConfig only supports creating health checks using the HTTP, HTTPS, or HTTP2 protocols. For more information, see Success criteria for HTTP, HTTPS, and HTTP/2. You cannot omit this parameter.",
									Required:    true,
								},
								"request_path": {
									Type:        schema.TypeString,
									Description: "For HTTP, HTTPS, or HTTP2 health checks, specify the request-path to which the probe system should connect. If you omit this parameter, Google Cloud uses the default of /.",
									Optional:    true,
								},
								"port": {
									Type:        schema.TypeInt,
									Description: "Specifies the port by using a port number. If you omit this parameter, Google Cloud uses the default of 80.",
									Optional:    true,
								},
							},
						},
					},
					"security_policy": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Description: "Add the name of your security policy to the BackendConfig. ",
									Required:    true,
								},
							},
						},
					},
					"logging": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable": {
									Type:        schema.TypeBool,
									Description: "TODO: If set to true, access logging will be enabled for this Ingress and logs is available in Cloud Logging. Otherwise, access logging is disabled for this Ingress.",
									Required:    true,
								},
								"sample_rate": {
									Type:        schema.TypeFloat,
									Description: "TODO: Specify a value from 0.0 through 1.0, where 0.0 means no packets are logged and 1.0 means 100% of packets are logged. This field is only relevant if enable is set to true. sampleRate is an optional field, but if it's configured then enable: true must also be set or else it is interpreted as enable: false",
									Optional:    true,
								},
							},
						},
					},
					"iap": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:        schema.TypeBool,
									Required:    true,
									Description: "TODO: To configure the BackendConfig for Identity-Aware Proxy IAP, you need to specify the enabled and secretName values to your iap block in your BackendConfig. To specify these values, ensure that you have the compute.backendServices.update permission.",
								},
								"oauthclient_credentials_secret_name": {
									Type:         schema.TypeString,
									RequiredWith: []string{"enabled"},
									Description:  "TODO: To configure the BackendConfig for Identity-Aware Proxy IAP, you need to specify the enabled and secretName values to your iap block in your BackendConfig. To specify these values, ensure that you have the compute.backendServices.update permission.",
								},
							},
						},
					},
					"session_affinity": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "TODO: to set session affinity to client IP or generated cookie.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"affinity_type": {
									Type:        schema.TypeString,
									Description: "TODO: set affinityType to GENERATED_COOKIE or CLIENT_IP",
									Required:    true,
								},
								"affinity_cookie_ttl_sec": {
									Type:        schema.TypeInt,
									Description: "TODO: To use a BackendConfig to set generated cookie affinity , set affinityType to GENERATED_COOKIE in your BackendConfig manifest. You can also use affinityCookieTtlSec to set the time period for the cookie to remain active.",
									Optional:    true,
								},
							},
						},
					},
					"custom_request_headers": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"headers": {
									Type:     schema.TypeSet,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString, Required: true},
									Set:      schema.HashString,
								},
							},
						},
					},
				},
			},
		},
	}
}
