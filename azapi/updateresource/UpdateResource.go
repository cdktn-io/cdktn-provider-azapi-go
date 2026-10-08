// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package updateresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/updateresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/update_resource azapi_update_resource}.
type UpdateResource interface {
	cdktn.TerraformResource
	Body() *map[string]interface{}
	SetBody(val *map[string]interface{})
	BodyInput() *map[string]interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IgnoreCasing() interface{}
	SetIgnoreCasing(val interface{})
	IgnoreCasingInput() interface{}
	IgnoreMissingProperty() interface{}
	SetIgnoreMissingProperty(val interface{})
	IgnoreMissingPropertyInput() interface{}
	IgnoreOtherItemsInList() *[]*string
	SetIgnoreOtherItemsInList(val *[]*string)
	IgnoreOtherItemsInListInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListUniqueIdProperty() *map[string]*string
	SetListUniqueIdProperty(val *map[string]*string)
	ListUniqueIdPropertyInput() *map[string]*string
	Locks() *[]*string
	SetLocks(val *[]*string)
	LocksInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Output() cdktn.AnyMap
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReadHeaders() *map[string]*string
	SetReadHeaders(val *map[string]*string)
	ReadHeadersInput() *map[string]*string
	ReadOverride() UpdateResourceReadOverrideOutputReference
	ReadOverrideInput() interface{}
	ReadQueryParameters() interface{}
	SetReadQueryParameters(val interface{})
	ReadQueryParametersInput() interface{}
	ReplaceTriggersExternalValues() *map[string]interface{}
	SetReplaceTriggersExternalValues(val *map[string]interface{})
	ReplaceTriggersExternalValuesInput() *map[string]interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResponseExportValues() *map[string]interface{}
	SetResponseExportValues(val *map[string]interface{})
	ResponseExportValuesInput() *map[string]interface{}
	Retry() UpdateResourceRetryOutputReference
	RetryInput() interface{}
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SensitiveBody() *map[string]interface{}
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetSensitiveBody(val *map[string]interface{})
	SensitiveBodyInput() *map[string]interface{}
	SensitiveBodyVersion() *map[string]*string
	SetSensitiveBodyVersion(val *map[string]*string)
	SensitiveBodyVersionInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() UpdateResourceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateHeaders() *map[string]*string
	SetUpdateHeaders(val *map[string]*string)
	UpdateHeadersInput() *map[string]*string
	UpdateQueryParameters() interface{}
	SetUpdateQueryParameters(val interface{})
	UpdateQueryParametersInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutReadOverride(value *UpdateResourceReadOverride)
	PutRetry(value *UpdateResourceRetry)
	PutTimeouts(value *UpdateResourceTimeouts)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetBody()
	ResetIgnoreCasing()
	ResetIgnoreMissingProperty()
	ResetIgnoreOtherItemsInList()
	ResetListUniqueIdProperty()
	ResetLocks()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentId()
	ResetReadHeaders()
	ResetReadOverride()
	ResetReadQueryParameters()
	ResetReplaceTriggersExternalValues()
	ResetResourceId()
	ResetResponseExportValues()
	ResetRetry()
	ResetSensitiveBody()
	ResetSensitiveBodyVersion()
	ResetTimeouts()
	ResetUpdateHeaders()
	ResetUpdateQueryParameters()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for UpdateResource
type jsiiProxy_UpdateResource struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_UpdateResource) Body() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) BodyInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) IgnoreCasing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCasing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) IgnoreCasingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCasingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) IgnoreMissingProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) IgnoreMissingPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) IgnoreOtherItemsInList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreOtherItemsInList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) IgnoreOtherItemsInListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreOtherItemsInListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ListUniqueIdProperty() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listUniqueIdProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ListUniqueIdPropertyInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listUniqueIdPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Locks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) LocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Output() cdktn.AnyMap {
	var returns cdktn.AnyMap
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReadHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"readHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReadHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"readHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReadOverride() UpdateResourceReadOverrideOutputReference {
	var returns UpdateResourceReadOverrideOutputReference
	_jsii_.Get(
		j,
		"readOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReadOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReadQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReadQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReplaceTriggersExternalValues() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"replaceTriggersExternalValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ReplaceTriggersExternalValuesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"replaceTriggersExternalValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ResponseExportValues() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseExportValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) ResponseExportValuesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseExportValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Retry() UpdateResourceRetryOutputReference {
	var returns UpdateResourceRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) RetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) SensitiveBody() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"sensitiveBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) SensitiveBodyInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"sensitiveBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) SensitiveBodyVersion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sensitiveBodyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) SensitiveBodyVersionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sensitiveBodyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Timeouts() UpdateResourceTimeoutsOutputReference {
	var returns UpdateResourceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) UpdateHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updateHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) UpdateHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updateHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) UpdateQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdateResource) UpdateQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateQueryParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/update_resource azapi_update_resource} Resource.
