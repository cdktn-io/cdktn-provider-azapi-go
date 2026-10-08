// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazapiresourceid

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzapiResourceIdConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_id#type DataAzapiResourceId#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the Azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_id#name DataAzapiResourceId#name}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_id#parent_id DataAzapiResourceId#parent_id}
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// The ID of an existing Azure source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_id#resource_id DataAzapiResourceId#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/data-sources/resource_id#timeouts DataAzapiResourceId#timeouts}
	Timeouts *DataAzapiResourceIdTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

