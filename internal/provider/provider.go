// providers.go
package provider

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/accountgroups"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/accounts"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/activationcode"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/advancedcomputersearches"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/advancedmobiledevicesearches"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/advancedusersearches"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/allowedfileextensions"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/apiintegrations"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/apiroles"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/buildings"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/categories"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/computercheckin"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/computerextensionattributes"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/computerinventory"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/computerinventorycollection"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/computerprestageenrollments"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/departments"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/diskencryptionconfigurations"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/dockitems"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/filesharedistributionpoints"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/macosconfigurationprofilesplist"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/mobiledeviceconfigurationprofilesplist"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/networksegments"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/policies"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/printers"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/restrictedsoftware"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/scripts"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/sites"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/smartcomputergroups"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/staticcomputergroups"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/usergroups"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/webhooks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// TerraformProviderProductUserAgent is included in the User-Agent header for
// any API requests made by the provider.
const (
	terraformProviderProductUserAgent = "terraform-provider-jamfpro"
	envKeyOAuthClientId               = "JAMFPRO_CLIENT_ID"
	envKeyOAuthClientSecret           = "JAMFPRO_CLIENT_SECRET"
	envKeyBasicAuthUsername           = "JAMFPRO_BASIC_USERNAME"
	envKeyBasicAuthPassword           = "JAMFPRO_BASIC_PASSWORD"
	envKeyJamfProUrlRoot              = "JAMFPRO_URL_ROOT" // e.g https://yourcompany.jamfcloud.com
	jamfLoadBalancerCookieName        = "jpro-ingress"
)

/*
GetJamfFqdn retrieves the instance domain name from the provided schema resource data.

If the instance domain is not found, it appends an error diagnostic to the diagnostics slice.

Parameters:

	d      - A pointer to the schema.ResourceData object which contains the resource data.
	diags  - A pointer to a slice of diag.Diagnostics where error messages will be appended.

Returns:

	A string representing the instance domain name. If the instance domain name is not provided,
	an error diagnostic is appended to diags and an empty string is returned.
*/
func GetJamfFqdn(d *schema.ResourceData, diags *diag.Diagnostics) string {
	jamf_fqdn, ok := d.GetOk("jamf_instance_fqdn")
	if jamf_fqdn.(string) == "" || !ok {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error getting instance name",
			Detail:   "instance_name must be provided either as an environment variable (JAMFPRO_INSTANCE_NAME) or in the Terraform configuration",
		})
		return ""
	}
	return jamf_fqdn.(string)
}

/*
GetClientID retrieves the client ID from the provided schema resource data.
If the client ID is not found, it appends an error diagnostic to the diagnostics slice.

Parameters:

	d      - A pointer to the schema.ResourceData object which contains the resource data.
	diags  - A pointer to a slice of diag.Diagnostics where error messages will be appended.

Returns:

	A string representing the client ID. If the client ID is not provided,
	an error diagnostic is appended to diags and an empty string is returned.
*/
func GetClientID(d *schema.ResourceData, diags *diag.Diagnostics) string {
	clientID := d.Get("client_id").(string)
	if clientID == "" {

		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error getting client id",
			Detail:   "client_id must be provided either as an environment variable (JAMFPRO_CLIENT_ID) or in the Terraform configuration",
		})

		return ""

	}
	return clientID
}

/*
GetClientSecret retrieves the client secret from the provided schema resource data.
If the client ID is not found, it appends an error diagnostic to the diagnostics slice.

Parameters:

	d      - A pointer to the schema.ResourceData object which contains the resource data.
	diags  - A pointer to a slice of diag.Diagnostics where error messages will be appended.

Returns:

	A string representing the client ID. If the client ID is not provided,
	an error diagnostic is appended to diags and an empty string is returned.
*/
func GetClientSecret(d *schema.ResourceData, diags *diag.Diagnostics) string {
	clientSecret := d.Get("client_secret").(string)
	if clientSecret == "" {

		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error getting client secret",
			Detail:   "client_secret must be provided either as an environment variable (JAMFPRO_CLIENT_SECRET) or in the Terraform configuration",
		})

		return ""

	}
	return clientSecret
}

/*
GetBasicAuthUsername retrieves the basic auth username from the provided schema resource data.
If the client ID is not found, it appends an error diagnostic to the diagnostics slice.

Parameters:

	d      - A pointer to the schema.ResourceData object which contains the resource data.
	diags  - A pointer to a slice of diag.Diagnostics where error messages will be appended.

Returns:

	A string representing the client ID. If the client ID is not provided,
	an error diagnostic is appended to diags and an empty string is returned.
*/
func GetBasicAuthUsername(d *schema.ResourceData, diags *diag.Diagnostics) string {
	username := d.Get("username").(string)
	if username == "" {

		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error getting basic auth username",
			Detail:   "username must be provided either as an environment variable (JAMFPRO_USERNAME) or in the Terraform configuration",
		})

		return ""

	}
	return username
}

