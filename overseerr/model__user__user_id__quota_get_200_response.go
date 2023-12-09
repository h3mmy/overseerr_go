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

// checks if the UserUserIdQuotaGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUserIdQuotaGet200Response{}

// UserUserIdQuotaGet200Response struct for UserUserIdQuotaGet200Response
type UserUserIdQuotaGet200Response struct {
	Movie *UserUserIdQuotaGet200ResponseMovie `json:"movie,omitempty"`
	Tv *UserUserIdQuotaGet200ResponseMovie `json:"tv,omitempty"`
}

// NewUserUserIdQuotaGet200Response instantiates a new UserUserIdQuotaGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUserIdQuotaGet200Response() *UserUserIdQuotaGet200Response {
	this := UserUserIdQuotaGet200Response{}
	return &this
}

// NewUserUserIdQuotaGet200ResponseWithDefaults instantiates a new UserUserIdQuotaGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUserIdQuotaGet200ResponseWithDefaults() *UserUserIdQuotaGet200Response {
	this := UserUserIdQuotaGet200Response{}
	return &this
}

// GetMovie returns the Movie field value if set, zero value otherwise.
func (o *UserUserIdQuotaGet200Response) GetMovie() UserUserIdQuotaGet200ResponseMovie {
	if o == nil || IsNil(o.Movie) {
		var ret UserUserIdQuotaGet200ResponseMovie
		return ret
	}
	return *o.Movie
}

// GetMovieOk returns a tuple with the Movie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserIdQuotaGet200Response) GetMovieOk() (*UserUserIdQuotaGet200ResponseMovie, bool) {
	if o == nil || IsNil(o.Movie) {
		return nil, false
	}
	return o.Movie, true
}

// HasMovie returns a boolean if a field has been set.
func (o *UserUserIdQuotaGet200Response) HasMovie() bool {
	if o != nil && !IsNil(o.Movie) {
		return true
	}

	return false
}

// SetMovie gets a reference to the given UserUserIdQuotaGet200ResponseMovie and assigns it to the Movie field.
func (o *UserUserIdQuotaGet200Response) SetMovie(v UserUserIdQuotaGet200ResponseMovie) {
	o.Movie = &v
}

// GetTv returns the Tv field value if set, zero value otherwise.
func (o *UserUserIdQuotaGet200Response) GetTv() UserUserIdQuotaGet200ResponseMovie {
	if o == nil || IsNil(o.Tv) {
		var ret UserUserIdQuotaGet200ResponseMovie
		return ret
	}
	return *o.Tv
}

// GetTvOk returns a tuple with the Tv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserIdQuotaGet200Response) GetTvOk() (*UserUserIdQuotaGet200ResponseMovie, bool) {
	if o == nil || IsNil(o.Tv) {
		return nil, false
	}
	return o.Tv, true
}

// HasTv returns a boolean if a field has been set.
func (o *UserUserIdQuotaGet200Response) HasTv() bool {
	if o != nil && !IsNil(o.Tv) {
		return true
	}

	return false
}

// SetTv gets a reference to the given UserUserIdQuotaGet200ResponseMovie and assigns it to the Tv field.
func (o *UserUserIdQuotaGet200Response) SetTv(v UserUserIdQuotaGet200ResponseMovie) {
	o.Tv = &v
}

func (o UserUserIdQuotaGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUserIdQuotaGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Movie) {
		toSerialize["movie"] = o.Movie
	}
	if !IsNil(o.Tv) {
		toSerialize["tv"] = o.Tv
	}
	return toSerialize, nil
}

type NullableUserUserIdQuotaGet200Response struct {
	value *UserUserIdQuotaGet200Response
	isSet bool
}

func (v NullableUserUserIdQuotaGet200Response) Get() *UserUserIdQuotaGet200Response {
	return v.value
}

func (v *NullableUserUserIdQuotaGet200Response) Set(val *UserUserIdQuotaGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUserIdQuotaGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUserIdQuotaGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUserIdQuotaGet200Response(val *UserUserIdQuotaGet200Response) *NullableUserUserIdQuotaGet200Response {
	return &NullableUserUserIdQuotaGet200Response{value: val, isSet: true}
}

func (v NullableUserUserIdQuotaGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUserIdQuotaGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

