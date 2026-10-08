// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package updateresource


type UpdateResourceReadOverride struct {
	// The name of the action appended to the resource ID, for example `list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/update_resource#action UpdateResource#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The HTTP method used to read the resource. The only supported value is `POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/update_resource#method UpdateResource#method}
	Method *string `field:"required" json:"method" yaml:"method"`
}

