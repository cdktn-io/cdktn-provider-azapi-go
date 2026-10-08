// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package updateresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/updateresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UpdateResourceRetryOutputReference interface {
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

// The jsii proxy struct for UpdateResourceRetryOutputReference
type jsiiProxy_UpdateResourceRetryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) ErrorMessageRegex() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorMessageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) ErrorMessageRegexInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorMessageRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) MaxIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) MaxIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) Multiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) MultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) RandomizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) RandomizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewUpdateResourceRetryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) UpdateResourceRetryOutputReference {
	_init_.Initialize()

	if err := validateNewUpdateResourceRetryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_UpdateResourceRetryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azapi.updateResource.UpdateResourceRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewUpdateResourceRetryOutputReference_Override(u UpdateResourceRetryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.updateResource.UpdateResourceRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		u,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetErrorMessageRegex(val *[]*string) {
	if err := j.validateSetErrorMessageRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorMessageRegex",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetMaxIntervalSeconds(val *float64) {
	if err := j.validateSetMaxIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetMultiplier(val *float64) {
	if err := j.validateSetMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplier",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetRandomizationFactor(val *float64) {
	if err := j.validateSetRandomizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomizationFactor",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_UpdateResourceRetryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) ResetIntervalSeconds() {
	_jsii_.InvokeVoid(
		u,
		"resetIntervalSeconds",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) ResetMaxIntervalSeconds() {
	_jsii_.InvokeVoid(
		u,
		"resetMaxIntervalSeconds",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) ResetMultiplier() {
	_jsii_.InvokeVoid(
		u,
		"resetMultiplier",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) ResetRandomizationFactor() {
	_jsii_.InvokeVoid(
		u,
		"resetRandomizationFactor",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := u.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		u,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResourceRetryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

