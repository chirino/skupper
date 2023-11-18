/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package control_api

import (
	"encoding/json"
)

// checks if the ModelsUpdateSecurityGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsUpdateSecurityGroup{}

// ModelsUpdateSecurityGroup struct for ModelsUpdateSecurityGroup
type ModelsUpdateSecurityGroup struct {
	Description   *string              `json:"description,omitempty"`
	InboundRules  []ModelsSecurityRule `json:"inbound_rules,omitempty"`
	OutboundRules []ModelsSecurityRule `json:"outbound_rules,omitempty"`
}

// NewModelsUpdateSecurityGroup instantiates a new ModelsUpdateSecurityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsUpdateSecurityGroup() *ModelsUpdateSecurityGroup {
	this := ModelsUpdateSecurityGroup{}
	return &this
}

// NewModelsUpdateSecurityGroupWithDefaults instantiates a new ModelsUpdateSecurityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsUpdateSecurityGroupWithDefaults() *ModelsUpdateSecurityGroup {
	this := ModelsUpdateSecurityGroup{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsUpdateSecurityGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateSecurityGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsUpdateSecurityGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsUpdateSecurityGroup) SetDescription(v string) {
	o.Description = &v
}

// GetInboundRules returns the InboundRules field value if set, zero value otherwise.
func (o *ModelsUpdateSecurityGroup) GetInboundRules() []ModelsSecurityRule {
	if o == nil || IsNil(o.InboundRules) {
		var ret []ModelsSecurityRule
		return ret
	}
	return o.InboundRules
}

// GetInboundRulesOk returns a tuple with the InboundRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateSecurityGroup) GetInboundRulesOk() ([]ModelsSecurityRule, bool) {
	if o == nil || IsNil(o.InboundRules) {
		return nil, false
	}
	return o.InboundRules, true
}

// HasInboundRules returns a boolean if a field has been set.
func (o *ModelsUpdateSecurityGroup) HasInboundRules() bool {
	if o != nil && !IsNil(o.InboundRules) {
		return true
	}

	return false
}

// SetInboundRules gets a reference to the given []ModelsSecurityRule and assigns it to the InboundRules field.
func (o *ModelsUpdateSecurityGroup) SetInboundRules(v []ModelsSecurityRule) {
	o.InboundRules = v
}

// GetOutboundRules returns the OutboundRules field value if set, zero value otherwise.
func (o *ModelsUpdateSecurityGroup) GetOutboundRules() []ModelsSecurityRule {
	if o == nil || IsNil(o.OutboundRules) {
		var ret []ModelsSecurityRule
		return ret
	}
	return o.OutboundRules
}

// GetOutboundRulesOk returns a tuple with the OutboundRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateSecurityGroup) GetOutboundRulesOk() ([]ModelsSecurityRule, bool) {
	if o == nil || IsNil(o.OutboundRules) {
		return nil, false
	}
	return o.OutboundRules, true
}

// HasOutboundRules returns a boolean if a field has been set.
func (o *ModelsUpdateSecurityGroup) HasOutboundRules() bool {
	if o != nil && !IsNil(o.OutboundRules) {
		return true
	}

	return false
}

// SetOutboundRules gets a reference to the given []ModelsSecurityRule and assigns it to the OutboundRules field.
func (o *ModelsUpdateSecurityGroup) SetOutboundRules(v []ModelsSecurityRule) {
	o.OutboundRules = v
}

func (o ModelsUpdateSecurityGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsUpdateSecurityGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InboundRules) {
		toSerialize["inbound_rules"] = o.InboundRules
	}
	if !IsNil(o.OutboundRules) {
		toSerialize["outbound_rules"] = o.OutboundRules
	}
	return toSerialize, nil
}

type NullableModelsUpdateSecurityGroup struct {
	value *ModelsUpdateSecurityGroup
	isSet bool
}

func (v NullableModelsUpdateSecurityGroup) Get() *ModelsUpdateSecurityGroup {
	return v.value
}

func (v *NullableModelsUpdateSecurityGroup) Set(val *ModelsUpdateSecurityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsUpdateSecurityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsUpdateSecurityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsUpdateSecurityGroup(val *ModelsUpdateSecurityGroup) *NullableModelsUpdateSecurityGroup {
	return &NullableModelsUpdateSecurityGroup{value: val, isSet: true}
}

func (v NullableModelsUpdateSecurityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsUpdateSecurityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