/*
GetBasicAuthPassword retrieves the basic auth password from the provided schema resource data.
If the client ID is not found, it appends an error diagnostic to the diagnostics slice.

Parameters:

	d      - A pointer to the schema.ResourceData object which contains the resource data.
	diags  - A pointer to a slice of diag.Diagnostics where error messages will be appended.

Returns:

	A string representing the client ID. If the client ID is not provided,
	an error diagnostic is appended to diags and an empty string is returned.
*/
func GetBasicAuthPassword(d *schema.ResourceData, diags *diag.Diagnostics) string {
	password := d.Get("password").(string)
	if password == "" {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error getting basic auth password",
			Detail:   "password must be provided either as an environment variable (JAMFPRO_PASSWORD) or in the Terraform configuration",
		})

		return ""

	}
	return password
}

// Schema defines the configuration attributes for the  within the JamfPro provider.
func Provider() *schema.Provider {

	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"jamf_instance_fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(envKeyJamfProUrlRoot, ""),
				Description: "The Jamf Pro FQDN (fully qualified domain name). example: https://mycompany.jamfcloud.com",
			},
			"auth_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Auth method chosen for Jamf.",
				ValidateFunc: validation.StringInSlice([]string{
					"basic", "oauth2",
				}, true),
			},
			// TODO the descs below
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(envKeyOAuthClientSecret, ""),
				Description: "The Jamf Pro Client ID for authentication.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(envKeyOAuthClientSecret, ""),
				Description: "The Jamf Pro Client secret for authentication.",
			},
			"basic_auth_username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(envKeyBasicAuthUsername, ""),
				Description: "The Jamf Pro username used for authentication.",
			},
			"basic_auth_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(envKeyBasicAuthPassword, ""),
				Description: "The Jamf Pro password used for authentication.",
			},
			"log_level": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "warning",
				ValidateFunc: validation.StringInSlice([]string{
					"debug", "info", "warning", "none",
				}, false),
				Description: "The logging level: debug, info, warning, or none",
			},
			"log_output_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "pretty",
				Description: "The output format of the logs. Use 'JSON' for JSON format, 'console' for human-readable format. Defaults to console if no value is supplied.",
			},
			"log_console_separator": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     " ",
				Description: "The separator character used in console log output.",
			},
			"log_export_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the path to export http client logs to.",
			},
			"export_logs": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "export logs to file",
			},
			"hide_sensitive_data": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Define whether sensitive fields should be hidden in logs. Default to hiding sensitive data in logs",
			},
			"custom_cookies": {
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "cookie key",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "cookie value",
						},
					},
				},
			},
			"jamf_load_balancer_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "programatically determines all available web app members in the load balance and locks all instances of httpclient to the app for faster executions. \nTEMP SOLUTION UNTIL JAMF PROVIDES SOLUTION",
			},
			"token_refresh_buffer_period_seconds": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     300,
				Description: "The buffer period in seconds for token refresh.",
			},

			"mandatory_request_delay_milliseconds": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     100,
				Description: "a mandatory delay after each request before returning to reduce high volume of requests in a short time",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{

			"jamfpro_account":                                   accounts.DataSourceJamfProAccounts(),
			"jamfpro_account_group":                             accountgroups.DataSourceJamfProAccountGroups(),
			"jamfpro_advanced_computer_search":                  advancedcomputersearches.DataSourceJamfProAdvancedComputerSearches(),
			"jamfpro_advanced_mobile_device_search":             advancedmobiledevicesearches.DataSourceJamfProAdvancedMobileDeviceSearches(),
			"jamfpro_advanced_user_search":                      advancedusersearches.DataSourceJamfProAdvancedUserSearches(),
			"jamfpro_api_integration":                           apiintegrations.DataSourceJamfProApiIntegrations(),
			"jamfpro_api_role":                                  apiroles.DataSourceJamfProAPIRoles(),
			"jamfpro_building":                                  buildings.DataSourceJamfProBuildings(),
			"jamfpro_category":                                  categories.DataSourceJamfProCategories(),
			"jamfpro_computer_extension_attribute":              computerextensionattributes.DataSourceJamfProComputerExtensionAttributes(),
			"jamfpro_computer_inventory":                        computerinventory.DataSourceJamfProComputerInventory(),
			"jamfpro_computer_prestage_enrollment":              computerprestageenrollments.DataSourceJamfProComputerPrestageEnrollmentEnrollment(),
			"jamfpro_department":                                departments.DataSourceJamfProDepartments(),
			"jamfpro_disk_encryption_configuration":             diskencryptionconfigurations.DataSourceJamfProDiskEncryptionConfigurations(),
			"jamfpro_dock_item":                                 dockitems.DataSourceJamfProDockItems(),
			"jamfpro_file_share_distribution_point":             filesharedistributionpoints.DataSourceJamfProFileShareDistributionPoints(),
			"jamfpro_network_segment":                           networksegments.DataSourceJamfProNetworkSegments(),
			"jamfpro_macos_configuration_profile_plist":         macosconfigurationprofilesplist.DataSourceJamfProMacOSConfigurationProfilesPlist(),
			"jamfpro_mobile_device_configuration_profile_plist": mobiledeviceconfigurationprofilesplist.DataSourceJamfProMobileDeviceConfigurationProfilesPlist(),
			// "jamfpro_package":                                   packages.DataSourceJamfProPackages(),
			"jamfpro_printer":               printers.DataSourceJamfProPrinters(),
			"jamfpro_script":                scripts.DataSourceJamfProScripts(),
			"jamfpro_site":                  sites.DataSourceJamfProSites(),
			"jamfpro_smart_computer_group":  smartcomputergroups.DataSourceJamfProSmartComputerGroups(),
			"jamfpro_static_computer_group": staticcomputergroups.DataSourceJamfProStaticComputerGroups(),
			"jamfpro_restricted_software":   restrictedsoftware.DataSourceJamfProRestrictedSoftwares(),
			"jamfpro_user_group":            usergroups.DataSourceJamfProUserGroups(),
			"jamfpro_webhook":               webhooks.DataSourceJamfProWebhooks(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"jamfpro_account":                                   accounts.ResourceJamfProAccounts(),
			"jamfpro_account_group":                             accountgroups.ResourceJamfProAccountGroups(),
			"jamfpro_activation_code":                           activationcode.ResourceJamfProActivationCode(),
			"jamfpro_advanced_computer_search":                  advancedcomputersearches.ResourceJamfProAdvancedComputerSearches(),
			"jamfpro_advanced_mobile_device_search":             advancedmobiledevicesearches.ResourceJamfProAdvancedMobileDeviceSearches(),
			"jamfpro_advanced_user_search":                      advancedusersearches.ResourceJamfProAdvancedUserSearches(),
			"jamfpro_allowed_file_extension":                    allowedfileextensions.ResourceJamfProAllowedFileExtensions(),
			"jamfpro_api_integration":                           apiintegrations.ResourceJamfProApiIntegrations(),
			"jamfpro_api_role":                                  apiroles.ResourceJamfProAPIRoles(),
			"jamfpro_building":                                  buildings.ResourceJamfProBuildings(),
			"jamfpro_category":                                  categories.ResourceJamfProCategories(),
			"jamfpro_computer_checkin":                          computercheckin.ResourceJamfProComputerCheckin(),
			"jamfpro_computer_extension_attribute":              computerextensionattributes.ResourceJamfProComputerExtensionAttributes(),
			"jamfpro_computer_inventory_collection":             computerinventorycollection.ResourceJamfProComputerInventoryCollection(),
			"jamfpro_computer_prestage_enrollment":              computerprestageenrollments.ResourceJamfProComputerPrestageEnrollmentEnrollment(),
			"jamfpro_department":                                departments.ResourceJamfProDepartments(),
			"jamfpro_disk_encryption_configuration":             diskencryptionconfigurations.ResourceJamfProDiskEncryptionConfigurations(),
			"jamfpro_dock_item":                                 dockitems.ResourceJamfProDockItems(),
			"jamfpro_file_share_distribution_point":             filesharedistributionpoints.ResourceJamfProFileShareDistributionPoints(),
			"jamfpro_network_segment":                           networksegments.ResourceJamfProNetworkSegments(),
			"jamfpro_macos_configuration_profile_plist":         macosconfigurationprofilesplist.ResourceJamfProMacOSConfigurationProfilesPlist(),
			"jamfpro_mobile_device_configuration_profile_plist": mobiledeviceconfigurationprofilesplist.ResourceJamfProMobileDeviceConfigurationProfilesPlist(),
			// "jamfpro_package":                                   packages.ResourceJamfProPackages(),
			"jamfpro_policy":                policies.ResourceJamfProPolicies(),
			"jamfpro_printer":               printers.ResourceJamfProPrinters(),
			"jamfpro_script":                scripts.ResourceJamfProScripts(),
			"jamfpro_site":                  sites.ResourceJamfProSites(),
			"jamfpro_smart_computer_group":  smartcomputergroups.ResourceJamfProSmartComputerGroups(),
			"jamfpro_static_computer_group": staticcomputergroups.ResourceJamfProStaticComputerGroups(),
			"jamfpro_restricted_software":   restrictedsoftware.ResourceJamfProRestrictedSoftwares(),
			"jamfpro_user_group":            usergroups.ResourceJamfProUserGroups(),
			"jamfpro_webhook":               webhooks.ResourceJamfProWebhooks(),
		},
	}

	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var err error
		var diags diag.Diagnostics
		var sharedLogger logger.Logger
		var jamfIntegration *jamfprointegration.Integration
		var jamfDomain,
			clientId,
			clientSecret,
			basicAuthUsername,
			basicAuthPassword string

		// Logger
		parsedLogLevel := logger.ParseLogLevelFromString(d.Get("log_level").(string))
		logOutputFormat := d.Get("log_output_format").(string)
		logConsoleSeparator := d.Get("log_console_separator").(string)
		logFilePath := d.Get("log_export_path").(string)
		exportLogs := d.Get("export_logs").(bool)

		sharedLogger = logger.BuildLogger(
			parsedLogLevel,
			logOutputFormat,
			logConsoleSeparator,
			logFilePath,
			exportLogs,
		)

		// Auth
		jamfDomain = GetJamfFqdn(d, &diags)
		tokenRefrshBufferPeriod := time.Duration(d.Get("token_refresh_buffer_period_seconds").(int)) * time.Second

		switch d.Get("auth_method").(string) {
		case "oauth2":
			clientId = GetClientID(d, &diags)
			clientSecret = GetClientSecret(d, &diags)
			jamfIntegration, err = jamfprointegration.BuildIntegrationWithOAuth(
				jamfDomain,
				sharedLogger,
				tokenRefrshBufferPeriod,
				clientId,
				clientSecret,
			)

		case "basic":
			basicAuthUsername = GetBasicAuthUsername(d, &diags)
			basicAuthPassword = GetBasicAuthPassword(d, &diags)
			jamfIntegration, err = jamfprointegration.BuildIntegrationWithBasicAuth(
				jamfDomain,
				sharedLogger,
				tokenRefrshBufferPeriod,
				basicAuthUsername,
				basicAuthPassword,
			)

		default:
			return nil, append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "invalid auth method supplied",
				Detail:   "You should not be able to find this error. If you have, please raise an issue with the schema.",
			})

		}

		if err != nil {
			return nil, append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error building jamf integration",
				Detail:   fmt.Sprintf("error: %v", err),
			})
		}

		// Cookies
		var cookiesList []*http.Cookie
		load_balancer_lock_enabled := d.Get("jamf_load_balancer_lock").(bool)
		customCookies := d.Get("custom_cookies")

		if load_balancer_lock_enabled {
			cookies, err := jamfIntegration.GetSessionCookies()
			if err != nil {
				return nil, append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Error getting session cookies",
					Detail:   fmt.Sprintf("error: %v", err),
				})
			}

			cookiesList = append(cookiesList, cookies...)

		}

		if customCookies != nil && len(customCookies.([]interface{})) > 0 {
			for _, v := range customCookies.([]interface{}) {
				name := v.(map[string]interface{})["name"]
				value := v.(map[string]interface{})["value"]

				if name == jamfLoadBalancerCookieName && load_balancer_lock_enabled {
					return nil, append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Cannot have load balancer lock and custom cookie of same name. (jpro-ingress)",
					})
				}

				httpCookie := &http.Cookie{
					Name:  name.(string),
					Value: value.(string),
				}

				cookiesList = append(cookiesList, httpCookie)
			}
		}

		// Amend timeouts
		// TODO make this exclusions list a lot prettier.
		// excludedResource := []string{"jamfpro_package"}
		for key, r := range provider.ResourcesMap {
			if key != "jamfpro_package" && key != "jamfpro_jamfpro_static_computer_group" && key != "jamfpro_smart_computer_group" {
				*r.Timeouts.Create = GetDefaultContextTimeoutCreate(load_balancer_lock_enabled)
				*r.Timeouts.Read = GetDefaultContextTimeoutRead(load_balancer_lock_enabled)
				*r.Timeouts.Update = GetDefaultContextTimeoutUpdate(load_balancer_lock_enabled)
				*r.Timeouts.Delete = GetDefaultContextTimeoutDelete(load_balancer_lock_enabled)
			}
		}

		// Packaging
		config := httpclient.ClientConfig{
			Integration:              jamfIntegration,
			HideSensitiveData:        d.Get("hide_sensitive_data").(bool),
			TokenRefreshBufferPeriod: tokenRefrshBufferPeriod,
			CustomCookies:            cookiesList,
			MandatoryRequestDelay:    time.Duration(d.Get("mandatory_request_delay_milliseconds").(int)) * time.Millisecond,
			RetryEligiableRequests:   false, // Forced off for now
		}

		goHttpClient, err := httpclient.BuildClient(config, false, sharedLogger)
		if err != nil {
			return nil, append(diags, diag.FromErr(err)...)
		}

		jamfClient := jamfpro.Client{
			HTTP: goHttpClient,
		}

		return &jamfClient, diags
	}

	return provider
}
