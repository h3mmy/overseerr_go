/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UserUserIdSettingsPermissionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUserIdSettingsPermissionsGet200Response{}

// UserUserIdSettingsPermissionsGet200Response struct for UserUserIdSettingsPermissionsGet200Response
type UserUserIdSettingsPermissionsGet200Response struct {
	Permissions *float32 `json:"permissions,omitempty"`
}

// NewUserUserIdSettingsPermissionsGet200Response instantiates a new UserUserIdSettingsPermissionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUserIdSettingsPermissionsGet200Response() *UserUserIdSettingsPermissionsGet200Response {
	this := UserUserIdSettingsPermissionsGet200Response{}
	return &this
}

// NewUserUserIdSettingsPermissionsGet200ResponseWithDefaults instantiates a new UserUserIdSettingsPermissionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUserIdSettingsPermissionsGet200ResponseWithDefaults() *UserUserIdSettingsPermissionsGet200Response {
	this := UserUserIdSettingsPermissionsGet200Response{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserUserIdSettingsPermissionsGet200Response) GetPermissions() float32 {
	if o == nil || IsNil(o.Permissions) {
		var ret float32
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserIdSettingsPermissionsGet200Response) GetPermissionsOk() (*float32, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserUserIdSettingsPermissionsGet200Response) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given float32 and assigns it to the Permissions field.
func (o *UserUserIdSettingsPermissionsGet200Response) SetPermissions(v float32) {
	o.Permissions = &v
}

func (o UserUserIdSettingsPermissionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUserIdSettingsPermissionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableUserUserIdSettingsPermissionsGet200Response struct {
	value *UserUserIdSettingsPermissionsGet200Response
	isSet bool
}

func (v NullableUserUserIdSettingsPermissionsGet200Response) Get() *UserUserIdSettingsPermissionsGet200Response {
	return v.value
}

func (v *NullableUserUserIdSettingsPermissionsGet200Response) Set(val *UserUserIdSettingsPermissionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUserIdSettingsPermissionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUserIdSettingsPermissionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUserIdSettingsPermissionsGet200Response(val *UserUserIdSettingsPermissionsGet200Response) *NullableUserUserIdSettingsPermissionsGet200Response {
	return &NullableUserUserIdSettingsPermissionsGet200Response{value: val, isSet: true}
}

func (v NullableUserUserIdSettingsPermissionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUserIdSettingsPermissionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


