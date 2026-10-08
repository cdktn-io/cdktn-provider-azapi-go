// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzapiProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AzapiProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_AzapiProvider) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateAzapiProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAzapiProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAzapiProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAzapiProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetAlwaysAcquirePolicyTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetDisableCorrelationRequestIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetDisableDefaultOutputParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetDisableInstanceDiscoveryParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetDisableTerraformPartnerIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetEnablePreflightParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetEndpointParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetIgnoreNoOpChangesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetPreserveResourceIdCasingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetSkipProviderRegistrationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetUseAksWorkloadIdentityParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetUseCliParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetUseMsiParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzapiProvider) validateSetUseOidcParameters(val interface{}) error {
	return nil
}

func validateNewAzapiProviderParameters(scope constructs.Construct, id *string, config *AzapiProviderConfig) error {
	return nil
}

