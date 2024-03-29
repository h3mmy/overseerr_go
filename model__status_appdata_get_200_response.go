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

// checks if the StatusAppdataGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusAppdataGet200Response{}

// StatusAppdataGet200Response struct for StatusAppdataGet200Response
type StatusAppdataGet200Response struct {
	AppData *bool `json:"appData,omitempty"`
	AppDataPath *string `json:"appDataPath,omitempty"`
}

// NewStatusAppdataGet200Response instantiates a new StatusAppdataGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusAppdataGet200Response() *StatusAppdataGet200Response {
	this := StatusAppdataGet200Response{}
	return &this
}

// NewStatusAppdataGet200ResponseWithDefaults instantiates a new StatusAppdataGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusAppdataGet200ResponseWithDefaults() *StatusAppdataGet200Response {
	this := StatusAppdataGet200Response{}
	return &this
}

// GetAppData returns the AppData field value if set, zero value otherwise.
func (o *StatusAppdataGet200Response) GetAppData() bool {
	if o == nil || IsNil(o.AppData) {
		var ret bool
		return ret
	}
	return *o.AppData
}

// GetAppDataOk returns a tuple with the AppData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusAppdataGet200Response) GetAppDataOk() (*bool, bool) {
	if o == nil || IsNil(o.AppData) {
		return nil, false
	}
	return o.AppData, true
}

// HasAppData returns a boolean if a field has been set.
func (o *StatusAppdataGet200Response) HasAppData() bool {
	if o != nil && !IsNil(o.AppData) {
		return true
	}

	return false
}

// SetAppData gets a reference to the given bool and assigns it to the AppData field.
func (o *StatusAppdataGet200Response) SetAppData(v bool) {
	o.AppData = &v
}

// GetAppDataPath returns the AppDataPath field value if set, zero value otherwise.
func (o *StatusAppdataGet200Response) GetAppDataPath() string {
	if o == nil || IsNil(o.AppDataPath) {
		var ret string
		return ret
	}
	return *o.AppDataPath
}

// GetAppDataPathOk returns a tuple with the AppDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusAppdataGet200Response) GetAppDataPathOk() (*string, bool) {
	if o == nil || IsNil(o.AppDataPath) {
		return nil, false
	}
	return o.AppDataPath, true
}

// HasAppDataPath returns a boolean if a field has been set.
func (o *StatusAppdataGet200Response) HasAppDataPath() bool {
	if o != nil && !IsNil(o.AppDataPath) {
		return true
	}

	return false
}

// SetAppDataPath gets a reference to the given string and assigns it to the AppDataPath field.
func (o *StatusAppdataGet200Response) SetAppDataPath(v string) {
	o.AppDataPath = &v
}

func (o StatusAppdataGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusAppdataGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppData) {
		toSerialize["appData"] = o.AppData
	}
	if !IsNil(o.AppDataPath) {
		toSerialize["appDataPath"] = o.AppDataPath
	}
	return toSerialize, nil
}

type NullableStatusAppdataGet200Response struct {
	value *StatusAppdataGet200Response
	isSet bool
}

func (v NullableStatusAppdataGet200Response) Get() *StatusAppdataGet200Response {
	return v.value
}

func (v *NullableStatusAppdataGet200Response) Set(val *StatusAppdataGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusAppdataGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusAppdataGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusAppdataGet200Response(val *StatusAppdataGet200Response) *NullableStatusAppdataGet200Response {
	return &NullableStatusAppdataGet200Response{value: val, isSet: true}
}

func (v NullableStatusAppdataGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusAppdataGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


