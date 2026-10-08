// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazapiresourcelist

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/dataazapiresourcelist/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzapiResourceListRetryOutputReference interface {
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

// The jsii proxy struct for DataAzapiResourceListRetryOutputReference
type jsiiProxy_DataAzapiResourceListRetryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) ErrorMessageRegex() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorMessageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) ErrorMessageRegexInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorMessageRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) MaxIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) MaxIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) Multiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) MultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) RandomizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) RandomizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAzapiResourceListRetryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAzapiResourceListRetryOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzapiResourceListRetryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzapiResourceListRetryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azapi.dataAzapiResourceList.DataAzapiResourceListRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAzapiResourceListRetryOutputReference_Override(d DataAzapiResourceListRetryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.dataAzapiResourceList.DataAzapiResourceListRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetErrorMessageRegex(val *[]*string) {
	if err := j.validateSetErrorMessageRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorMessageRegex",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetMaxIntervalSeconds(val *float64) {
	if err := j.validateSetMaxIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetMultiplier(val *float64) {
	if err := j.validateSetMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplier",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetRandomizationFactor(val *float64) {
	if err := j.validateSetRandomizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomizationFactor",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzapiResourceListRetryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) ResetIntervalSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetIntervalSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) ResetMaxIntervalSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIntervalSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) ResetMultiplier() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiplier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) ResetRandomizationFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetRandomizationFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzapiResourceListRetryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

