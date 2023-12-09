/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserUserIdSettingsPasswordGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUserIdSettingsPasswordGet200Response{}

// UserUserIdSettingsPasswordGet200Response struct for UserUserIdSettingsPasswordGet200Response
type UserUserIdSettingsPasswordGet200Response struct {
	HasPassword *bool `json:"hasPassword,omitempty"`
}

// NewUserUserIdSettingsPasswordGet200Response instantiates a new UserUserIdSettingsPasswordGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUserIdSettingsPasswordGet200Response() *UserUserIdSettingsPasswordGet200Response {
	this := UserUserIdSettingsPasswordGet200Response{}
	return &this
}

// NewUserUserIdSettingsPasswordGet200ResponseWithDefaults instantiates a new UserUserIdSettingsPasswordGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUserIdSettingsPasswordGet200ResponseWithDefaults() *UserUserIdSettingsPasswordGet200Response {
	this := UserUserIdSettingsPasswordGet200Response{}
	return &this
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *UserUserIdSettingsPasswordGet200Response) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserIdSettingsPasswordGet200Response) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *UserUserIdSettingsPasswordGet200Response) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *UserUserIdSettingsPasswordGet200Response) SetHasPassword(v bool) {
	o.HasPassword = &v
}

func (o UserUserIdSettingsPasswordGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUserIdSettingsPasswordGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}
	return toSerialize, nil
}

type NullableUserUserIdSettingsPasswordGet200Response struct {
	value *UserUserIdSettingsPasswordGet200Response
	isSet bool
}

func (v NullableUserUserIdSettingsPasswordGet200Response) Get() *UserUserIdSettingsPasswordGet200Response {
	return v.value
}

func (v *NullableUserUserIdSettingsPasswordGet200Response) Set(val *UserUserIdSettingsPasswordGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUserIdSettingsPasswordGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUserIdSettingsPasswordGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUserIdSettingsPasswordGet200Response(val *UserUserIdSettingsPasswordGet200Response) *NullableUserUserIdSettingsPasswordGet200Response {
	return &NullableUserUserIdSettingsPasswordGet200Response{value: val, isSet: true}
}

func (v NullableUserUserIdSettingsPasswordGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUserIdSettingsPasswordGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

