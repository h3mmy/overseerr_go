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

// checks if the IssueCommentCommentIdPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueCommentCommentIdPutRequest{}

// IssueCommentCommentIdPutRequest struct for IssueCommentCommentIdPutRequest
type IssueCommentCommentIdPutRequest struct {
	Message *string `json:"message,omitempty"`
}

// NewIssueCommentCommentIdPutRequest instantiates a new IssueCommentCommentIdPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueCommentCommentIdPutRequest() *IssueCommentCommentIdPutRequest {
	this := IssueCommentCommentIdPutRequest{}
	return &this
}

// NewIssueCommentCommentIdPutRequestWithDefaults instantiates a new IssueCommentCommentIdPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueCommentCommentIdPutRequestWithDefaults() *IssueCommentCommentIdPutRequest {
	this := IssueCommentCommentIdPutRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IssueCommentCommentIdPutRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCommentCommentIdPutRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IssueCommentCommentIdPutRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IssueCommentCommentIdPutRequest) SetMessage(v string) {
	o.Message = &v
}

func (o IssueCommentCommentIdPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueCommentCommentIdPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableIssueCommentCommentIdPutRequest struct {
	value *IssueCommentCommentIdPutRequest
	isSet bool
}

func (v NullableIssueCommentCommentIdPutRequest) Get() *IssueCommentCommentIdPutRequest {
	return v.value
}

func (v *NullableIssueCommentCommentIdPutRequest) Set(val *IssueCommentCommentIdPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueCommentCommentIdPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueCommentCommentIdPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueCommentCommentIdPutRequest(val *IssueCommentCommentIdPutRequest) *NullableIssueCommentCommentIdPutRequest {
	return &NullableIssueCommentCommentIdPutRequest{value: val, isSet: true}
}

func (v NullableIssueCommentCommentIdPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueCommentCommentIdPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


