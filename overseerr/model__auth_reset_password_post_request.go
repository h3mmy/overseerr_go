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

// checks if the AuthResetPasswordPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthResetPasswordPostRequest{}

// AuthResetPasswordPostRequest struct for AuthResetPasswordPostRequest
type AuthResetPasswordPostRequest struct {
	Email string `json:"email"`
}

type _AuthResetPasswordPostRequest AuthResetPasswordPostRequest

// NewAuthResetPasswordPostRequest instantiates a new AuthResetPasswordPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthResetPasswordPostRequest(email string) *AuthResetPasswordPostRequest {
	this := AuthResetPasswordPostRequest{}
	this.Email = email
	return &this
}

// NewAuthResetPasswordPostRequestWithDefaults instantiates a new AuthResetPasswordPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthResetPasswordPostRequestWithDefaults() *AuthResetPasswordPostRequest {
	this := AuthResetPasswordPostRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *AuthResetPasswordPostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AuthResetPasswordPostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AuthResetPasswordPostRequest) SetEmail(v string) {
	o.Email = v
}

func (o AuthResetPasswordPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthResetPasswordPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *AuthResetPasswordPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varAuthResetPasswordPostRequest := _AuthResetPasswordPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthResetPasswordPostRequest)

	if err != nil {
		return err
	}

	*o = AuthResetPasswordPostRequest(varAuthResetPasswordPostRequest)

	return err
}

type NullableAuthResetPasswordPostRequest struct {
	value *AuthResetPasswordPostRequest
	isSet bool
}

func (v NullableAuthResetPasswordPostRequest) Get() *AuthResetPasswordPostRequest {
	return v.value
}

func (v *NullableAuthResetPasswordPostRequest) Set(val *AuthResetPasswordPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResetPasswordPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResetPasswordPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResetPasswordPostRequest(val *AuthResetPasswordPostRequest) *NullableAuthResetPasswordPostRequest {
	return &NullableAuthResetPasswordPostRequest{value: val, isSet: true}
}

func (v NullableAuthResetPasswordPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResetPasswordPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


