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

// checks if the ModelsNotAllowedError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsNotAllowedError{}

// ModelsNotAllowedError struct for ModelsNotAllowedError
type ModelsNotAllowedError struct {
	Error  *string `json:"error,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewModelsNotAllowedError instantiates a new ModelsNotAllowedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsNotAllowedError() *ModelsNotAllowedError {
	this := ModelsNotAllowedError{}
	return &this
}

// NewModelsNotAllowedErrorWithDefaults instantiates a new ModelsNotAllowedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsNotAllowedErrorWithDefaults() *ModelsNotAllowedError {
	this := ModelsNotAllowedError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelsNotAllowedError) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsNotAllowedError) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelsNotAllowedError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModelsNotAllowedError) SetError(v string) {
	o.Error = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ModelsNotAllowedError) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsNotAllowedError) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ModelsNotAllowedError) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ModelsNotAllowedError) SetReason(v string) {
	o.Reason = &v
}

func (o ModelsNotAllowedError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsNotAllowedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableModelsNotAllowedError struct {
	value *ModelsNotAllowedError
	isSet bool
}

func (v NullableModelsNotAllowedError) Get() *ModelsNotAllowedError {
	return v.value
}

func (v *NullableModelsNotAllowedError) Set(val *ModelsNotAllowedError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsNotAllowedError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsNotAllowedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsNotAllowedError(val *ModelsNotAllowedError) *NullableModelsNotAllowedError {
	return &NullableModelsNotAllowedError{value: val, isSet: true}
}

func (v NullableModelsNotAllowedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsNotAllowedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
