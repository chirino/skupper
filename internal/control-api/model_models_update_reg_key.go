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

// checks if the ModelsUpdateRegKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsUpdateRegKey{}

// ModelsUpdateRegKey struct for ModelsUpdateRegKey
type ModelsUpdateRegKey struct {
	// Description of the registration key.
	Description *string `json:"description,omitempty"`
	// ExpiresAt is optional, if set the registration key is only valid until the ExpiresAt time.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// SecurityGroupId is the ID of the security group to assign to the device.
	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// Settings contains general settings for the device.
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// NewModelsUpdateRegKey instantiates a new ModelsUpdateRegKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsUpdateRegKey() *ModelsUpdateRegKey {
	this := ModelsUpdateRegKey{}
	return &this
}

// NewModelsUpdateRegKeyWithDefaults instantiates a new ModelsUpdateRegKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsUpdateRegKeyWithDefaults() *ModelsUpdateRegKey {
	this := ModelsUpdateRegKey{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsUpdateRegKey) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateRegKey) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsUpdateRegKey) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsUpdateRegKey) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ModelsUpdateRegKey) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateRegKey) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ModelsUpdateRegKey) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *ModelsUpdateRegKey) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *ModelsUpdateRegKey) GetSecurityGroupId() string {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateRegKey) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *ModelsUpdateRegKey) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *ModelsUpdateRegKey) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ModelsUpdateRegKey) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUpdateRegKey) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ModelsUpdateRegKey) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *ModelsUpdateRegKey) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o ModelsUpdateRegKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsUpdateRegKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.SecurityGroupId) {
		toSerialize["security_group_id"] = o.SecurityGroupId
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableModelsUpdateRegKey struct {
	value *ModelsUpdateRegKey
	isSet bool
}

func (v NullableModelsUpdateRegKey) Get() *ModelsUpdateRegKey {
	return v.value
}

func (v *NullableModelsUpdateRegKey) Set(val *ModelsUpdateRegKey) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsUpdateRegKey) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsUpdateRegKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsUpdateRegKey(val *ModelsUpdateRegKey) *NullableModelsUpdateRegKey {
	return &NullableModelsUpdateRegKey{value: val, isSet: true}
}

func (v NullableModelsUpdateRegKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsUpdateRegKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
