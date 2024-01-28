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

// checks if the SettingsCacheGet200ResponseApiCachesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsCacheGet200ResponseApiCachesInner{}

// SettingsCacheGet200ResponseApiCachesInner struct for SettingsCacheGet200ResponseApiCachesInner
type SettingsCacheGet200ResponseApiCachesInner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Stats *SettingsCacheGet200ResponseApiCachesInnerStats `json:"stats,omitempty"`
}

// NewSettingsCacheGet200ResponseApiCachesInner instantiates a new SettingsCacheGet200ResponseApiCachesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsCacheGet200ResponseApiCachesInner() *SettingsCacheGet200ResponseApiCachesInner {
	this := SettingsCacheGet200ResponseApiCachesInner{}
	return &this
}

// NewSettingsCacheGet200ResponseApiCachesInnerWithDefaults instantiates a new SettingsCacheGet200ResponseApiCachesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsCacheGet200ResponseApiCachesInnerWithDefaults() *SettingsCacheGet200ResponseApiCachesInner {
	this := SettingsCacheGet200ResponseApiCachesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettingsCacheGet200ResponseApiCachesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCacheGet200ResponseApiCachesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SettingsCacheGet200ResponseApiCachesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SettingsCacheGet200ResponseApiCachesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SettingsCacheGet200ResponseApiCachesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCacheGet200ResponseApiCachesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SettingsCacheGet200ResponseApiCachesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SettingsCacheGet200ResponseApiCachesInner) SetName(v string) {
	o.Name = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *SettingsCacheGet200ResponseApiCachesInner) GetStats() SettingsCacheGet200ResponseApiCachesInnerStats {
	if o == nil || IsNil(o.Stats) {
		var ret SettingsCacheGet200ResponseApiCachesInnerStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCacheGet200ResponseApiCachesInner) GetStatsOk() (*SettingsCacheGet200ResponseApiCachesInnerStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *SettingsCacheGet200ResponseApiCachesInner) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given SettingsCacheGet200ResponseApiCachesInnerStats and assigns it to the Stats field.
func (o *SettingsCacheGet200ResponseApiCachesInner) SetStats(v SettingsCacheGet200ResponseApiCachesInnerStats) {
	o.Stats = &v
}

func (o SettingsCacheGet200ResponseApiCachesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsCacheGet200ResponseApiCachesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	return toSerialize, nil
}

type NullableSettingsCacheGet200ResponseApiCachesInner struct {
	value *SettingsCacheGet200ResponseApiCachesInner
	isSet bool
}

func (v NullableSettingsCacheGet200ResponseApiCachesInner) Get() *SettingsCacheGet200ResponseApiCachesInner {
	return v.value
}

func (v *NullableSettingsCacheGet200ResponseApiCachesInner) Set(val *SettingsCacheGet200ResponseApiCachesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsCacheGet200ResponseApiCachesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsCacheGet200ResponseApiCachesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsCacheGet200ResponseApiCachesInner(val *SettingsCacheGet200ResponseApiCachesInner) *NullableSettingsCacheGet200ResponseApiCachesInner {
	return &NullableSettingsCacheGet200ResponseApiCachesInner{value: val, isSet: true}
}

func (v NullableSettingsCacheGet200ResponseApiCachesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsCacheGet200ResponseApiCachesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


