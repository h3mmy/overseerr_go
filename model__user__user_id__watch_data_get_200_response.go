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

// checks if the UserUserIdWatchDataGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUserIdWatchDataGet200Response{}

// UserUserIdWatchDataGet200Response struct for UserUserIdWatchDataGet200Response
type UserUserIdWatchDataGet200Response struct {
	RecentlyWatched []MediaInfo `json:"recentlyWatched,omitempty"`
	PlayCount *float32 `json:"playCount,omitempty"`
}

// NewUserUserIdWatchDataGet200Response instantiates a new UserUserIdWatchDataGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUserIdWatchDataGet200Response() *UserUserIdWatchDataGet200Response {
	this := UserUserIdWatchDataGet200Response{}
	return &this
}

// NewUserUserIdWatchDataGet200ResponseWithDefaults instantiates a new UserUserIdWatchDataGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUserIdWatchDataGet200ResponseWithDefaults() *UserUserIdWatchDataGet200Response {
	this := UserUserIdWatchDataGet200Response{}
	return &this
}

// GetRecentlyWatched returns the RecentlyWatched field value if set, zero value otherwise.
func (o *UserUserIdWatchDataGet200Response) GetRecentlyWatched() []MediaInfo {
	if o == nil || IsNil(o.RecentlyWatched) {
		var ret []MediaInfo
		return ret
	}
	return o.RecentlyWatched
}

// GetRecentlyWatchedOk returns a tuple with the RecentlyWatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserIdWatchDataGet200Response) GetRecentlyWatchedOk() ([]MediaInfo, bool) {
	if o == nil || IsNil(o.RecentlyWatched) {
		return nil, false
	}
	return o.RecentlyWatched, true
}

// HasRecentlyWatched returns a boolean if a field has been set.
func (o *UserUserIdWatchDataGet200Response) HasRecentlyWatched() bool {
	if o != nil && !IsNil(o.RecentlyWatched) {
		return true
	}

	return false
}

// SetRecentlyWatched gets a reference to the given []MediaInfo and assigns it to the RecentlyWatched field.
func (o *UserUserIdWatchDataGet200Response) SetRecentlyWatched(v []MediaInfo) {
	o.RecentlyWatched = v
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise.
func (o *UserUserIdWatchDataGet200Response) GetPlayCount() float32 {
	if o == nil || IsNil(o.PlayCount) {
		var ret float32
		return ret
	}
	return *o.PlayCount
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserIdWatchDataGet200Response) GetPlayCountOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount) {
		return nil, false
	}
	return o.PlayCount, true
}

// HasPlayCount returns a boolean if a field has been set.
func (o *UserUserIdWatchDataGet200Response) HasPlayCount() bool {
	if o != nil && !IsNil(o.PlayCount) {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given float32 and assigns it to the PlayCount field.
func (o *UserUserIdWatchDataGet200Response) SetPlayCount(v float32) {
	o.PlayCount = &v
}

func (o UserUserIdWatchDataGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUserIdWatchDataGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecentlyWatched) {
		toSerialize["recentlyWatched"] = o.RecentlyWatched
	}
	if !IsNil(o.PlayCount) {
		toSerialize["playCount"] = o.PlayCount
	}
	return toSerialize, nil
}

type NullableUserUserIdWatchDataGet200Response struct {
	value *UserUserIdWatchDataGet200Response
	isSet bool
}

func (v NullableUserUserIdWatchDataGet200Response) Get() *UserUserIdWatchDataGet200Response {
	return v.value
}

func (v *NullableUserUserIdWatchDataGet200Response) Set(val *UserUserIdWatchDataGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUserIdWatchDataGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUserIdWatchDataGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUserIdWatchDataGet200Response(val *UserUserIdWatchDataGet200Response) *NullableUserUserIdWatchDataGet200Response {
	return &NullableUserUserIdWatchDataGet200Response{value: val, isSet: true}
}

func (v NullableUserUserIdWatchDataGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUserIdWatchDataGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


