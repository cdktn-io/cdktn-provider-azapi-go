// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/provider/internal"
	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/providerfunctions"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs azapi}.
type AzapiProvider interface {
	cdktn.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AlwaysAcquirePolicyToken() interface{}
	SetAlwaysAcquirePolicyToken(val interface{})
	AlwaysAcquirePolicyTokenInput() interface{}
	AuxiliaryTenantIds() *[]*string
	SetAuxiliaryTenantIds(val *[]*string)
	AuxiliaryTenantIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientCertificatePassword() *string
	SetClientCertificatePassword(val *string)
	ClientCertificatePasswordInput() *string
	ClientCertificatePath() *string
	SetClientCertificatePath(val *string)
	ClientCertificatePathInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdFilePath() *string
	SetClientIdFilePath(val *string)
	ClientIdFilePathInput() *string
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretFilePath() *string
	SetClientSecretFilePath(val *string)
	ClientSecretFilePathInput() *string
	ClientSecretInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CustomCorrelationRequestId() *string
	SetCustomCorrelationRequestId(val *string)
	CustomCorrelationRequestIdInput() *string
	DefaultLocation() *string
	SetDefaultLocation(val *string)
	DefaultLocationInput() *string
	DefaultName() *string
	SetDefaultName(val *string)
	DefaultNameInput() *string
	DefaultTags() *map[string]*string
	SetDefaultTags(val *map[string]*string)
	DefaultTagsInput() *map[string]*string
	DisableCorrelationRequestId() interface{}
	SetDisableCorrelationRequestId(val interface{})
	DisableCorrelationRequestIdInput() interface{}
	DisableDefaultOutput() interface{}
	SetDisableDefaultOutput(val interface{})
	DisableDefaultOutputInput() interface{}
	DisableInstanceDiscovery() interface{}
	SetDisableInstanceDiscovery(val interface{})
	DisableInstanceDiscoveryInput() interface{}
	DisableTerraformPartnerId() interface{}
	SetDisableTerraformPartnerId(val interface{})
	DisableTerraformPartnerIdInput() interface{}
	EnablePreflight() interface{}
	SetEnablePreflight(val interface{})
	EnablePreflightInput() interface{}
	Endpoint() interface{}
	SetEndpoint(val interface{})
	EndpointInput() interface{}
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Provider-defined functions of the azapi provider.
	Functions() providerfunctions.AzapiProviderFunctions
	IgnoreNoOpChanges() interface{}
	SetIgnoreNoOpChanges(val interface{})
	IgnoreNoOpChangesInput() interface{}
	MaximumBusyRetryAttempts() *float64
	SetMaximumBusyRetryAttempts(val *float64)
	MaximumBusyRetryAttemptsInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	OidcAzureServiceConnectionId() *string
	SetOidcAzureServiceConnectionId(val *string)
	OidcAzureServiceConnectionIdInput() *string
	OidcRequestToken() *string
	SetOidcRequestToken(val *string)
	OidcRequestTokenInput() *string
	OidcRequestUrl() *string
	SetOidcRequestUrl(val *string)
	OidcRequestUrlInput() *string
	OidcToken() *string
	SetOidcToken(val *string)
	OidcTokenFilePath() *string
	SetOidcTokenFilePath(val *string)
	OidcTokenFilePathInput() *string
	OidcTokenInput() *string
	PartnerId() *string
	SetPartnerId(val *string)
	PartnerIdInput() *string
	PreserveResourceIdCasing() interface{}
	SetPreserveResourceIdCasing(val interface{})
	PreserveResourceIdCasingInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SkipProviderRegistration() interface{}
	SetSkipProviderRegistration(val interface{})
	SkipProviderRegistrationInput() interface{}
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	UseAksWorkloadIdentity() interface{}
	SetUseAksWorkloadIdentity(val interface{})
	UseAksWorkloadIdentityInput() interface{}
	UseCli() interface{}
	SetUseCli(val interface{})
	UseCliInput() interface{}
	UseMsi() interface{}
	SetUseMsi(val interface{})
	UseMsiInput() interface{}
	UseOidc() interface{}
	SetUseOidc(val interface{})
	UseOidcInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAlias()
	ResetAlwaysAcquirePolicyToken()
	ResetAuxiliaryTenantIds()
	ResetClientCertificate()
	ResetClientCertificatePassword()
	ResetClientCertificatePath()
	ResetClientId()
	ResetClientIdFilePath()
	ResetClientSecret()
	ResetClientSecretFilePath()
	ResetCustomCorrelationRequestId()
	ResetDefaultLocation()
	ResetDefaultName()
	ResetDefaultTags()
	ResetDisableCorrelationRequestId()
	ResetDisableDefaultOutput()
	ResetDisableInstanceDiscovery()
	ResetDisableTerraformPartnerId()
	ResetEnablePreflight()
	ResetEndpoint()
	ResetEnvironment()
	ResetIgnoreNoOpChanges()
	ResetMaximumBusyRetryAttempts()
	ResetOidcAzureServiceConnectionId()
	ResetOidcRequestToken()
	ResetOidcRequestUrl()
	ResetOidcToken()
	ResetOidcTokenFilePath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartnerId()
	ResetPreserveResourceIdCasing()
	ResetSkipProviderRegistration()
	ResetSubscriptionId()
	ResetTenantId()
	ResetUseAksWorkloadIdentity()
	ResetUseCli()
	ResetUseMsi()
	ResetUseOidc()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AzapiProvider
type jsiiProxy_AzapiProvider struct {
	internal.Type__cdktnTerraformProvider
}

func (j *jsiiProxy_AzapiProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) AlwaysAcquirePolicyToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysAcquirePolicyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) AlwaysAcquirePolicyTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysAcquirePolicyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) AuxiliaryTenantIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auxiliaryTenantIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) AuxiliaryTenantIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auxiliaryTenantIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientCertificatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientCertificatePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientIdFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientIdFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientSecretFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientSecretFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) CustomCorrelationRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCorrelationRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) CustomCorrelationRequestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCorrelationRequestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DefaultLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DefaultLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DefaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DefaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DefaultTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DefaultTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableCorrelationRequestId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorrelationRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableCorrelationRequestIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorrelationRequestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableDefaultOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDefaultOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableDefaultOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDefaultOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableInstanceDiscovery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableInstanceDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableInstanceDiscoveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableInstanceDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableTerraformPartnerId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTerraformPartnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) DisableTerraformPartnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTerraformPartnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) EnablePreflight() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePreflight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) EnablePreflightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePreflightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) Endpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) Functions() providerfunctions.AzapiProviderFunctions {
	var returns providerfunctions.AzapiProviderFunctions
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) IgnoreNoOpChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNoOpChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) IgnoreNoOpChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNoOpChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) MaximumBusyRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBusyRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) MaximumBusyRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBusyRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcAzureServiceConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcAzureServiceConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcAzureServiceConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcAzureServiceConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcRequestToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcRequestTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcRequestUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcRequestUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcTokenFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcTokenFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) OidcTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) PartnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) PartnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) PreserveResourceIdCasing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveResourceIdCasing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) PreserveResourceIdCasingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveResourceIdCasingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) SkipProviderRegistration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipProviderRegistration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) SkipProviderRegistrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipProviderRegistrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseAksWorkloadIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAksWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseAksWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAksWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseCli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseCliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseMsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseMsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseOidc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzapiProvider) UseOidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs azapi} Resource.
func NewAzapiProvider(scope constructs.Construct, id *string, config *AzapiProviderConfig) AzapiProvider {
	_init_.Initialize()

	if err := validateNewAzapiProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzapiProvider{}

	_jsii_.Create(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs azapi} Resource.
func NewAzapiProvider_Override(a AzapiProvider, scope constructs.Construct, id *string, config *AzapiProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzapiProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetAlwaysAcquirePolicyToken(val interface{}) {
	if err := j.validateSetAlwaysAcquirePolicyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysAcquirePolicyToken",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetAuxiliaryTenantIds(val *[]*string) {
	_jsii_.Set(
		j,
		"auxiliaryTenantIds",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientCertificate(val *string) {
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientCertificatePath(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePath",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientIdFilePath(val *string) {
	_jsii_.Set(
		j,
		"clientIdFilePath",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetClientSecretFilePath(val *string) {
	_jsii_.Set(
		j,
		"clientSecretFilePath",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetCustomCorrelationRequestId(val *string) {
	_jsii_.Set(
		j,
		"customCorrelationRequestId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDefaultLocation(val *string) {
	_jsii_.Set(
		j,
		"defaultLocation",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDefaultName(val *string) {
	_jsii_.Set(
		j,
		"defaultName",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDefaultTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDisableCorrelationRequestId(val interface{}) {
	if err := j.validateSetDisableCorrelationRequestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCorrelationRequestId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDisableDefaultOutput(val interface{}) {
	if err := j.validateSetDisableDefaultOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDefaultOutput",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDisableInstanceDiscovery(val interface{}) {
	if err := j.validateSetDisableInstanceDiscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableInstanceDiscovery",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetDisableTerraformPartnerId(val interface{}) {
	if err := j.validateSetDisableTerraformPartnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTerraformPartnerId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetEnablePreflight(val interface{}) {
	if err := j.validateSetEnablePreflightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePreflight",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetEndpoint(val interface{}) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetIgnoreNoOpChanges(val interface{}) {
	if err := j.validateSetIgnoreNoOpChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreNoOpChanges",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetMaximumBusyRetryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumBusyRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetOidcAzureServiceConnectionId(val *string) {
	_jsii_.Set(
		j,
		"oidcAzureServiceConnectionId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetOidcRequestToken(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestToken",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetOidcRequestUrl(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestUrl",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetOidcToken(val *string) {
	_jsii_.Set(
		j,
		"oidcToken",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetOidcTokenFilePath(val *string) {
	_jsii_.Set(
		j,
		"oidcTokenFilePath",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetPartnerId(val *string) {
	_jsii_.Set(
		j,
		"partnerId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetPreserveResourceIdCasing(val interface{}) {
	if err := j.validateSetPreserveResourceIdCasingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveResourceIdCasing",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetSkipProviderRegistration(val interface{}) {
	if err := j.validateSetSkipProviderRegistrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipProviderRegistration",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetUseAksWorkloadIdentity(val interface{}) {
	if err := j.validateSetUseAksWorkloadIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAksWorkloadIdentity",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetUseCli(val interface{}) {
	if err := j.validateSetUseCliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCli",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetUseMsi(val interface{}) {
	if err := j.validateSetUseMsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMsi",
		val,
	)
}

func (j *jsiiProxy_AzapiProvider)SetUseOidc(val interface{}) {
	if err := j.validateSetUseOidcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOidc",
		val,
	)
}

// Generates CDKTN code for importing a AzapiProvider resource upon running "cdktn plan <stack-name>".
func AzapiProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAzapiProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AzapiProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzapiProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzapiProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzapiProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzapiProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzapiProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzapiProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azapi.provider.AzapiProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzapiProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzapiProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzapiProvider) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AzapiProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetAlwaysAcquirePolicyToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAlwaysAcquirePolicyToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetAuxiliaryTenantIds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuxiliaryTenantIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientCertificatePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientIdFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientIdFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetClientSecretFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetCustomCorrelationRequestId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomCorrelationRequestId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDefaultLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDefaultName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDisableCorrelationRequestId() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableCorrelationRequestId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDisableDefaultOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableDefaultOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDisableInstanceDiscovery() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableInstanceDiscovery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetDisableTerraformPartnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableTerraformPartnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetEnablePreflight() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePreflight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetIgnoreNoOpChanges() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreNoOpChanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetMaximumBusyRetryAttempts() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumBusyRetryAttempts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetOidcAzureServiceConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcAzureServiceConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetOidcRequestToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcRequestToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetOidcRequestUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcRequestUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetOidcToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetOidcTokenFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcTokenFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetPartnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetPartnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetPreserveResourceIdCasing() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveResourceIdCasing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetSkipProviderRegistration() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipProviderRegistration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetTenantId() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetUseAksWorkloadIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetUseAksWorkloadIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetUseCli() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetUseMsi() {
	_jsii_.InvokeVoid(
		a,
		"resetUseMsi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) ResetUseOidc() {
	_jsii_.InvokeVoid(
		a,
		"resetUseOidc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzapiProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

