// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"
)

func (a *jsiiProxy_AzapiProviderFunctions) validateBuildResourceIdParameters(parentId *string, resourceType *string, name *string) error {
	if parentId == nil {
		return fmt.Errorf("parameter parentId is required, but nil was provided")
	}

	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateExtensionResourceIdParameters(baseResourceId *string, resourceType *string, resourceNames *[]*string) error {
	if baseResourceId == nil {
		return fmt.Errorf("parameter baseResourceId is required, but nil was provided")
	}

	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if resourceNames == nil {
		return fmt.Errorf("parameter resourceNames is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateManagementGroupResourceIdParameters(managementGroupName *string, resourceType *string, resourceNames *[]*string) error {
	if managementGroupName == nil {
		return fmt.Errorf("parameter managementGroupName is required, but nil was provided")
	}

	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if resourceNames == nil {
		return fmt.Errorf("parameter resourceNames is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateParseResourceIdParameters(resourceType *string, resourceId *string) error {
	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if resourceId == nil {
		return fmt.Errorf("parameter resourceId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateResourceGroupResourceIdParameters(subscriptionId *string, resourceGroupName *string, resourceType *string, resourceNames *[]*string) error {
	if subscriptionId == nil {
		return fmt.Errorf("parameter subscriptionId is required, but nil was provided")
	}

	if resourceGroupName == nil {
		return fmt.Errorf("parameter resourceGroupName is required, but nil was provided")
	}

	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if resourceNames == nil {
		return fmt.Errorf("parameter resourceNames is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateSubscriptionResourceIdParameters(subscriptionId *string, resourceType *string, resourceNames *[]*string) error {
	if subscriptionId == nil {
		return fmt.Errorf("parameter subscriptionId is required, but nil was provided")
	}

	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if resourceNames == nil {
		return fmt.Errorf("parameter resourceNames is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateTenantResourceIdParameters(resourceType *string, resourceNames *[]*string) error {
	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	if resourceNames == nil {
		return fmt.Errorf("parameter resourceNames is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AzapiProviderFunctions) validateUniqueStringParameters(baseString *[]*string) error {
	if baseString == nil {
		return fmt.Errorf("parameter baseString is required, but nil was provided")
	}

	return nil
}

func validateNewAzapiProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

