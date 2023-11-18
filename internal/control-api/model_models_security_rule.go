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

// checks if the ModelsSecurityRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSecurityRule{}

// ModelsSecurityRule struct for ModelsSecurityRule
type ModelsSecurityRule struct {
	FromPort   *int32   `json:"from_port,omitempty"`
	IpProtocol *string  `json:"ip_protocol,omitempty"`
	IpRanges   []string `json:"ip_ranges,omitempty"`
	ToPort     *int32   `json:"to_port,omitempty"`
}

// NewModelsSecurityRule instantiates a new ModelsSecurityRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSecurityRule() *ModelsSecurityRule {
	this := ModelsSecurityRule{}
	return &this
}

// NewModelsSecurityRuleWithDefaults instantiates a new ModelsSecurityRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSecurityRuleWithDefaults() *ModelsSecurityRule {
	this := ModelsSecurityRule{}
	return &this
}

// GetFromPort returns the FromPort field value if set, zero value otherwise.
func (o *ModelsSecurityRule) GetFromPort() int32 {
	if o == nil || IsNil(o.FromPort) {
		var ret int32
		return ret
	}
	return *o.FromPort
}

// GetFromPortOk returns a tuple with the FromPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecurityRule) GetFromPortOk() (*int32, bool) {
	if o == nil || IsNil(o.FromPort) {
		return nil, false
	}
	return o.FromPort, true
}

// HasFromPort returns a boolean if a field has been set.
func (o *ModelsSecurityRule) HasFromPort() bool {
	if o != nil && !IsNil(o.FromPort) {
		return true
	}

	return false
}

// SetFromPort gets a reference to the given int32 and assigns it to the FromPort field.
func (o *ModelsSecurityRule) SetFromPort(v int32) {
	o.FromPort = &v
}

// GetIpProtocol returns the IpProtocol field value if set, zero value otherwise.
func (o *ModelsSecurityRule) GetIpProtocol() string {
	if o == nil || IsNil(o.IpProtocol) {
		var ret string
		return ret
	}
	return *o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecurityRule) GetIpProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.IpProtocol) {
		return nil, false
	}
	return o.IpProtocol, true
}

// HasIpProtocol returns a boolean if a field has been set.
func (o *ModelsSecurityRule) HasIpProtocol() bool {
	if o != nil && !IsNil(o.IpProtocol) {
		return true
	}

	return false
}

// SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.
func (o *ModelsSecurityRule) SetIpProtocol(v string) {
	o.IpProtocol = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *ModelsSecurityRule) GetIpRanges() []string {
	if o == nil || IsNil(o.IpRanges) {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecurityRule) GetIpRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *ModelsSecurityRule) HasIpRanges() bool {
	if o != nil && !IsNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *ModelsSecurityRule) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetToPort returns the ToPort field value if set, zero value otherwise.
func (o *ModelsSecurityRule) GetToPort() int32 {
	if o == nil || IsNil(o.ToPort) {
		var ret int32
		return ret
	}
	return *o.ToPort
}

// GetToPortOk returns a tuple with the ToPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecurityRule) GetToPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ToPort) {
		return nil, false
	}
	return o.ToPort, true
}

// HasToPort returns a boolean if a field has been set.
func (o *ModelsSecurityRule) HasToPort() bool {
	if o != nil && !IsNil(o.ToPort) {
		return true
	}

	return false
}

// SetToPort gets a reference to the given int32 and assigns it to the ToPort field.
func (o *ModelsSecurityRule) SetToPort(v int32) {
	o.ToPort = &v
}

func (o ModelsSecurityRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSecurityRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromPort) {
		toSerialize["from_port"] = o.FromPort
	}
	if !IsNil(o.IpProtocol) {
		toSerialize["ip_protocol"] = o.IpProtocol
	}
	if !IsNil(o.IpRanges) {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if !IsNil(o.ToPort) {
		toSerialize["to_port"] = o.ToPort
	}
	return toSerialize, nil
}

type NullableModelsSecurityRule struct {
	value *ModelsSecurityRule
	isSet bool
}

func (v NullableModelsSecurityRule) Get() *ModelsSecurityRule {
	return v.value
}

func (v *NullableModelsSecurityRule) Set(val *ModelsSecurityRule) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSecurityRule) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSecurityRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSecurityRule(val *ModelsSecurityRule) *NullableModelsSecurityRule {
	return &NullableModelsSecurityRule{value: val, isSet: true}
}

func (v NullableModelsSecurityRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSecurityRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
