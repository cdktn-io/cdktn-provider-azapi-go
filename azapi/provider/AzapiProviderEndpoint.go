// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AzapiProviderEndpoint struct {
	// The Azure Active Directory login endpoint to use.
	//
	// This can also be sourced from the `ARM_ACTIVE_DIRECTORY_AUTHORITY_HOST` Environment Variable. Defaults to `https://login.microsoftonline.com/` for public cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#active_directory_authority_host AzapiProvider#active_directory_authority_host}
	ActiveDirectoryAuthorityHost *string `field:"optional" json:"activeDirectoryAuthorityHost" yaml:"activeDirectoryAuthorityHost"`
	// The resource ID to obtain AD tokens for.
	//
	// This can also be sourced from the `ARM_RESOURCE_MANAGER_AUDIENCE` Environment Variable. Defaults to `https://management.core.windows.net/` for public cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#resource_manager_audience AzapiProvider#resource_manager_audience}
	ResourceManagerAudience *string `field:"optional" json:"resourceManagerAudience" yaml:"resourceManagerAudience"`
	// The Azure Resource Manager endpoint to use.
	//
	// This can also be sourced from the `ARM_RESOURCE_MANAGER_ENDPOINT` Environment Variable. Defaults to `https://management.azure.com/` for public cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs#resource_manager_endpoint AzapiProvider#resource_manager_endpoint}
	ResourceManagerEndpoint *string `field:"optional" json:"resourceManagerEndpoint" yaml:"resourceManagerEndpoint"`
}

