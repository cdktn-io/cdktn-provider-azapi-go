// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azapi-go/azapi/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azapi-go/azapi/resource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource azapi_resource}.
type Resource interface {
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
	Identity() ResourceIdentityList
	IdentityInput() interface{}
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	IgnoreBodyChanges() *[]*string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetIgnoreBodyChanges(val *[]*string)
	IgnoreBodyChangesInput() *[]*string
	IgnoreCasing() interface{}
	SetIgnoreCasing(val interface{})
	IgnoreCasingInput() interface{}
	IgnoreMissingProperty() interface{}
	SetIgnoreMissingProperty(val interface{})
	IgnoreMissingPropertyInput() interface{}
	IgnoreNullProperty() interface{}
	SetIgnoreNullProperty(val interface{})
	IgnoreNullPropertyInput() interface{}
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	Retry() ResourceRetryOutputReference
	RetryInput() interface{}
	SchemaValidationEnabled() interface{}
	SetSchemaValidationEnabled(val interface{})
	SchemaValidationEnabledInput() interface{}
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SensitiveBody() *map[string]interface{}
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetSensitiveBody(val *map[string]interface{})
	SensitiveBodyInput() *map[string]interface{}
	SensitiveBodyVersion() *map[string]*string
	SetSensitiveBodyVersion(val *map[string]*string)
	SensitiveBodyVersionInput() *map[string]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ResourceTimeoutsOutputReference
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
	PutIdentity(value interface{})
	PutRetry(value *ResourceRetry)
	PutTimeouts(value *ResourceTimeouts)
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
	ResetIdentity()
	ResetIgnoreBodyChanges()
	ResetIgnoreCasing()
	ResetIgnoreMissingProperty()
	ResetIgnoreNullProperty()
	ResetIgnoreOtherItemsInList()
	ResetListUniqueIdProperty()
	ResetLocation()
	ResetLocks()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentId()
	ResetReadHeaders()
	ResetReadQueryParameters()
	ResetReplaceTriggersExternalValues()
	ResetReplaceTriggersRefs()
	ResetResponseExportValues()
	ResetRetry()
	ResetSchemaValidationEnabled()
	ResetSensitiveBody()
	ResetSensitiveBodyVersion()
	ResetTags()
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

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Resource) Body() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) BodyInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) CreateHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"createHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) CreateHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"createHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) CreateQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) CreateQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DeleteHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deleteHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DeleteHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deleteHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DeleteQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DeleteQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Identity() ResourceIdentityList {
	var returns ResourceIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreBodyChanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreBodyChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreBodyChangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreBodyChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreCasing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCasing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreCasingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCasingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreMissingProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreMissingPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreNullProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNullProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreNullPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNullPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreOtherItemsInList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreOtherItemsInList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) IgnoreOtherItemsInListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreOtherItemsInListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ListUniqueIdProperty() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listUniqueIdProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ListUniqueIdPropertyInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listUniqueIdPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Locks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) LocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Output() cdktn.AnyMap {
	var returns cdktn.AnyMap
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReadHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"readHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReadHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"readHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReadQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReadQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReplaceTriggersExternalValues() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"replaceTriggersExternalValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReplaceTriggersExternalValuesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"replaceTriggersExternalValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReplaceTriggersRefs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replaceTriggersRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ReplaceTriggersRefsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replaceTriggersRefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ResponseExportValues() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseExportValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ResponseExportValuesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseExportValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Retry() ResourceRetryOutputReference {
	var returns ResourceRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) RetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) SchemaValidationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaValidationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) SchemaValidationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaValidationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) SensitiveBody() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"sensitiveBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) SensitiveBodyInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"sensitiveBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) SensitiveBodyVersion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sensitiveBodyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) SensitiveBodyVersionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sensitiveBodyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Timeouts() ResourceTimeoutsOutputReference {
	var returns ResourceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) UpdateHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updateHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) UpdateHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updateHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) UpdateQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) UpdateQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateQueryParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource azapi_resource} Resource.
func NewResource(scope constructs.Construct, id *string, config *ResourceConfig) Resource {
	_init_.Initialize()

	if err := validateNewResourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Resource{}

	_jsii_.Create(
		"@cdktn/provider-azapi.resource.Resource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/azure/azapi/2.13.0/docs/resources/resource azapi_resource} Resource.
func NewResource_Override(r Resource, scope constructs.Construct, id *string, config *ResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azapi.resource.Resource",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Resource)SetBody(val *map[string]interface{}) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_Resource)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Resource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Resource)SetCreateHeaders(val *map[string]*string) {
	if err := j.validateSetCreateHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createHeaders",
		val,
	)
}

func (j *jsiiProxy_Resource)SetCreateQueryParameters(val interface{}) {
	if err := j.validateSetCreateQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createQueryParameters",
		val,
	)
}

func (j *jsiiProxy_Resource)SetDeleteHeaders(val *map[string]*string) {
	if err := j.validateSetDeleteHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteHeaders",
		val,
	)
}

func (j *jsiiProxy_Resource)SetDeleteQueryParameters(val interface{}) {
	if err := j.validateSetDeleteQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteQueryParameters",
		val,
	)
}

func (j *jsiiProxy_Resource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Resource)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Resource)SetIgnoreBodyChanges(val *[]*string) {
	if err := j.validateSetIgnoreBodyChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreBodyChanges",
		val,
	)
}

func (j *jsiiProxy_Resource)SetIgnoreCasing(val interface{}) {
	if err := j.validateSetIgnoreCasingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCasing",
		val,
	)
}

