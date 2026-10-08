// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazapiresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzapiResourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// In a format like `<resource-type>@<api-version>`.
	//
	// `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`. `<api-version>` is version of the API used to manage this azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#type DataAzapiResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A map of headers to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#headers DataAzapiResource#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// If set to `true`, the data source will not fail when the specified resource is not found (HTTP 404).
	//
	// Identifier attributes (`id`, `name`, `parent_id`, `resource_id`) will still be populated based on inputs; other computed attributes (`output`, `location`, `identity`, `tags`) will be null/empty. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#ignore_not_found DataAzapiResource#ignore_not_found}
	IgnoreNotFound interface{} `field:"optional" json:"ignoreNotFound" yaml:"ignoreNotFound"`
	// Specifies the name of the Azure resource.
	//
	// Exactly one of the arguments `name` or `resource_id` must be set. It could be omitted if the `type` is `Microsoft.Resources/subscriptions`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#name DataAzapiResource#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the azure resource in which this resource is created.
	//
	// It supports different kinds of deployment scope for **top level** resources:
	//
	// - resource group scope: `parent_id` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
	// - management group scope: `parent_id` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
	// - extension scope: `parent_id` should be the ID of the resource you're adding the extension to.
	// - subscription scope: `parent_id` should be like \x60/subscriptions/00000000-0000-0000-0000-000000000000\x60
	// - tenant scope: `parent_id` should be /
	//
	// For child level resources, the `parent_id` should be the ID of its parent resource, for example, subnet resource's `parent_id` is the ID of the vnet.
	//
	// For type `Microsoft.Resources/resourceGroups`, the `parent_id` could be omitted, it defaults to subscription ID specified in provider or the default subscription (You could check the default subscription by azure cli command: `az account show`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#parent_id DataAzapiResource#parent_id}
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// A map of query parameters to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#query_parameters DataAzapiResource#query_parameters}
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The ID of the Azure resource to retrieve.
	//
	// Exactly one of the arguments `name` or `resource_id` must be set. It could be omitted if the `type` is `Microsoft.Resources/subscriptions`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#resource_id DataAzapiResource#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The attribute can accept either a list or a map.
	//
	// - **List**: A list of paths that need to be exported from the response body. Setting it to `["*"]` will export the full response body. Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following HCL object to the computed property output.
	//
	// 	```text
	// 	{
	// 		properties = {
	// 			loginServer = "registry1.azurecr.io"
	// 			policies = {
	// 				quarantinePolicy = {
	// 					status = "disabled"
	// 				}
	// 			}
	// 		}
	// 	}
	// 	```
	//
	// - **Map**: A map where the key is the name for the result and the value is a JMESPath query string to filter the response. Here's an example. If it sets to `{"login_server": "properties.loginServer", "quarantine_status": "properties.policies.quarantinePolicy.status"}`, it will set the following HCL object to the computed property output.
	//
	// 	```text
	// 	{
	// 		"login_server" = "registry1.azurecr.io"
	// 		"quarantine_status" = "disabled"
	// 	}
	// 	```
	//
	// To learn more about JMESPath, visit [JMESPath](https://jmespath.org/).
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#response_export_values DataAzapiResource#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#retry DataAzapiResource#retry}
	Retry *DataAzapiResourceRetry `field:"optional" json:"retry" yaml:"retry"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource#timeouts DataAzapiResource#timeouts}
	Timeouts *DataAzapiResourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

