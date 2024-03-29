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

// checks if the SettingsPlexUsersGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsPlexUsersGet200ResponseInner{}

// SettingsPlexUsersGet200ResponseInner struct for SettingsPlexUsersGet200ResponseInner
type SettingsPlexUsersGet200ResponseInner struct {
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Username *string `json:"username,omitempty"`
	Email *string `json:"email,omitempty"`
	Thumb *string `json:"thumb,omitempty"`
}

// NewSettingsPlexUsersGet200ResponseInner instantiates a new SettingsPlexUsersGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsPlexUsersGet200ResponseInner() *SettingsPlexUsersGet200ResponseInner {
	this := SettingsPlexUsersGet200ResponseInner{}
	return &this
}

// NewSettingsPlexUsersGet200ResponseInnerWithDefaults instantiates a new SettingsPlexUsersGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsPlexUsersGet200ResponseInnerWithDefaults() *SettingsPlexUsersGet200ResponseInner {
	this := SettingsPlexUsersGet200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettingsPlexUsersGet200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexUsersGet200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SettingsPlexUsersGet200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SettingsPlexUsersGet200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SettingsPlexUsersGet200ResponseInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexUsersGet200ResponseInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SettingsPlexUsersGet200ResponseInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SettingsPlexUsersGet200ResponseInner) SetTitle(v string) {
	o.Title = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SettingsPlexUsersGet200ResponseInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexUsersGet200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SettingsPlexUsersGet200ResponseInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SettingsPlexUsersGet200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SettingsPlexUsersGet200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexUsersGet200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SettingsPlexUsersGet200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SettingsPlexUsersGet200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetThumb returns the Thumb field value if set, zero value otherwise.
func (o *SettingsPlexUsersGet200ResponseInner) GetThumb() string {
	if o == nil || IsNil(o.Thumb) {
		var ret string
		return ret
	}
	return *o.Thumb
}

// GetThumbOk returns a tuple with the Thumb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexUsersGet200ResponseInner) GetThumbOk() (*string, bool) {
	if o == nil || IsNil(o.Thumb) {
		return nil, false
	}
	return o.Thumb, true
}

// HasThumb returns a boolean if a field has been set.
func (o *SettingsPlexUsersGet200ResponseInner) HasThumb() bool {
	if o != nil && !IsNil(o.Thumb) {
		return true
	}

	return false
}

// SetThumb gets a reference to the given string and assigns it to the Thumb field.
func (o *SettingsPlexUsersGet200ResponseInner) SetThumb(v string) {
	o.Thumb = &v
}

func (o SettingsPlexUsersGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsPlexUsersGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Thumb) {
		toSerialize["thumb"] = o.Thumb
	}
	return toSerialize, nil
}

type NullableSettingsPlexUsersGet200ResponseInner struct {
	value *SettingsPlexUsersGet200ResponseInner
	isSet bool
}

func (v NullableSettingsPlexUsersGet200ResponseInner) Get() *SettingsPlexUsersGet200ResponseInner {
	return v.value
}

func (v *NullableSettingsPlexUsersGet200ResponseInner) Set(val *SettingsPlexUsersGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsPlexUsersGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsPlexUsersGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsPlexUsersGet200ResponseInner(val *SettingsPlexUsersGet200ResponseInner) *NullableSettingsPlexUsersGet200ResponseInner {
	return &NullableSettingsPlexUsersGet200ResponseInner{value: val, isSet: true}
}

func (v NullableSettingsPlexUsersGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsPlexUsersGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


