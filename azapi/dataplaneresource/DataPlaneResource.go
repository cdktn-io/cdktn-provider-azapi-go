// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplaneresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/dataplaneresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource azapi_data_plane_resource}.
type DataPlaneResource interface {
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
	CreateHeaders() *map[string]*string
	SetCreateHeaders(val *map[string]*string)
	CreateHeadersInput() *map[string]*string
	CreateQueryParameters() interface{}
	SetCreateQueryParameters(val interface{})
	CreateQueryParametersInput() interface{}
	DeleteHeaders() *map[string]*string
	SetDeleteHeaders(val *map[string]*string)
	DeleteHeadersInput() *map[string]*string
	DeleteQueryParameters() interface{}
	SetDeleteQueryParameters(val interface{})
	DeleteQueryParametersInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	ReadQueryParameters() interface{}
	SetReadQueryParameters(val interface{})
	ReadQueryParametersInput() interface{}
	ReplaceTriggersExternalValues() *map[string]interface{}
	SetReplaceTriggersExternalValues(val *map[string]interface{})
	ReplaceTriggersExternalValuesInput() *map[string]interface{}
	ReplaceTriggersRefs() *[]*string
	SetReplaceTriggersRefs(val *[]*string)
	ReplaceTriggersRefsInput() *[]*string
	ResponseExportValues() *map[string]interface{}
	SetResponseExportValues(val *map[string]interface{})
	ResponseExportValuesInput() *map[string]interface{}
	Retry() DataPlaneResourceRetryOutputReference
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
	Timeouts() DataPlaneResourceTimeoutsOutputReference
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
	PutRetry(value *DataPlaneResourceRetry)
	PutTimeouts(value *DataPlaneResourceTimeouts)
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
	ResetCreateHeaders()
	ResetCreateQueryParameters()
	ResetDeleteHeaders()
	ResetDeleteQueryParameters()
	ResetIgnoreCasing()
	ResetIgnoreMissingProperty()
	ResetLocks()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadHeaders()
	ResetReadQueryParameters()
	ResetReplaceTriggersExternalValues()
	ResetReplaceTriggersRefs()
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

// The jsii proxy struct for DataPlaneResource
type jsiiProxy_DataPlaneResource struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DataPlaneResource) Body() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) BodyInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) CreateHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"createHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) CreateHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"createHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) CreateQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) CreateQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) DeleteHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deleteHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) DeleteHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deleteHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) DeleteQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) DeleteQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) IgnoreCasing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCasing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) IgnoreCasingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCasingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) IgnoreMissingProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) IgnoreMissingPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Locks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) LocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Output() cdktn.AnyMap {
	var returns cdktn.AnyMap
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReadHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"readHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReadHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"readHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReadQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReadQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReplaceTriggersExternalValues() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"replaceTriggersExternalValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReplaceTriggersExternalValuesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"replaceTriggersExternalValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReplaceTriggersRefs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replaceTriggersRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ReplaceTriggersRefsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replaceTriggersRefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ResponseExportValues() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseExportValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) ResponseExportValuesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseExportValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Retry() DataPlaneResourceRetryOutputReference {
	var returns DataPlaneResourceRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) RetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) SensitiveBody() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"sensitiveBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) SensitiveBodyInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"sensitiveBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) SensitiveBodyVersion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sensitiveBodyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) SensitiveBodyVersionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sensitiveBodyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Timeouts() DataPlaneResourceTimeoutsOutputReference {
	var returns DataPlaneResourceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) UpdateHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updateHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) UpdateHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updateHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) UpdateQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPlaneResource) UpdateQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateQueryParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource azapi_data_plane_resource} Resource.
