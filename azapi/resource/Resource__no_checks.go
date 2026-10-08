// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package resource

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Resource) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (r *jsiiProxy_Resource) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_Resource) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validatePutIdentityParameters(value interface{}) error {
	return nil
}

func (r *jsiiProxy_Resource) validatePutRetryParameters(value *ResourceRetry) error {
	return nil
}

func (r *jsiiProxy_Resource) validatePutTimeoutsParameters(value *ResourceTimeouts) error {
	return nil
}

func (r *jsiiProxy_Resource) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateResource_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResource_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateResource_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetBodyParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetCreateHeadersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetCreateQueryParametersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetDeleteHeadersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetDeleteQueryParametersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetIgnoreBodyChangesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetIgnoreCasingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetIgnoreMissingPropertyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetIgnoreNullPropertyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetIgnoreOtherItemsInListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetListUniqueIdPropertyParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetLocksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetParentIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetReadHeadersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetReadQueryParametersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetReplaceTriggersExternalValuesParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetReplaceTriggersRefsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetResponseExportValuesParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetSchemaValidationEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetSensitiveBodyParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetSensitiveBodyVersionParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetUpdateHeadersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetUpdateQueryParametersParameters(val interface{}) error {
	return nil
}

func validateNewResourceParameters(scope constructs.Construct, id *string, config *ResourceConfig) error {
	return nil
}

