// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralazapiresourceaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAzapiResourceActionConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The ID of an existing Azure source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#resource_id EphemeralAzapiResourceAction#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// In a format like `<resource-type>@<api-version>`.
	//
	// `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`. `<api-version>` is version of the API used to manage this azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#type EphemeralAzapiResourceAction#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the resource action.
	//
	// It's also possible to make HTTP requests towards the resource ID if leave this field empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#action EphemeralAzapiResourceAction#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// A dynamic attribute that contains the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#body EphemeralAzapiResourceAction#body}
	Body *map[string]interface{} `field:"optional" json:"body" yaml:"body"`
	// A map of headers to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#headers EphemeralAzapiResourceAction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#locks EphemeralAzapiResourceAction#locks}
	Locks *[]*string `field:"optional" json:"locks" yaml:"locks"`
	// Specifies the HTTP method of the azure resource action. Defaults to `POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#method EphemeralAzapiResourceAction#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// A map of query parameters to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#query_parameters EphemeralAzapiResourceAction#query_parameters}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#response_export_values EphemeralAzapiResourceAction#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#retry EphemeralAzapiResourceAction#retry}
	Retry *EphemeralAzapiResourceActionRetry `field:"optional" json:"retry" yaml:"retry"`
	// A dynamic attribute that contains the write-only properties of the request body.
	//
	// This will be merge-patched to the body to construct the actual request body. If a property is defined in both `body` and `sensitive_body`, the `sensitive_body` value takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#sensitive_body EphemeralAzapiResourceAction#sensitive_body}
	SensitiveBody *map[string]interface{} `field:"optional" json:"sensitiveBody" yaml:"sensitiveBody"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#timeouts EphemeralAzapiResourceAction#timeouts}
	Timeouts *EphemeralAzapiResourceActionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

