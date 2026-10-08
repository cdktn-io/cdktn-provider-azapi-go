// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AzapiProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#alias AzapiProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Always acquire a policy token for write requests, regardless of whether one is required.
	//
	// The default is `false`. The default behaviour is to wait for a qualifying `403` response indicating that a policy token is required, and then retry the request with an acquired policy token. When this attribute is set to `true`, the provider proactively acquires a policy token and attaches it to every write request, avoiding the extra round-trip per request. Performance will be improved if the number of changed resources is known to be large beforehand. This can also be sourced from the `ARM_ALWAYS_ACQUIRE_POLICY_TOKEN` Environment Variable. See [Feature: Acquire Policy Token](guides/feature_acquire_policy_token.html) to learn more.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#always_acquire_policy_token AzapiProvider#always_acquire_policy_token}
	AlwaysAcquirePolicyToken interface{} `field:"optional" json:"alwaysAcquirePolicyToken" yaml:"alwaysAcquirePolicyToken"`
	// List of auxiliary Tenant IDs required for multi-tenancy and cross-tenant scenarios.
	//
	// This can also be sourced from the `ARM_AUXILIARY_TENANT_IDS` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#auxiliary_tenant_ids AzapiProvider#auxiliary_tenant_ids}
	AuxiliaryTenantIds *[]*string `field:"optional" json:"auxiliaryTenantIds" yaml:"auxiliaryTenantIds"`
	// A base64-encoded PKCS#12 bundle to be used as the client certificate for authentication.
	//
	// This can also be sourced from the `ARM_CLIENT_CERTIFICATE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_certificate AzapiProvider#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// The password associated with the Client Certificate. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PASSWORD` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_certificate_password AzapiProvider#client_certificate_password}
	ClientCertificatePassword *string `field:"optional" json:"clientCertificatePassword" yaml:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal which should be used.
	//
	// This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_certificate_path AzapiProvider#client_certificate_path}
	ClientCertificatePath *string `field:"optional" json:"clientCertificatePath" yaml:"clientCertificatePath"`
	// The Client ID which should be used. This can also be sourced from the `ARM_CLIENT_ID`, `AZURE_CLIENT_ID` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_id AzapiProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The path to a file containing the Client ID which should be used.
	//
	// This can also be sourced from the `ARM_CLIENT_ID_FILE_PATH` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_id_file_path AzapiProvider#client_id_file_path}
	ClientIdFilePath *string `field:"optional" json:"clientIdFilePath" yaml:"clientIdFilePath"`
	// The Client Secret which should be used. This can also be sourced from the `ARM_CLIENT_SECRET` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_secret AzapiProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The path to a file containing the Client Secret which should be used.
	//
	// For use When authenticating as a Service Principal using a Client Secret. This can also be sourced from the `ARM_CLIENT_SECRET_FILE_PATH` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#client_secret_file_path AzapiProvider#client_secret_file_path}
	ClientSecretFilePath *string `field:"optional" json:"clientSecretFilePath" yaml:"clientSecretFilePath"`
	// The value of the `x-ms-correlation-request-id` header, otherwise an auto-generated UUID will be used.
	//
	// This can also be sourced from the `ARM_CORRELATION_REQUEST_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#custom_correlation_request_id AzapiProvider#custom_correlation_request_id}
	CustomCorrelationRequestId *string `field:"optional" json:"customCorrelationRequestId" yaml:"customCorrelationRequestId"`
	// The default Azure Region where the azure resource should exist.
	//
	// The `location` in each resource block can override the `default_location`. Changing this forces new resources to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#default_location AzapiProvider#default_location}
	DefaultLocation *string `field:"optional" json:"defaultLocation" yaml:"defaultLocation"`
	// The default name to create the azure resource.
	//
	// The `name` in each resource block can override the `default_name`. Changing this forces new resources to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#default_name AzapiProvider#default_name}
	DefaultName *string `field:"optional" json:"defaultName" yaml:"defaultName"`
	// A mapping of tags which should be assigned to the azure resource as default tags.
	//
	// The `tags` in each resource block can override the `default_tags`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#default_tags AzapiProvider#default_tags}
	DefaultTags *map[string]*string `field:"optional" json:"defaultTags" yaml:"defaultTags"`
	// This will disable the x-ms-correlation-request-id header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#disable_correlation_request_id AzapiProvider#disable_correlation_request_id}
	DisableCorrelationRequestId interface{} `field:"optional" json:"disableCorrelationRequestId" yaml:"disableCorrelationRequestId"`
	// Disable default output.
	//
	// The default is false. When set to false, the provider will output the read-only properties if `response_export_values` is not specified in the resource block. When set to true, the provider will disable this output. This can also be sourced from the `ARM_DISABLE_DEFAULT_OUTPUT` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#disable_default_output AzapiProvider#disable_default_output}
	DisableDefaultOutput interface{} `field:"optional" json:"disableDefaultOutput" yaml:"disableDefaultOutput"`
	// Disables Instance Discovery, which validates that the Authority is valid and known by the Microsoft Entra instance metadata service at `https://login.microsoft.com` before authenticating. This should only be enabled when the configured authority is known to be valid and trustworthy - such as when running against Azure Stack or when `environment` is set to `custom`. This can also be specified via the `ARM_DISABLE_INSTANCE_DISCOVERY` environment variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#disable_instance_discovery AzapiProvider#disable_instance_discovery}
	DisableInstanceDiscovery interface{} `field:"optional" json:"disableInstanceDiscovery" yaml:"disableInstanceDiscovery"`
	// Disable sending the Terraform Partner ID if a custom `partner_id` isn't specified, which allows Microsoft to better understand the usage of Terraform.
	//
	// The Partner ID does not give HashiCorp any direct access to usage information. This can also be sourced from the `ARM_DISABLE_TERRAFORM_PARTNER_ID` environment variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#disable_terraform_partner_id AzapiProvider#disable_terraform_partner_id}
	DisableTerraformPartnerId interface{} `field:"optional" json:"disableTerraformPartnerId" yaml:"disableTerraformPartnerId"`
	// Enable Preflight Validation.
	//
	// The default is false. When set to true, the provider will use Preflight to do static validation before really deploying a new resource. When set to false, the provider will disable this validation. This can also be sourced from the `ARM_ENABLE_PREFLIGHT` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#enable_preflight AzapiProvider#enable_preflight}
	EnablePreflight interface{} `field:"optional" json:"enablePreflight" yaml:"enablePreflight"`
	// The Azure API Endpoint Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#endpoint AzapiProvider#endpoint}
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The Cloud Environment which should be used.
	//
	// Defaults to `public`. This can also be sourced from the `ARM_ENVIRONMENT` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#environment AzapiProvider#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Ignore no-op changes for `azapi_resource`.
	//
	// The default is true. When set to true, the provider will suppress changes in the `azapi_resource` if the `body` in the new API version still matches the remote state. When set to false, the provider will not suppress these changes. This can also be sourced from the `ARM_IGNORE_NO_OP_CHANGES` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#ignore_no_op_changes AzapiProvider#ignore_no_op_changes}
	IgnoreNoOpChanges interface{} `field:"optional" json:"ignoreNoOpChanges" yaml:"ignoreNoOpChanges"`
	// DEPRECATED - The maximum number of retries to attempt if the Azure API returns an HTTP 408, 429, 500, 502, 503, or 504 response.
	//
	// The default is `32767`, this allows the provider to rely on the resource timeout values rather than a maximum retry count. The resource-specific retry configuration may additionally be used to retry on other errors and conditions. This property will be removed in a future version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#maximum_busy_retry_attempts AzapiProvider#maximum_busy_retry_attempts}
	MaximumBusyRetryAttempts *float64 `field:"optional" json:"maximumBusyRetryAttempts" yaml:"maximumBusyRetryAttempts"`
	// The Azure Pipelines Service Connection ID to use for authentication.
	//
	// This can also be sourced from the `ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID`, `ARM_OIDC_AZURE_SERVICE_CONNECTION_ID`, or `AZURESUBSCRIPTION_SERVICE_CONNECTION_ID` Environment Variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#oidc_azure_service_connection_id AzapiProvider#oidc_azure_service_connection_id}
	OidcAzureServiceConnectionId *string `field:"optional" json:"oidcAzureServiceConnectionId" yaml:"oidcAzureServiceConnectionId"`
	// The bearer token for the request to the OIDC provider.
	//
	// This can also be sourced from the `ARM_OIDC_REQUEST_TOKEN`, `ACTIONS_ID_TOKEN_REQUEST_TOKEN`, or `SYSTEM_ACCESSTOKEN` Environment Variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#oidc_request_token AzapiProvider#oidc_request_token}
	OidcRequestToken *string `field:"optional" json:"oidcRequestToken" yaml:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token.
	//
	// This can also be sourced from the `ARM_OIDC_REQUEST_URL`, `ACTIONS_ID_TOKEN_REQUEST_URL`, or `SYSTEM_OIDCREQUESTURI` Environment Variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#oidc_request_url AzapiProvider#oidc_request_url}
	OidcRequestUrl *string `field:"optional" json:"oidcRequestUrl" yaml:"oidcRequestUrl"`
	// The ID token when authenticating using OpenID Connect (OIDC). This can also be sourced from the `ARM_OIDC_TOKEN` environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#oidc_token AzapiProvider#oidc_token}
	OidcToken *string `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// The path to a file containing an ID token when authenticating using OpenID Connect (OIDC).
	//
	// This can also be sourced from the `ARM_OIDC_TOKEN_FILE_PATH`, `AZURE_FEDERATED_TOKEN_FILE` environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#oidc_token_file_path AzapiProvider#oidc_token_file_path}
	OidcTokenFilePath *string `field:"optional" json:"oidcTokenFilePath" yaml:"oidcTokenFilePath"`
	// A GUID/UUID that is [registered](https://docs.microsoft.com/azure/marketplace/azure-partner-customer-usage-attribution#register-guids-and-offers) with Microsoft to facilitate partner resource usage attribution. This can also be sourced from the `ARM_PARTNER_ID` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#partner_id AzapiProvider#partner_id}
	PartnerId *string `field:"optional" json:"partnerId" yaml:"partnerId"`
	// Preserve the existing casing of the resource ID in state.
	//
	// The default is false. When set to true, if the resource ID the provider would write back to state differs from the value already in state only by casing, the existing casing is kept. This is useful when consumers of the resource ID (or the `azapi_resource` identity) require a specific casing that the Azure API may not preserve. This only affects the `id` (and `resource_id`) attributes; other properties are unaffected. This can also be sourced from the `ARM_PRESERVE_RESOURCE_ID_CASING` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#preserve_resource_id_casing AzapiProvider#preserve_resource_id_casing}
	PreserveResourceIdCasing interface{} `field:"optional" json:"preserveResourceIdCasing" yaml:"preserveResourceIdCasing"`
	// Should the Provider skip registering the Resource Providers it supports?
	//
	// This can also be sourced from the `ARM_SKIP_PROVIDER_REGISTRATION` Environment Variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#skip_provider_registration AzapiProvider#skip_provider_registration}
	SkipProviderRegistration interface{} `field:"optional" json:"skipProviderRegistration" yaml:"skipProviderRegistration"`
	// The Subscription ID which should be used. This can also be sourced from the `ARM_SUBSCRIPTION_ID` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#subscription_id AzapiProvider#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The Tenant ID should be used. This can also be sourced from the `ARM_TENANT_ID` Environment Variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#tenant_id AzapiProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Should AKS Workload Identity be used for Authentication?
	//
	// This can also be sourced from the `ARM_USE_AKS_WORKLOAD_IDENTITY` Environment Variable. Defaults to `false`. When set, `client_id`, `tenant_id` and `oidc_token_file_path` will be detected from the environment and do not need to be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#use_aks_workload_identity AzapiProvider#use_aks_workload_identity}
	UseAksWorkloadIdentity interface{} `field:"optional" json:"useAksWorkloadIdentity" yaml:"useAksWorkloadIdentity"`
	// Should Azure CLI be used for authentication?
	//
	// This can also be sourced from the `ARM_USE_CLI` environment variable. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#use_cli AzapiProvider#use_cli}
	UseCli interface{} `field:"optional" json:"useCli" yaml:"useCli"`
	// Should Managed Identity be used for Authentication?
	//
	// This can also be sourced from the `ARM_USE_MSI` Environment Variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#use_msi AzapiProvider#use_msi}
	UseMsi interface{} `field:"optional" json:"useMsi" yaml:"useMsi"`
	// Should OIDC be used for Authentication? This can also be sourced from the `ARM_USE_OIDC` Environment Variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#use_oidc AzapiProvider#use_oidc}
	UseOidc interface{} `field:"optional" json:"useOidc" yaml:"useOidc"`
}

