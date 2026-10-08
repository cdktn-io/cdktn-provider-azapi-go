// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazapiresourceaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzapiResourceActionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#type DataAzapiResourceAction#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the resource action.
	//
	// It's also possible to make HTTP requests towards the resource ID if leave this field empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#action DataAzapiResourceAction#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The body attribute is a dynamic attribute that only allows users to specify the resource body as an HCL object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#body DataAzapiResourceAction#body}
	Body *map[string]interface{} `field:"optional" json:"body" yaml:"body"`
	// A map of headers to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#headers DataAzapiResourceAction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// The HTTP method to use when performing the action. Defaults to `POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#method DataAzapiResourceAction#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// A map of query parameters to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#query_parameters DataAzapiResourceAction#query_parameters}
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The ID of the Azure resource to perform the action on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#resource_id DataAzapiResourceAction#resource_id}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#response_export_values DataAzapiResourceAction#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#retry DataAzapiResourceAction#retry}
	Retry *DataAzapiResourceActionRetry `field:"optional" json:"retry" yaml:"retry"`
	// The attribute can accept either a list or a map.
	//
	// - **List**: A list of paths that need to be exported from the response body. Setting it to `["*"]` will export the full response body. Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following HCL object to the computed property sensitive_output.
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
	// - **Map**: A map where the key is the name for the result and the value is a JMESPath query string to filter the response. Here's an example. If it sets to `{"login_server": "properties.loginServer", "quarantine_status": "properties.policies.quarantinePolicy.status"}`, it will set the following HCL object to the computed property sensitive_output.
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#sensitive_response_export_values DataAzapiResourceAction#sensitive_response_export_values}
	SensitiveResponseExportValues *map[string]interface{} `field:"optional" json:"sensitiveResponseExportValues" yaml:"sensitiveResponseExportValues"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_action#timeouts DataAzapiResourceAction#timeouts}
	Timeouts *DataAzapiResourceActionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

