// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralazapiresourceaction


type EphemeralAzapiResourceActionRetry struct {
	// A list of regular expressions to match against error messages.
	//
	// If any of the regular expressions match, the request will be retried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#error_message_regex EphemeralAzapiResourceAction#error_message_regex}
	ErrorMessageRegex *[]*string `field:"required" json:"errorMessageRegex" yaml:"errorMessageRegex"`
	// The base number of seconds to wait between retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#interval_seconds EphemeralAzapiResourceAction#interval_seconds}
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The maximum number of seconds to wait between retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#max_interval_seconds EphemeralAzapiResourceAction#max_interval_seconds}
	MaxIntervalSeconds *float64 `field:"optional" json:"maxIntervalSeconds" yaml:"maxIntervalSeconds"`
	// The multiplier to apply to the interval between retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#multiplier EphemeralAzapiResourceAction#multiplier}
	Multiplier *float64 `field:"optional" json:"multiplier" yaml:"multiplier"`
	// The randomization factor to apply to the interval between retries.
	//
	// The formula for the randomized interval is: `RetryInterval * (random value in range [1 - RandomizationFactor, 1 + RandomizationFactor])`. Therefore set to zero `0.0` for no randomization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/ephemeral-resources/resource_action#randomization_factor EphemeralAzapiResourceAction#randomization_factor}
	RandomizationFactor *float64 `field:"optional" json:"randomizationFactor" yaml:"randomizationFactor"`
}

