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

// checks if the ModelsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsUser{}

// ModelsUser struct for ModelsUser
type ModelsUser struct {
	Id       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewModelsUser instantiates a new ModelsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsUser() *ModelsUser {
	this := ModelsUser{}
	return &this
}

// NewModelsUserWithDefaults instantiates a new ModelsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsUserWithDefaults() *ModelsUser {
	this := ModelsUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsUser) SetId(v string) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ModelsUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ModelsUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ModelsUser) SetUsername(v string) {
	o.Username = &v
}

func (o ModelsUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableModelsUser struct {
	value *ModelsUser
	isSet bool
}

func (v NullableModelsUser) Get() *ModelsUser {
	return v.value
}

func (v *NullableModelsUser) Set(val *ModelsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsUser(val *ModelsUser) *NullableModelsUser {
	return &NullableModelsUser{value: val, isSet: true}
}

func (v NullableModelsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