func (j *jsiiProxy_Resource)SetIgnoreMissingProperty(val interface{}) {
	if err := j.validateSetIgnoreMissingPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMissingProperty",
		val,
	)
}

func (j *jsiiProxy_Resource)SetIgnoreNullProperty(val interface{}) {
	if err := j.validateSetIgnoreNullPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreNullProperty",
		val,
	)
}

func (j *jsiiProxy_Resource)SetIgnoreOtherItemsInList(val *[]*string) {
	if err := j.validateSetIgnoreOtherItemsInListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOtherItemsInList",
		val,
	)
}

func (j *jsiiProxy_Resource)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Resource)SetListUniqueIdProperty(val *map[string]*string) {
	if err := j.validateSetListUniqueIdPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listUniqueIdProperty",
		val,
	)
}

func (j *jsiiProxy_Resource)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_Resource)SetLocks(val *[]*string) {
	if err := j.validateSetLocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locks",
		val,
	)
}

func (j *jsiiProxy_Resource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Resource)SetParentId(val *string) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_Resource)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Resource)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Resource)SetReadHeaders(val *map[string]*string) {
	if err := j.validateSetReadHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readHeaders",
		val,
	)
}

func (j *jsiiProxy_Resource)SetReadQueryParameters(val interface{}) {
	if err := j.validateSetReadQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readQueryParameters",
		val,
	)
}

func (j *jsiiProxy_Resource)SetReplaceTriggersExternalValues(val *map[string]interface{}) {
	if err := j.validateSetReplaceTriggersExternalValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTriggersExternalValues",
		val,
	)
}

func (j *jsiiProxy_Resource)SetReplaceTriggersRefs(val *[]*string) {
	if err := j.validateSetReplaceTriggersRefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTriggersRefs",
		val,
	)
}

func (j *jsiiProxy_Resource)SetResponseExportValues(val *map[string]interface{}) {
	if err := j.validateSetResponseExportValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseExportValues",
		val,
	)
}

func (j *jsiiProxy_Resource)SetSchemaValidationEnabled(val interface{}) {
	if err := j.validateSetSchemaValidationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaValidationEnabled",
		val,
	)
}

func (j *jsiiProxy_Resource)SetSensitiveBody(val *map[string]interface{}) {
	if err := j.validateSetSensitiveBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveBody",
		val,
	)
}

func (j *jsiiProxy_Resource)SetSensitiveBodyVersion(val *map[string]*string) {
	if err := j.validateSetSensitiveBodyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveBodyVersion",
		val,
	)
}

func (j *jsiiProxy_Resource)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Resource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Resource)SetUpdateHeaders(val *map[string]*string) {
	if err := j.validateSetUpdateHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateHeaders",
		val,
	)
}

func (j *jsiiProxy_Resource)SetUpdateQueryParameters(val interface{}) {
	if err := j.validateSetUpdateQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateQueryParameters",
		val,
	)
}

// Generates CDKTN code for importing a Resource resource upon running "cdktn plan <stack-name>".
func Resource_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateResource_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.resource.Resource",
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
func Resource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.resource.Resource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Resource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.resource.Resource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Resource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azapi.resource.Resource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Resource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azapi.resource.Resource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Resource) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Resource) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Resource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Resource) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := r.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Resource) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Resource) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Resource) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Resource) PutIdentity(value interface{}) {
	if err := r.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resource) PutRetry(value *ResourceRetry) {
	if err := r.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRetry",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resource) PutTimeouts(value *ResourceTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resource) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_Resource) ResetBody() {
	_jsii_.InvokeVoid(
		r,
		"resetBody",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetCreateHeaders() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateHeaders",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetCreateQueryParameters() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateQueryParameters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetDeleteHeaders() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteHeaders",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetDeleteQueryParameters() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteQueryParameters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetIdentity() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetIgnoreBodyChanges() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreBodyChanges",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetIgnoreCasing() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreCasing",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetIgnoreMissingProperty() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreMissingProperty",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetIgnoreNullProperty() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreNullProperty",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetIgnoreOtherItemsInList() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreOtherItemsInList",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetListUniqueIdProperty() {
	_jsii_.InvokeVoid(
		r,
		"resetListUniqueIdProperty",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetLocation() {
	_jsii_.InvokeVoid(
		r,
		"resetLocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetLocks() {
	_jsii_.InvokeVoid(
		r,
		"resetLocks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetParentId() {
	_jsii_.InvokeVoid(
		r,
		"resetParentId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetReadHeaders() {
	_jsii_.InvokeVoid(
		r,
		"resetReadHeaders",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetReadQueryParameters() {
	_jsii_.InvokeVoid(
		r,
		"resetReadQueryParameters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetReplaceTriggersExternalValues() {
	_jsii_.InvokeVoid(
		r,
		"resetReplaceTriggersExternalValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetReplaceTriggersRefs() {
	_jsii_.InvokeVoid(
		r,
		"resetReplaceTriggersRefs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetResponseExportValues() {
	_jsii_.InvokeVoid(
		r,
		"resetResponseExportValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetRetry() {
	_jsii_.InvokeVoid(
		r,
		"resetRetry",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetSchemaValidationEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetSchemaValidationEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetSensitiveBody() {
	_jsii_.InvokeVoid(
		r,
		"resetSensitiveBody",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetSensitiveBodyVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetSensitiveBodyVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetUpdateHeaders() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdateHeaders",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) ResetUpdateQueryParameters() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdateQueryParameters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

