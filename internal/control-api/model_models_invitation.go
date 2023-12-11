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

// checks if the ModelsInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsInvitation{}

// ModelsInvitation struct for ModelsInvitation
type ModelsInvitation struct {
	// The email address to invite
	Email          *string             `json:"email,omitempty"`
	ExpiresAt      *string             `json:"expires_at,omitempty"`
	From           *ModelsUser         `json:"from,omitempty"`
	Id             *string             `json:"id,omitempty"`
	Organization   *ModelsOrganization `json:"organization,omitempty"`
	OrganizationId *string             `json:"organization_id,omitempty"`
	UserId         *string             `json:"user_id,omitempty"`
}

// NewModelsInvitation instantiates a new ModelsInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsInvitation() *ModelsInvitation {
	this := ModelsInvitation{}
	return &this
}

// NewModelsInvitationWithDefaults instantiates a new ModelsInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsInvitationWithDefaults() *ModelsInvitation {
	this := ModelsInvitation{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ModelsInvitation) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ModelsInvitation) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ModelsInvitation) SetEmail(v string) {
	o.Email = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ModelsInvitation) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ModelsInvitation) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *ModelsInvitation) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModelsInvitation) GetFrom() ModelsUser {
	if o == nil || IsNil(o.From) {
		var ret ModelsUser
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetFromOk() (*ModelsUser, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModelsInvitation) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given ModelsUser and assigns it to the From field.
func (o *ModelsInvitation) SetFrom(v ModelsUser) {
	o.From = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsInvitation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsInvitation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsInvitation) SetId(v string) {
	o.Id = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ModelsInvitation) GetOrganization() ModelsOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret ModelsOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetOrganizationOk() (*ModelsOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ModelsInvitation) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given ModelsOrganization and assigns it to the Organization field.
func (o *ModelsInvitation) SetOrganization(v ModelsOrganization) {
	o.Organization = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ModelsInvitation) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ModelsInvitation) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ModelsInvitation) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ModelsInvitation) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsInvitation) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ModelsInvitation) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ModelsInvitation) SetUserId(v string) {
	o.UserId = &v
}

func (o ModelsInvitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableModelsInvitation struct {
	value *ModelsInvitation
	isSet bool
}

func (v NullableModelsInvitation) Get() *ModelsInvitation {
	return v.value
}

func (v *NullableModelsInvitation) Set(val *ModelsInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsInvitation(val *ModelsInvitation) *NullableModelsInvitation {
	return &NullableModelsInvitation{value: val, isSet: true}
}

func (v NullableModelsInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
