// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package resourceaction

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceAction) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validatePutRetryParameters(value *ResourceActionRetry) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validatePutTimeoutsParameters(value *ResourceActionTimeouts) error {
	return nil
}

func (r *jsiiProxy_ResourceAction) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateResourceAction_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateResourceAction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourceAction_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateResourceAction_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetActionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetBodyParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetHeadersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetIgnoreNotFoundParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetLocksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetMethodParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetQueryParametersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetResourceIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetResponseExportValuesParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetSensitiveBodyParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetSensitiveBodyVersionParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetSensitiveResponseExportValuesParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceAction) validateSetWhenParameters(val *string) error {
	return nil
}

func validateNewResourceActionParameters(scope constructs.Construct, id *string, config *ResourceActionConfig) error {
	return nil
}

