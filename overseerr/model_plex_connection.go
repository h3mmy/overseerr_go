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

// checks if the PlexConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlexConnection{}

// PlexConnection struct for PlexConnection
type PlexConnection struct {
	Protocol string `json:"protocol"`
	Address string `json:"address"`
	Port float32 `json:"port"`
	Uri string `json:"uri"`
	Local bool `json:"local"`
	Status *float32 `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

type _PlexConnection PlexConnection

// NewPlexConnection instantiates a new PlexConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexConnection(protocol string, address string, port float32, uri string, local bool) *PlexConnection {
	this := PlexConnection{}
	this.Protocol = protocol
	this.Address = address
	this.Port = port
	this.Uri = uri
	this.Local = local
	return &this
}

// NewPlexConnectionWithDefaults instantiates a new PlexConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexConnectionWithDefaults() *PlexConnection {
	this := PlexConnection{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *PlexConnection) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *PlexConnection) SetProtocol(v string) {
	o.Protocol = v
}

// GetAddress returns the Address field value
func (o *PlexConnection) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PlexConnection) SetAddress(v string) {
	o.Address = v
}

// GetPort returns the Port field value
func (o *PlexConnection) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PlexConnection) SetPort(v float32) {
	o.Port = v
}

// GetUri returns the Uri field value
func (o *PlexConnection) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *PlexConnection) SetUri(v string) {
	o.Uri = v
}

// GetLocal returns the Local field value
func (o *PlexConnection) GetLocal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Local
}

// GetLocalOk returns a tuple with the Local field value
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetLocalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Local, true
}

// SetLocal sets field value
func (o *PlexConnection) SetLocal(v bool) {
	o.Local = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlexConnection) GetStatus() float32 {
	if o == nil || IsNil(o.Status) {
		var ret float32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetStatusOk() (*float32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlexConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given float32 and assigns it to the Status field.
func (o *PlexConnection) SetStatus(v float32) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PlexConnection) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexConnection) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PlexConnection) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PlexConnection) SetMessage(v string) {
	o.Message = &v
}

func (o PlexConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlexConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["address"] = o.Address
	toSerialize["port"] = o.Port
	toSerialize["uri"] = o.Uri
	toSerialize["local"] = o.Local
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

func (o *PlexConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol",
		"address",
		"port",
		"uri",
		"local",
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

	varPlexConnection := _PlexConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlexConnection)

	if err != nil {
		return err
	}

	*o = PlexConnection(varPlexConnection)

	return err
}

type NullablePlexConnection struct {
	value *PlexConnection
	isSet bool
}

func (v NullablePlexConnection) Get() *PlexConnection {
	return v.value
}

func (v *NullablePlexConnection) Set(val *PlexConnection) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexConnection) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexConnection(val *PlexConnection) *NullablePlexConnection {
	return &NullablePlexConnection{value: val, isSet: true}
}

func (v NullablePlexConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


