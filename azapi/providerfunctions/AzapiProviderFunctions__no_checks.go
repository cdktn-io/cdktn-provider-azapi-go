// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzapiProviderFunctions) validateBuildResourceIdParameters(parentId *string, resourceType *string, name *string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateExtensionResourceIdParameters(baseResourceId *string, resourceType *string, resourceNames *[]*string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateManagementGroupResourceIdParameters(managementGroupName *string, resourceType *string, resourceNames *[]*string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateParseResourceIdParameters(resourceType *string, resourceId *string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateResourceGroupResourceIdParameters(subscriptionId *string, resourceGroupName *string, resourceType *string, resourceNames *[]*string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateSubscriptionResourceIdParameters(subscriptionId *string, resourceType *string, resourceNames *[]*string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateTenantResourceIdParameters(resourceType *string, resourceNames *[]*string) error {
	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateUniqueStringParameters(baseString *[]*string) error {
	return nil
}

func validateNewAzapiProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

