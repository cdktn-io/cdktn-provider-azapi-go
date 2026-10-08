// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazapiresource

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzapiResourceIdentityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzapiResourceIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzapiResourceIdentityList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzapiResourceIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzapiResourceIdentityList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzapiResourceIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzapiResourceIdentityListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