func NewDataPlaneResource(scope constructs.Construct, id *string, config *DataPlaneResourceConfig) DataPlaneResource {
	_init_.Initialize()

	if err := validateNewDataPlaneResourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPlaneResource{}

	_jsii_.Create(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/data_plane_resource azapi_data_plane_resource} Resource.
func NewDataPlaneResource_Override(d DataPlaneResource, scope constructs.Construct, id *string, config *DataPlaneResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetBody(val *map[string]interface{}) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetCreateHeaders(val *map[string]*string) {
	if err := j.validateSetCreateHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createHeaders",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetCreateQueryParameters(val interface{}) {
	if err := j.validateSetCreateQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createQueryParameters",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetDeleteHeaders(val *map[string]*string) {
	if err := j.validateSetDeleteHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteHeaders",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetDeleteQueryParameters(val interface{}) {
	if err := j.validateSetDeleteQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteQueryParameters",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetIgnoreCasing(val interface{}) {
	if err := j.validateSetIgnoreCasingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCasing",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetIgnoreMissingProperty(val interface{}) {
	if err := j.validateSetIgnoreMissingPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMissingProperty",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetLocks(val *[]*string) {
	if err := j.validateSetLocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locks",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetParentId(val *string) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetReadHeaders(val *map[string]*string) {
	if err := j.validateSetReadHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readHeaders",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetReadQueryParameters(val interface{}) {
	if err := j.validateSetReadQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readQueryParameters",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetReplaceTriggersExternalValues(val *map[string]interface{}) {
	if err := j.validateSetReplaceTriggersExternalValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTriggersExternalValues",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetReplaceTriggersRefs(val *[]*string) {
	if err := j.validateSetReplaceTriggersRefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTriggersRefs",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetResponseExportValues(val *map[string]interface{}) {
	if err := j.validateSetResponseExportValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseExportValues",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetSensitiveBody(val *map[string]interface{}) {
	if err := j.validateSetSensitiveBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveBody",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetSensitiveBodyVersion(val *map[string]*string) {
	if err := j.validateSetSensitiveBodyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveBodyVersion",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetUpdateHeaders(val *map[string]*string) {
	if err := j.validateSetUpdateHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateHeaders",
		val,
	)
}

func (j *jsiiProxy_DataPlaneResource)SetUpdateQueryParameters(val interface{}) {
	if err := j.validateSetUpdateQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateQueryParameters",
		val,
	)
}

// Generates CDKTN code for importing a DataPlaneResource resource upon running "cdktn plan <stack-name>".
func DataPlaneResource_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataPlaneResource_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
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
func DataPlaneResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPlaneResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPlaneResource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPlaneResource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPlaneResource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPlaneResource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataPlaneResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azapi.dataPlaneResource.DataPlaneResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataPlaneResource) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataPlaneResource) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataPlaneResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataPlaneResource) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataPlaneResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataPlaneResource) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataPlaneResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataPlaneResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataPlaneResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataPlaneResource) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataPlaneResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataPlaneResource) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataPlaneResource) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataPlaneResource) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataPlaneResource) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataPlaneResource) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataPlaneResource) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataPlaneResource) PutRetry(value *DataPlaneResourceRetry) {
	if err := d.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetry",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataPlaneResource) PutTimeouts(value *DataPlaneResourceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataPlaneResource) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetBody() {
	_jsii_.InvokeVoid(
		d,
		"resetBody",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetCreateHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetCreateQueryParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateQueryParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetDeleteHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetDeleteQueryParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteQueryParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetIgnoreCasing() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreCasing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetIgnoreMissingProperty() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreMissingProperty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetLocks() {
	_jsii_.InvokeVoid(
		d,
		"resetLocks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetReadHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetReadHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetReadQueryParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetReadQueryParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetReplaceTriggersExternalValues() {
	_jsii_.InvokeVoid(
		d,
		"resetReplaceTriggersExternalValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetReplaceTriggersRefs() {
	_jsii_.InvokeVoid(
		d,
		"resetReplaceTriggersRefs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetResponseExportValues() {
	_jsii_.InvokeVoid(
		d,
		"resetResponseExportValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetRetry() {
	_jsii_.InvokeVoid(
		d,
		"resetRetry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetSensitiveBody() {
	_jsii_.InvokeVoid(
		d,
		"resetSensitiveBody",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetSensitiveBodyVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetSensitiveBodyVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetUpdateHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdateHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) ResetUpdateQueryParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdateQueryParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPlaneResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPlaneResource) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

