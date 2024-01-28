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

// checks if the PersonResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonResult{}

// PersonResult struct for PersonResult
type PersonResult struct {
	Id *float32 `json:"id,omitempty"`
	ProfilePath *string `json:"profilePath,omitempty"`
	Adult *bool `json:"adult,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	KnownFor []PersonResultKnownForInner `json:"knownFor,omitempty"`
}

// NewPersonResult instantiates a new PersonResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonResult() *PersonResult {
	this := PersonResult{}
	var mediaType string = "person"
	this.MediaType = &mediaType
	return &this
}

// NewPersonResultWithDefaults instantiates a new PersonResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonResultWithDefaults() *PersonResult {
	this := PersonResult{}
	var mediaType string = "person"
	this.MediaType = &mediaType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersonResult) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonResult) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersonResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *PersonResult) SetId(v float32) {
	o.Id = &v
}

// GetProfilePath returns the ProfilePath field value if set, zero value otherwise.
func (o *PersonResult) GetProfilePath() string {
	if o == nil || IsNil(o.ProfilePath) {
		var ret string
		return ret
	}
	return *o.ProfilePath
}

// GetProfilePathOk returns a tuple with the ProfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonResult) GetProfilePathOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePath) {
		return nil, false
	}
	return o.ProfilePath, true
}

// HasProfilePath returns a boolean if a field has been set.
func (o *PersonResult) HasProfilePath() bool {
	if o != nil && !IsNil(o.ProfilePath) {
		return true
	}

	return false
}

// SetProfilePath gets a reference to the given string and assigns it to the ProfilePath field.
func (o *PersonResult) SetProfilePath(v string) {
	o.ProfilePath = &v
}

// GetAdult returns the Adult field value if set, zero value otherwise.
func (o *PersonResult) GetAdult() bool {
	if o == nil || IsNil(o.Adult) {
		var ret bool
		return ret
	}
	return *o.Adult
}

// GetAdultOk returns a tuple with the Adult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonResult) GetAdultOk() (*bool, bool) {
	if o == nil || IsNil(o.Adult) {
		return nil, false
	}
	return o.Adult, true
}

// HasAdult returns a boolean if a field has been set.
func (o *PersonResult) HasAdult() bool {
	if o != nil && !IsNil(o.Adult) {
		return true
	}

	return false
}

// SetAdult gets a reference to the given bool and assigns it to the Adult field.
func (o *PersonResult) SetAdult(v bool) {
	o.Adult = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *PersonResult) GetMediaType() string {
	if o == nil || IsNil(o.MediaType) {
		var ret string
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonResult) GetMediaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *PersonResult) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given string and assigns it to the MediaType field.
func (o *PersonResult) SetMediaType(v string) {
	o.MediaType = &v
}

// GetKnownFor returns the KnownFor field value if set, zero value otherwise.
func (o *PersonResult) GetKnownFor() []PersonResultKnownForInner {
	if o == nil || IsNil(o.KnownFor) {
		var ret []PersonResultKnownForInner
		return ret
	}
	return o.KnownFor
}

// GetKnownForOk returns a tuple with the KnownFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonResult) GetKnownForOk() ([]PersonResultKnownForInner, bool) {
	if o == nil || IsNil(o.KnownFor) {
		return nil, false
	}
	return o.KnownFor, true
}

// HasKnownFor returns a boolean if a field has been set.
func (o *PersonResult) HasKnownFor() bool {
	if o != nil && !IsNil(o.KnownFor) {
		return true
	}

	return false
}

// SetKnownFor gets a reference to the given []PersonResultKnownForInner and assigns it to the KnownFor field.
func (o *PersonResult) SetKnownFor(v []PersonResultKnownForInner) {
	o.KnownFor = v
}

func (o PersonResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProfilePath) {
		toSerialize["profilePath"] = o.ProfilePath
	}
	if !IsNil(o.Adult) {
		toSerialize["adult"] = o.Adult
	}
	if !IsNil(o.MediaType) {
		toSerialize["mediaType"] = o.MediaType
	}
	if !IsNil(o.KnownFor) {
		toSerialize["knownFor"] = o.KnownFor
	}
	return toSerialize, nil
}

type NullablePersonResult struct {
	value *PersonResult
	isSet bool
}

func (v NullablePersonResult) Get() *PersonResult {
	return v.value
}

func (v *NullablePersonResult) Set(val *PersonResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonResult(val *PersonResult) *NullablePersonResult {
	return &NullablePersonResult{value: val, isSet: true}
}

func (v NullablePersonResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


