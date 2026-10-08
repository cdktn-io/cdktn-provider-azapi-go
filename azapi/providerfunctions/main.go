// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-azapi.providerFunctions.AzapiProviderFunctions",
		reflect.TypeOf((*AzapiProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildResourceId", GoMethod: "BuildResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "extensionResourceId", GoMethod: "ExtensionResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "managementGroupResourceId", GoMethod: "ManagementGroupResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "parseResourceId", GoMethod: "ParseResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "resourceGroupResourceId", GoMethod: "ResourceGroupResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "snake2Camel", GoMethod: "Snake2Camel"},
			_jsii_.MemberMethod{JsiiMethod: "subscriptionResourceId", GoMethod: "SubscriptionResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "tenantResourceId", GoMethod: "TenantResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "uniqueString", GoMethod: "UniqueString"},
		},
		func() interface{} {
			return &jsiiProxy_AzapiProviderFunctions{}
		},
	)
}
