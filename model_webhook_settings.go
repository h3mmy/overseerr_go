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

// checks if the WebhookSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSettings{}

// WebhookSettings struct for WebhookSettings
type WebhookSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Types *float32 `json:"types,omitempty"`
	Options *WebhookSettingsOptions `json:"options,omitempty"`
}

// NewWebhookSettings instantiates a new WebhookSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSettings() *WebhookSettings {
	this := WebhookSettings{}
	return &this
}

// NewWebhookSettingsWithDefaults instantiates a new WebhookSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSettingsWithDefaults() *WebhookSettings {
	this := WebhookSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WebhookSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WebhookSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WebhookSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *WebhookSettings) GetTypes() float32 {
	if o == nil || IsNil(o.Types) {
		var ret float32
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSettings) GetTypesOk() (*float32, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WebhookSettings) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given float32 and assigns it to the Types field.
func (o *WebhookSettings) SetTypes(v float32) {
	o.Types = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *WebhookSettings) GetOptions() WebhookSettingsOptions {
	if o == nil || IsNil(o.Options) {
		var ret WebhookSettingsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSettings) GetOptionsOk() (*WebhookSettingsOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *WebhookSettings) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given WebhookSettingsOptions and assigns it to the Options field.
func (o *WebhookSettings) SetOptions(v WebhookSettingsOptions) {
	o.Options = &v
}

func (o WebhookSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableWebhookSettings struct {
	value *WebhookSettings
	isSet bool
}

func (v NullableWebhookSettings) Get() *WebhookSettings {
	return v.value
}

func (v *NullableWebhookSettings) Set(val *WebhookSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSettings(val *WebhookSettings) *NullableWebhookSettings {
	return &NullableWebhookSettings{value: val, isSet: true}
}

func (v NullableWebhookSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