func NewUpdateResource(scope constructs.Construct, id *string, config *UpdateResourceConfig) UpdateResource {
	_init_.Initialize()

	if err := validateNewUpdateResourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_UpdateResource{}

	_jsii_.Create(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/update_resource azapi_update_resource} Resource.
func NewUpdateResource_Override(u UpdateResource, scope constructs.Construct, id *string, config *UpdateResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_UpdateResource)SetBody(val *map[string]interface{}) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetIgnoreCasing(val interface{}) {
	if err := j.validateSetIgnoreCasingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCasing",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetIgnoreMissingProperty(val interface{}) {
	if err := j.validateSetIgnoreMissingPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMissingProperty",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetIgnoreOtherItemsInList(val *[]*string) {
	if err := j.validateSetIgnoreOtherItemsInListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOtherItemsInList",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetListUniqueIdProperty(val *map[string]*string) {
	if err := j.validateSetListUniqueIdPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listUniqueIdProperty",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetLocks(val *[]*string) {
	if err := j.validateSetLocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locks",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetParentId(val *string) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetReadHeaders(val *map[string]*string) {
	if err := j.validateSetReadHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readHeaders",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetReadQueryParameters(val interface{}) {
	if err := j.validateSetReadQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readQueryParameters",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetReplaceTriggersExternalValues(val *map[string]interface{}) {
	if err := j.validateSetReplaceTriggersExternalValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTriggersExternalValues",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetResponseExportValues(val *map[string]interface{}) {
	if err := j.validateSetResponseExportValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseExportValues",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetSensitiveBody(val *map[string]interface{}) {
	if err := j.validateSetSensitiveBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveBody",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetSensitiveBodyVersion(val *map[string]*string) {
	if err := j.validateSetSensitiveBodyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveBodyVersion",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetUpdateHeaders(val *map[string]*string) {
	if err := j.validateSetUpdateHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateHeaders",
		val,
	)
}

func (j *jsiiProxy_UpdateResource)SetUpdateQueryParameters(val interface{}) {
	if err := j.validateSetUpdateQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateQueryParameters",
		val,
	)
}

// Generates CDKTN code for importing a UpdateResource resource upon running "cdktn plan <stack-name>".
func UpdateResource_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateUpdateResource_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func UpdateResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpdateResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UpdateResource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpdateResource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UpdateResource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpdateResource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func UpdateResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azapi.updateResource.UpdateResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UpdateResource) AddMoveTarget(moveTarget *string) {
	if err := u.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (u *jsiiProxy_UpdateResource) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_UpdateResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_UpdateResource) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (u *jsiiProxy_UpdateResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_UpdateResource) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_UpdateResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_UpdateResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_UpdateResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_UpdateResource) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_UpdateResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_UpdateResource) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := u.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (u *jsiiProxy_UpdateResource) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (u *jsiiProxy_UpdateResource) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := u.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		u,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) MoveFromId(id *string) {
	if err := u.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveFromId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_UpdateResource) MoveTo(moveTarget *string, index interface{}) {
	if err := u.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (u *jsiiProxy_UpdateResource) MoveToId(id *string) {
	if err := u.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveToId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_UpdateResource) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_UpdateResource) PutReadOverride(value *UpdateResourceReadOverride) {
	if err := u.validatePutReadOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putReadOverride",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UpdateResource) PutRetry(value *UpdateResourceRetry) {
	if err := u.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putRetry",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UpdateResource) PutTimeouts(value *UpdateResourceTimeouts) {
	if err := u.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UpdateResource) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := u.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (u *jsiiProxy_UpdateResource) ResetBody() {
	_jsii_.InvokeVoid(
		u,
		"resetBody",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetIgnoreCasing() {
	_jsii_.InvokeVoid(
		u,
		"resetIgnoreCasing",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetIgnoreMissingProperty() {
	_jsii_.InvokeVoid(
		u,
		"resetIgnoreMissingProperty",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetIgnoreOtherItemsInList() {
	_jsii_.InvokeVoid(
		u,
		"resetIgnoreOtherItemsInList",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetListUniqueIdProperty() {
	_jsii_.InvokeVoid(
		u,
		"resetListUniqueIdProperty",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetLocks() {
	_jsii_.InvokeVoid(
		u,
		"resetLocks",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetName() {
	_jsii_.InvokeVoid(
		u,
		"resetName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetParentId() {
	_jsii_.InvokeVoid(
		u,
		"resetParentId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetReadHeaders() {
	_jsii_.InvokeVoid(
		u,
		"resetReadHeaders",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetReadOverride() {
	_jsii_.InvokeVoid(
		u,
		"resetReadOverride",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetReadQueryParameters() {
	_jsii_.InvokeVoid(
		u,
		"resetReadQueryParameters",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetReplaceTriggersExternalValues() {
	_jsii_.InvokeVoid(
		u,
		"resetReplaceTriggersExternalValues",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetResourceId() {
	_jsii_.InvokeVoid(
		u,
		"resetResourceId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetResponseExportValues() {
	_jsii_.InvokeVoid(
		u,
		"resetResponseExportValues",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetRetry() {
	_jsii_.InvokeVoid(
		u,
		"resetRetry",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetSensitiveBody() {
	_jsii_.InvokeVoid(
		u,
		"resetSensitiveBody",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetSensitiveBodyVersion() {
	_jsii_.InvokeVoid(
		u,
		"resetSensitiveBodyVersion",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetTimeouts() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetUpdateHeaders() {
	_jsii_.InvokeVoid(
		u,
		"resetUpdateHeaders",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) ResetUpdateQueryParameters() {
	_jsii_.InvokeVoid(
		u,
		"resetUpdateQueryParameters",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdateResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdateResource) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		u,
		"with",
		args,
		&returns,
	)

	return returns
}

