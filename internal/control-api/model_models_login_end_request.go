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

// checks if the ModelsLoginEndRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsLoginEndRequest{}

// ModelsLoginEndRequest struct for ModelsLoginEndRequest
type ModelsLoginEndRequest struct {
	RequestUrl *string `json:"request_url,omitempty"`
}

// NewModelsLoginEndRequest instantiates a new ModelsLoginEndRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsLoginEndRequest() *ModelsLoginEndRequest {
	this := ModelsLoginEndRequest{}
	return &this
}

// NewModelsLoginEndRequestWithDefaults instantiates a new ModelsLoginEndRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsLoginEndRequestWithDefaults() *ModelsLoginEndRequest {
	this := ModelsLoginEndRequest{}
	return &this
}

// GetRequestUrl returns the RequestUrl field value if set, zero value otherwise.
func (o *ModelsLoginEndRequest) GetRequestUrl() string {
	if o == nil || IsNil(o.RequestUrl) {
		var ret string
		return ret
	}
	return *o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLoginEndRequest) GetRequestUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RequestUrl) {
		return nil, false
	}
	return o.RequestUrl, true
}

// HasRequestUrl returns a boolean if a field has been set.
func (o *ModelsLoginEndRequest) HasRequestUrl() bool {
	if o != nil && !IsNil(o.RequestUrl) {
		return true
	}

	return false
}

// SetRequestUrl gets a reference to the given string and assigns it to the RequestUrl field.
func (o *ModelsLoginEndRequest) SetRequestUrl(v string) {
	o.RequestUrl = &v
}

func (o ModelsLoginEndRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsLoginEndRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestUrl) {
		toSerialize["request_url"] = o.RequestUrl
	}
	return toSerialize, nil
}

type NullableModelsLoginEndRequest struct {
	value *ModelsLoginEndRequest
	isSet bool
}

func (v NullableModelsLoginEndRequest) Get() *ModelsLoginEndRequest {
	return v.value
}

func (v *NullableModelsLoginEndRequest) Set(val *ModelsLoginEndRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsLoginEndRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsLoginEndRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsLoginEndRequest(val *ModelsLoginEndRequest) *NullableModelsLoginEndRequest {
	return &NullableModelsLoginEndRequest{value: val, isSet: true}
}

func (v NullableModelsLoginEndRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsLoginEndRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
