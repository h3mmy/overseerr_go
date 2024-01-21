/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"encoding/json"
)

// checks if the SettingsNotificationsPushoverSoundsGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsNotificationsPushoverSoundsGet200ResponseInner{}

// SettingsNotificationsPushoverSoundsGet200ResponseInner struct for SettingsNotificationsPushoverSoundsGet200ResponseInner
type SettingsNotificationsPushoverSoundsGet200ResponseInner struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSettingsNotificationsPushoverSoundsGet200ResponseInner instantiates a new SettingsNotificationsPushoverSoundsGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsNotificationsPushoverSoundsGet200ResponseInner() *SettingsNotificationsPushoverSoundsGet200ResponseInner {
	this := SettingsNotificationsPushoverSoundsGet200ResponseInner{}
	return &this
}

// NewSettingsNotificationsPushoverSoundsGet200ResponseInnerWithDefaults instantiates a new SettingsNotificationsPushoverSoundsGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsNotificationsPushoverSoundsGet200ResponseInnerWithDefaults() *SettingsNotificationsPushoverSoundsGet200ResponseInner {
	this := SettingsNotificationsPushoverSoundsGet200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SettingsNotificationsPushoverSoundsGet200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

func (o SettingsNotificationsPushoverSoundsGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsNotificationsPushoverSoundsGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSettingsNotificationsPushoverSoundsGet200ResponseInner struct {
	value *SettingsNotificationsPushoverSoundsGet200ResponseInner
	isSet bool
}

func (v NullableSettingsNotificationsPushoverSoundsGet200ResponseInner) Get() *SettingsNotificationsPushoverSoundsGet200ResponseInner {
	return v.value
}

func (v *NullableSettingsNotificationsPushoverSoundsGet200ResponseInner) Set(val *SettingsNotificationsPushoverSoundsGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsNotificationsPushoverSoundsGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsNotificationsPushoverSoundsGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsNotificationsPushoverSoundsGet200ResponseInner(val *SettingsNotificationsPushoverSoundsGet200ResponseInner) *NullableSettingsNotificationsPushoverSoundsGet200ResponseInner {
	return &NullableSettingsNotificationsPushoverSoundsGet200ResponseInner{value: val, isSet: true}
}

func (v NullableSettingsNotificationsPushoverSoundsGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsNotificationsPushoverSoundsGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

