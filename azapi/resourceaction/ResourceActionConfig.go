// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourceaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourceActionConfig struct {
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
	// The ID of an existing Azure source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#resource_id ResourceAction#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// In a format like `<resource-type>@<api-version>`.
	//
	// `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`. `<api-version>` is version of the API used to manage this azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#type ResourceAction#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the resource action.
	//
	// It's also possible to make HTTP requests towards the resource ID if leave this field empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#action ResourceAction#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// A dynamic attribute that contains the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#body ResourceAction#body}
	Body *map[string]interface{} `field:"optional" json:"body" yaml:"body"`
	// A map of headers to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#headers ResourceAction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// If set to `true`, the resource action will ignore `Not Found` errors returned from the Azure API.
	//
	// Default is `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#ignore_not_found ResourceAction#ignore_not_found}
	IgnoreNotFound interface{} `field:"optional" json:"ignoreNotFound" yaml:"ignoreNotFound"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#locks ResourceAction#locks}
	Locks *[]*string `field:"optional" json:"locks" yaml:"locks"`
	// Specifies the HTTP method of the azure resource action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#method ResourceAction#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// A map of query parameters to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#query_parameters ResourceAction#query_parameters}
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#response_export_values ResourceAction#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#retry ResourceAction#retry}
	Retry *ResourceActionRetry `field:"optional" json:"retry" yaml:"retry"`
	// A dynamic attribute that contains the write-only properties of the request body.
	//
	// This will be merge-patched to the body to construct the actual request body. If a property is defined in both `body` and `sensitive_body`, the `sensitive_body` value takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#sensitive_body ResourceAction#sensitive_body}
	SensitiveBody *map[string]interface{} `field:"optional" json:"sensitiveBody" yaml:"sensitiveBody"`
	// A map where the key is the path to the property in `sensitive_body` and the value is the version of the property.
	//
	// The key is a string in the format of `path.to.property[index].subproperty`, where `index` is the index of the item in an array. When the version is changed, the property will be included in the request body, otherwise it will be omitted from the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#sensitive_body_version ResourceAction#sensitive_body_version}
	SensitiveBodyVersion *map[string]*string `field:"optional" json:"sensitiveBodyVersion" yaml:"sensitiveBodyVersion"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#sensitive_response_export_values ResourceAction#sensitive_response_export_values}
	SensitiveResponseExportValues *map[string]interface{} `field:"optional" json:"sensitiveResponseExportValues" yaml:"sensitiveResponseExportValues"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#timeouts ResourceAction#timeouts}
	Timeouts *ResourceActionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// When to perform the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource_action#when ResourceAction#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

