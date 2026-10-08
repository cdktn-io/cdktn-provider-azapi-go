// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplaneresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataPlaneResourceConfig struct {
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
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#parent_id DataPlaneResource#parent_id}
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// In a format like `<resource-type>@<api-version>`.
	//
	// `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`. `<api-version>` is version of the API used to manage this azure resource. For a list of supported data plane resource types, see the [Available Resources](https://registry.terraform.io/providers/Azure/azapi/latest/docs/resources/data_plane_resource#available-resources) documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#type DataPlaneResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A dynamic attribute that contains the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#body DataPlaneResource#body}
	Body *map[string]interface{} `field:"optional" json:"body" yaml:"body"`
	// A mapping of headers to be sent with the create request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#create_headers DataPlaneResource#create_headers}
	CreateHeaders *map[string]*string `field:"optional" json:"createHeaders" yaml:"createHeaders"`
	// A mapping of query parameters to be sent with the create request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#create_query_parameters DataPlaneResource#create_query_parameters}
	CreateQueryParameters interface{} `field:"optional" json:"createQueryParameters" yaml:"createQueryParameters"`
	// A mapping of headers to be sent with the delete request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#delete_headers DataPlaneResource#delete_headers}
	DeleteHeaders *map[string]*string `field:"optional" json:"deleteHeaders" yaml:"deleteHeaders"`
	// A mapping of query parameters to be sent with the delete request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#delete_query_parameters DataPlaneResource#delete_query_parameters}
	DeleteQueryParameters interface{} `field:"optional" json:"deleteQueryParameters" yaml:"deleteQueryParameters"`
	// A dynamic attribute that contains the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#ignore_casing DataPlaneResource#ignore_casing}
	IgnoreCasing interface{} `field:"optional" json:"ignoreCasing" yaml:"ignoreCasing"`
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff.
	//
	// It's recommend to enable this option when some sensitive properties are not returned in response body, instead of setting them in `lifecycle.ignore_changes` because it will make the sensitive fields unable to update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#ignore_missing_property DataPlaneResource#ignore_missing_property}
	IgnoreMissingProperty interface{} `field:"optional" json:"ignoreMissingProperty" yaml:"ignoreMissingProperty"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#locks DataPlaneResource#locks}
	Locks *[]*string `field:"optional" json:"locks" yaml:"locks"`
	// Specifies the name (identifier segment) of the data plane resource. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#name DataPlaneResource#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A mapping of headers to be sent with the read request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#read_headers DataPlaneResource#read_headers}
	ReadHeaders *map[string]*string `field:"optional" json:"readHeaders" yaml:"readHeaders"`
	// A mapping of query parameters to be sent with the read request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#read_query_parameters DataPlaneResource#read_query_parameters}
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
	// resource "azapi_data_plane_resource" "example" {
	//   name = var.name
	//   type = "Microsoft.AppConfiguration/configurationStores/keyValues@1.0"
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#replace_triggers_external_values DataPlaneResource#replace_triggers_external_values}
	ReplaceTriggersExternalValues *map[string]interface{} `field:"optional" json:"replaceTriggersExternalValues" yaml:"replaceTriggersExternalValues"`
	// A list of paths in the current Terraform configuration.
	//
	// When the values at these paths change, the resource will be replaced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#replace_triggers_refs DataPlaneResource#replace_triggers_refs}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#response_export_values DataPlaneResource#response_export_values}
	ResponseExportValues *map[string]interface{} `field:"optional" json:"responseExportValues" yaml:"responseExportValues"`
	// The retry object supports the following attributes:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#retry DataPlaneResource#retry}
	Retry *DataPlaneResourceRetry `field:"optional" json:"retry" yaml:"retry"`
	// A dynamic attribute that contains the write-only properties of the request body.
	//
	// This will be merge-patched to the body to construct the actual request body. If a property is defined in both `body` and `sensitive_body`, the `sensitive_body` value takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#sensitive_body DataPlaneResource#sensitive_body}
	SensitiveBody *map[string]interface{} `field:"optional" json:"sensitiveBody" yaml:"sensitiveBody"`
	// A map where the key is the path to the property in `sensitive_body` and the value is the version of the property.
	//
	// The key is a string in the format of `path.to.property[index].subproperty`, where `index` is the index of the item in an array. When the version is changed, the property will be included in the request body, otherwise it will be omitted from the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#sensitive_body_version DataPlaneResource#sensitive_body_version}
	SensitiveBodyVersion *map[string]*string `field:"optional" json:"sensitiveBodyVersion" yaml:"sensitiveBodyVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#timeouts DataPlaneResource#timeouts}
	Timeouts *DataPlaneResourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A mapping of headers to be sent with the update request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#update_headers DataPlaneResource#update_headers}
	UpdateHeaders *map[string]*string `field:"optional" json:"updateHeaders" yaml:"updateHeaders"`
	// A mapping of query parameters to be sent with the update request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource#update_query_parameters DataPlaneResource#update_query_parameters}
	UpdateQueryParameters interface{} `field:"optional" json:"updateQueryParameters" yaml:"updateQueryParameters"`
}

