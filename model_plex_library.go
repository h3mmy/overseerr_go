/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PlexLibrary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlexLibrary{}

// PlexLibrary struct for PlexLibrary
type PlexLibrary struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
}

type _PlexLibrary PlexLibrary

// NewPlexLibrary instantiates a new PlexLibrary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexLibrary(id string, name string, enabled bool) *PlexLibrary {
	this := PlexLibrary{}
	this.Id = id
	this.Name = name
	this.Enabled = enabled
	return &this
}

// NewPlexLibraryWithDefaults instantiates a new PlexLibrary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexLibraryWithDefaults() *PlexLibrary {
	this := PlexLibrary{}
	return &this
}

// GetId returns the Id field value
func (o *PlexLibrary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlexLibrary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlexLibrary) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PlexLibrary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlexLibrary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlexLibrary) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *PlexLibrary) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PlexLibrary) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PlexLibrary) SetEnabled(v bool) {
	o.Enabled = v
}

func (o PlexLibrary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlexLibrary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *PlexLibrary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPlexLibrary := _PlexLibrary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlexLibrary)

	if err != nil {
		return err
	}

	*o = PlexLibrary(varPlexLibrary)

	return err
}

type NullablePlexLibrary struct {
	value *PlexLibrary
	isSet bool
}

func (v NullablePlexLibrary) Get() *PlexLibrary {
	return v.value
}

func (v *NullablePlexLibrary) Set(val *PlexLibrary) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexLibrary) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexLibrary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexLibrary(val *PlexLibrary) *NullablePlexLibrary {
	return &NullablePlexLibrary{value: val, isSet: true}
}

func (v NullablePlexLibrary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexLibrary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

