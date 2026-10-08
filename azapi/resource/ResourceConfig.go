// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#type Resource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A dynamic attribute that contains the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#body Resource#body}
	Body *map[string]interface{} `field:"optional" json:"body" yaml:"body"`
	// A mapping of headers to be sent with the create request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#create_headers Resource#create_headers}
	CreateHeaders *map[string]*string `field:"optional" json:"createHeaders" yaml:"createHeaders"`
	// A mapping of query parameters to be sent with the create request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#create_query_parameters Resource#create_query_parameters}
	CreateQueryParameters interface{} `field:"optional" json:"createQueryParameters" yaml:"createQueryParameters"`
	// A mapping of headers to be sent with the delete request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#delete_headers Resource#delete_headers}
	DeleteHeaders *map[string]*string `field:"optional" json:"deleteHeaders" yaml:"deleteHeaders"`
	// A mapping of query parameters to be sent with the delete request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#delete_query_parameters Resource#delete_query_parameters}
	DeleteQueryParameters interface{} `field:"optional" json:"deleteQueryParameters" yaml:"deleteQueryParameters"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#identity Resource#identity}
	Identity interface{} `field:"optional" json:"identity" yaml:"identity"`
	// A list of paths in the resource body whose changes should be ignored.
	//
	// Prefer Terraform's `lifecycle.ignore_changes` when possible. Use this argument only when the paths must be derived from variables or other non-static values. Changes to this argument take effect only after an apply because its value is stored in provider-private state. Paths use dot notation, for example `properties.sku.name`. Individual list items cannot be targeted, ignore the entire list property instead. Configuration changes at an ignored path will not be sent to Azure until that path is removed from this list. This write-only argument requires Terraform 1.11 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#ignore_body_changes Resource#ignore_body_changes}
	IgnoreBodyChanges *[]*string `field:"optional" json:"ignoreBodyChanges" yaml:"ignoreBodyChanges"`
	// Whether ignore the casing of the property names in the response body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#ignore_casing Resource#ignore_casing}
	IgnoreCasing interface{} `field:"optional" json:"ignoreCasing" yaml:"ignoreCasing"`
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff.
	//
	// It's recommend to enable this option when some sensitive properties are not returned in response body, instead of setting them in `lifecycle.ignore_changes` because it will make the sensitive fields unable to update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#ignore_missing_property Resource#ignore_missing_property}
	IgnoreMissingProperty interface{} `field:"optional" json:"ignoreMissingProperty" yaml:"ignoreMissingProperty"`
	// When set to `true`, the provider will ignore properties whose values are `null` in the `body`.
	//
	// These properties will not be included in the request body sent to the API, and the difference will not be shown in the plan output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#ignore_null_property Resource#ignore_null_property}
	IgnoreNullProperty interface{} `field:"optional" json:"ignoreNullProperty" yaml:"ignoreNullProperty"`
	// A list of list property paths where items not specified in configuration should be ignored.
	//
	// This is intended for partial list management when combined with `list_unique_id_property` (for example, to avoid perpetual drift from server-side ordering).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#ignore_other_items_in_list Resource#ignore_other_items_in_list}
	IgnoreOtherItemsInList *[]*string `field:"optional" json:"ignoreOtherItemsInList" yaml:"ignoreOtherItemsInList"`
	// A mapping of list property paths to the field name used as a unique identifier when comparing and merging list items.
	//
	// When not set, list items are matched by a `name` property (if present) or by list ordering. To match using multiple fields, specify a comma-separated list of field names (e.g., `"category, categoryGroup"`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#list_unique_id_property Resource#list_unique_id_property}
	ListUniqueIdProperty *map[string]*string `field:"optional" json:"listUniqueIdProperty" yaml:"listUniqueIdProperty"`
	// The location of the Azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#location Resource#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#locks Resource#locks}
	Locks *[]*string `field:"optional" json:"locks" yaml:"locks"`
	// Specifies the name of the azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#name Resource#name}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#parent_id Resource#parent_id}
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// A mapping of headers to be sent with the read request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#read_headers Resource#read_headers}
	ReadHeaders *map[string]*string `field:"optional" json:"readHeaders" yaml:"readHeaders"`
	// A mapping of query parameters to be sent with the read request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#read_query_parameters Resource#read_query_parameters}
	ReadQueryParameters interface{} `field:"optional" json:"readQueryParameters" yaml:"readQueryParameters"`
	// Will trigger a replace of the resource when the value changes and is not `null`.
	//
	// This can be used by practitioners to force a replace of the resource when certain values change, e.g. changing the SKU of a virtual machine based on the value of variables or locals. The value is a `dynamic`, so practitioners can compose the input however they wish. For a "break glass" set the value to `null` to prevent the plan modifier taking effect.
	// If you have `null` values that you do want to be tracked as affecting the resource replacement, include these inside an object.
	// Advanced use cases are possible and resource replacement can be triggered by values external to the resource, for example when a dependent resource changes.
	//
	// e.g. to replace a resource when either the SKU or os_type attributes change:
	//
	// ```hcl
	// resource "azapi_resource" "example" {
	//   name      = var.name
	//   type      = "Microsoft.Network/publicIPAddresses@2023-11-01"
	//   parent_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example"
	//   body = {
	//     properties = {
	//       sku   = var.sku
	//       zones = var.zones
	//     }
	//   }
	//
	//   replace_triggers_external_values = [
	//     var.sku,
	//     var.zones,
	//   ]
	// }
	// ```
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#replace_triggers_external_values Resource#replace_triggers_external_values}
	ReplaceTriggersExternalValues *map[string]interface{} `field:"optional" json:"replaceTriggersExternalValues" yaml:"replaceTriggersExternalValues"`
	// A list of paths in the current Terraform configuration.
	//
	// When the values at these paths change, the resource will be replaced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#replace_triggers_refs Resource#replace_triggers_refs}
	ReplaceTriggersRefs *[]*string `field:"optional" json:"replaceTriggersRefs" yaml:"replaceTriggersRefs"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#response_export_values Resource#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#retry Resource#retry}
	Retry *ResourceRetry `field:"optional" json:"retry" yaml:"retry"`
	// Whether enabled the validation on `type` and `body` with embedded schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#schema_validation_enabled Resource#schema_validation_enabled}
	SchemaValidationEnabled interface{} `field:"optional" json:"schemaValidationEnabled" yaml:"schemaValidationEnabled"`
	// A dynamic attribute that contains the write-only properties of the request body.
	//
	// This will be merge-patched to the body to construct the actual request body. If a property is defined in both `body` and `sensitive_body`, the `sensitive_body` value takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#sensitive_body Resource#sensitive_body}
	SensitiveBody *map[string]interface{} `field:"optional" json:"sensitiveBody" yaml:"sensitiveBody"`
	// A map where the key is the path to the property in `sensitive_body` and the value is the version of the property.
	//
	// The key is a string in the format of `path.to.property[index].subproperty`, where `index` is the index of the item in an array. When the version is changed, the property will be included in the request body, otherwise it will be omitted from the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#sensitive_body_version Resource#sensitive_body_version}
	SensitiveBodyVersion *map[string]*string `field:"optional" json:"sensitiveBodyVersion" yaml:"sensitiveBodyVersion"`
	// A mapping of tags which should be assigned to the Azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#tags Resource#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#timeouts Resource#timeouts}
	Timeouts *ResourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A mapping of headers to be sent with the update request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#update_headers Resource#update_headers}
	UpdateHeaders *map[string]*string `field:"optional" json:"updateHeaders" yaml:"updateHeaders"`
	// A mapping of query parameters to be sent with the update request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#update_query_parameters Resource#update_query_parameters}
	UpdateQueryParameters interface{} `field:"optional" json:"updateQueryParameters" yaml:"updateQueryParameters"`
}

