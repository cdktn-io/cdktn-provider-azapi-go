// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the azapi provider.
type AzapiProviderFunctions interface {
	// This function constructs an Azure resource ID given the parent ID, resource type, and resource name.
	//
	// It is useful for creating resource IDs for top-level and nested resources within a specific scope.
	BuildResourceId(parentId *string, resourceType *string, name *string) *string
	// This function constructs an Azure extension resource ID given the base resource ID, resource type, and additional resource names.
	ExtensionResourceId(baseResourceId *string, resourceType *string, resourceNames *[]*string) *string
	// This function constructs an Azure management group scope resource ID given the management group name, resource type, and resource names.
	ManagementGroupResourceId(managementGroupName *string, resourceType *string, resourceNames *[]*string) *string
	// This function takes an Azure resource ID and a resource type and parses the ID into its individual components such as subscription ID, resource group name, provider namespace, and other parts.
	ParseResourceId(resourceType *string, resourceId *string) cdktn.IResolvable
	// This function constructs an Azure resource group scope resource ID given the subscription ID, resource group name, resource type, and resource names.
	ResourceGroupResourceId(subscriptionId *string, resourceGroupName *string, resourceType *string, resourceNames *[]*string) *string
	// Converts all keys in the input from snake_case to camelCase.
	//
	// Retains the original structure and values.
	Snake2Camel(input interface{}) cdktn.IResolvable
	// This function constructs an Azure subscription scope resource ID given the subscription ID, resource type, and resource names.
	SubscriptionResourceId(subscriptionId *string, resourceType *string, resourceNames *[]*string) *string
	// This function constructs an Azure tenant scope resource ID given the resource type and resource names.
	TenantResourceId(resourceType *string, resourceNames *[]*string) *string
	// This function constructs an Azure equivalent `uniqueString` value.
	//
	// It is useful for migrating existing resources based on the ARM `uniqueString` function.
	UniqueString(baseString *[]*string) *string
}

// The jsii proxy struct for AzapiProviderFunctions
type jsiiProxy_AzapiProviderFunctions struct {
	_ byte // padding
}

func NewAzapiProviderFunctions(providerLocalName *string) AzapiProviderFunctions {
	_init_.Initialize()

	if err := validateNewAzapiProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzapiProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-azapi.providerFunctions.AzapiProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewAzapiProviderFunctions_Override(a AzapiProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.providerFunctions.AzapiProviderFunctions",
		[]interface{}{providerLocalName},
		a,
	)
}

func (a *jsiiProxy_AzapiProviderFunctions) BuildResourceId(parentId *string, resourceType *string, name *string) *string {
	if err := a.validateBuildResourceIdParameters(parentId, resourceType, name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"buildResourceId",
		[]interface{}{parentId, resourceType, name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) ExtensionResourceId(baseResourceId *string, resourceType *string, resourceNames *[]*string) *string {
	if err := a.validateExtensionResourceIdParameters(baseResourceId, resourceType, resourceNames); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"extensionResourceId",
		[]interface{}{baseResourceId, resourceType, resourceNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) ManagementGroupResourceId(managementGroupName *string, resourceType *string, resourceNames *[]*string) *string {
	if err := a.validateManagementGroupResourceIdParameters(managementGroupName, resourceType, resourceNames); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"managementGroupResourceId",
		[]interface{}{managementGroupName, resourceType, resourceNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) ParseResourceId(resourceType *string, resourceId *string) cdktn.IResolvable {
	if err := a.validateParseResourceIdParameters(resourceType, resourceId); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"parseResourceId",
		[]interface{}{resourceType, resourceId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) ResourceGroupResourceId(subscriptionId *string, resourceGroupName *string, resourceType *string, resourceNames *[]*string) *string {
	if err := a.validateResourceGroupResourceIdParameters(subscriptionId, resourceGroupName, resourceType, resourceNames); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"resourceGroupResourceId",
		[]interface{}{subscriptionId, resourceGroupName, resourceType, resourceNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) Snake2Camel(input interface{}) cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"snake2Camel",
		[]interface{}{input},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) SubscriptionResourceId(subscriptionId *string, resourceType *string, resourceNames *[]*string) *string {
	if err := a.validateSubscriptionResourceIdParameters(subscriptionId, resourceType, resourceNames); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"subscriptionResourceId",
		[]interface{}{subscriptionId, resourceType, resourceNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) TenantResourceId(resourceType *string, resourceNames *[]*string) *string {
	if err := a.validateTenantResourceIdParameters(resourceType, resourceNames); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"tenantResourceId",
		[]interface{}{resourceType, resourceNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzapiProviderFunctions) UniqueString(baseString *[]*string) *string {
	if err := a.validateUniqueStringParameters(baseString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"uniqueString",
		[]interface{}{baseString},
		&returns,
	)

	return returns
}

