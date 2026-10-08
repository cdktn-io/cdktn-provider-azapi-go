// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazapiresourcelist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzapiResourceListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#parent_id DataAzapiResourceList#parent_id}
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// In a format like `<resource-type>@<api-version>`.
	//
	// `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`. `<api-version>` is version of the API used to manage this azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#type DataAzapiResourceList#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A map of headers to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#headers DataAzapiResourceList#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// A map of query parameters to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#query_parameters DataAzapiResourceList#query_parameters}
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The attribute can accept either a list or a map.
	//
	// - **List**: A list of paths that need to be exported from the response body. Setting it to `["*"]` will export the full response body. Here's an example. If it sets to `["value"]`, it will set the following HCL object to the computed property output.
	//
	// 	```text
	// 	{
	// 	  "value" = [
	// 		{
	// 		  "id" = "/subscriptions/000000/resourceGroups/demo-rg/providers/Microsoft.Automation/automationAccounts/example"
	// 		  "location" = "eastus2"
	// 		  "name" = "example"
	// 		  "properties" = {
	// 			"creationTime" = "2024-10-11T08:18:38.737+00:00"
	// 			"disableLocalAuth" = false
	// 			"lastModifiedTime" = "2024-10-11T08:18:38.737+00:00"
	// 			"publicNetworkAccess" = true
	// 		  }
	// 		  "tags" = {}
	// 		  "type" = "Microsoft.Automation/AutomationAccounts"
	// 		}
	// 	  ]
	// 	}
	// 	```
	//
	// - **Map**: A map where the key is the name for the result and the value is a JMESPath query string to filter the response. Here's an example. If it sets to `{"values": "value[].{name: name, publicNetworkAccess: properties.publicNetworkAccess}", "names": "value[].name"}`, it will set the following HCL object to the computed property output.
	//
	// 	```text
	// 	{
	// 		"names" = [
	// 			"example",
	// 			"fredaccount01",
	// 		]
	// 		"values" = [
	// 			{
	// 			  "name" = "example"
	// 			  "publicNetworkAccess" = true
	// 			},
	// 			{
	// 			  "name" = "fredaccount01"
	// 			  "publicNetworkAccess" = null
	// 			},
	// 		]
	// 	}
	// 	```
	//
	// To learn more about JMESPath, visit [JMESPath](https://jmespath.org/).
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#response_export_values DataAzapiResourceList#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#retry DataAzapiResourceList#retry}
	Retry *DataAzapiResourceListRetry `field:"optional" json:"retry" yaml:"retry"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_list#timeouts DataAzapiResourceList#timeouts}
	Timeouts *DataAzapiResourceListTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

