// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resource


type ResourceIdentity struct {
	// The Type of Identity which should be used for this azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#type Resource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of User Managed Identity ID's which should be assigned to the azure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource#identity_ids Resource#identity_ids}
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

