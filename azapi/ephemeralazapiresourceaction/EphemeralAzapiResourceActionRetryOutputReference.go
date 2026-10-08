// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralazapiresourceaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/ephemeralazapiresourceaction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAzapiResourceActionRetryOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ErrorMessageRegex() *[]*string
	SetErrorMessageRegex(val *[]*string)
	ErrorMessageRegexInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IntervalSeconds() *float64
	SetIntervalSeconds(val *float64)
	IntervalSecondsInput() *float64
	MaxIntervalSeconds() *float64
	SetMaxIntervalSeconds(val *float64)
	MaxIntervalSecondsInput() *float64
	Multiplier() *float64
	SetMultiplier(val *float64)
	MultiplierInput() *float64
	RandomizationFactor() *float64
	SetRandomizationFactor(val *float64)
	RandomizationFactorInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetIntervalSeconds()
	ResetMaxIntervalSeconds()
	ResetMultiplier()
	ResetRandomizationFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EphemeralAzapiResourceActionRetryOutputReference
type jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ErrorMessageRegex() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorMessageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ErrorMessageRegexInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorMessageRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) MaxIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) MaxIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) Multiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) MultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) RandomizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) RandomizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEphemeralAzapiResourceActionRetryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EphemeralAzapiResourceActionRetryOutputReference {
	_init_.Initialize()

	if err := validateNewEphemeralAzapiResourceActionRetryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azapi.ephemeralAzapiResourceAction.EphemeralAzapiResourceActionRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEphemeralAzapiResourceActionRetryOutputReference_Override(e EphemeralAzapiResourceActionRetryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.ephemeralAzapiResourceAction.EphemeralAzapiResourceActionRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetErrorMessageRegex(val *[]*string) {
	if err := j.validateSetErrorMessageRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorMessageRegex",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetMaxIntervalSeconds(val *float64) {
	if err := j.validateSetMaxIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetMultiplier(val *float64) {
	if err := j.validateSetMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplier",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetRandomizationFactor(val *float64) {
	if err := j.validateSetRandomizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomizationFactor",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ResetIntervalSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetIntervalSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ResetMaxIntervalSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxIntervalSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ResetMultiplier() {
	_jsii_.InvokeVoid(
		e,
		"resetMultiplier",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ResetRandomizationFactor() {
	_jsii_.InvokeVoid(
		e,
		"resetRandomizationFactor",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAzapiResourceActionRetryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

