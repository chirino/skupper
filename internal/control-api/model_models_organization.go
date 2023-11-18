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

// checks if the ModelsOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsOrganization{}

// ModelsOrganization struct for ModelsOrganization
type ModelsOrganization struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	OwnerId     *string `json:"owner_id,omitempty"`
}

// NewModelsOrganization instantiates a new ModelsOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsOrganization() *ModelsOrganization {
	this := ModelsOrganization{}
	return &this
}

// NewModelsOrganizationWithDefaults instantiates a new ModelsOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsOrganizationWithDefaults() *ModelsOrganization {
	this := ModelsOrganization{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsOrganization) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrganization) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsOrganization) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsOrganization) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsOrganization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrganization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsOrganization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsOrganization) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsOrganization) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrganization) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsOrganization) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsOrganization) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ModelsOrganization) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrganization) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ModelsOrganization) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ModelsOrganization) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o ModelsOrganization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	return toSerialize, nil
}

type NullableModelsOrganization struct {
	value *ModelsOrganization
	isSet bool
}

func (v NullableModelsOrganization) Get() *ModelsOrganization {
	return v.value
}

func (v *NullableModelsOrganization) Set(val *ModelsOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsOrganization(val *ModelsOrganization) *NullableModelsOrganization {
	return &NullableModelsOrganization{value: val, isSet: true}
}

func (v NullableModelsOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
