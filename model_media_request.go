/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MediaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaRequest{}

// MediaRequest struct for MediaRequest
type MediaRequest struct {
	Id float32 `json:"id"`
	// Status of the request. 1 = PENDING APPROVAL, 2 = APPROVED, 3 = DECLINED
	Status float32 `json:"status"`
	Media *MediaInfo `json:"media,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	RequestedBy *User `json:"requestedBy,omitempty"`
	ModifiedBy *MediaRequestModifiedBy `json:"modifiedBy,omitempty"`
	Is4k *bool `json:"is4k,omitempty"`
	ServerId *float32 `json:"serverId,omitempty"`
	ProfileId *float32 `json:"profileId,omitempty"`
	RootFolder *string `json:"rootFolder,omitempty"`
}

type _MediaRequest MediaRequest

// NewMediaRequest instantiates a new MediaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaRequest(id float32, status float32) *MediaRequest {
	this := MediaRequest{}
	this.Id = id
	this.Status = status
	return &this
}

// NewMediaRequestWithDefaults instantiates a new MediaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaRequestWithDefaults() *MediaRequest {
	this := MediaRequest{}
	return &this
}

// GetId returns the Id field value
func (o *MediaRequest) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MediaRequest) SetId(v float32) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *MediaRequest) GetStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MediaRequest) SetStatus(v float32) {
	o.Status = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *MediaRequest) GetMedia() MediaInfo {
	if o == nil || IsNil(o.Media) {
		var ret MediaInfo
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetMediaOk() (*MediaInfo, bool) {
	if o == nil || IsNil(o.Media) {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *MediaRequest) HasMedia() bool {
	if o != nil && !IsNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given MediaInfo and assigns it to the Media field.
func (o *MediaRequest) SetMedia(v MediaInfo) {
	o.Media = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MediaRequest) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MediaRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MediaRequest) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MediaRequest) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MediaRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MediaRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetRequestedBy returns the RequestedBy field value if set, zero value otherwise.
func (o *MediaRequest) GetRequestedBy() User {
	if o == nil || IsNil(o.RequestedBy) {
		var ret User
		return ret
	}
	return *o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetRequestedByOk() (*User, bool) {
	if o == nil || IsNil(o.RequestedBy) {
		return nil, false
	}
	return o.RequestedBy, true
}

// HasRequestedBy returns a boolean if a field has been set.
func (o *MediaRequest) HasRequestedBy() bool {
	if o != nil && !IsNil(o.RequestedBy) {
		return true
	}

	return false
}

// SetRequestedBy gets a reference to the given User and assigns it to the RequestedBy field.
func (o *MediaRequest) SetRequestedBy(v User) {
	o.RequestedBy = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *MediaRequest) GetModifiedBy() MediaRequestModifiedBy {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret MediaRequestModifiedBy
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetModifiedByOk() (*MediaRequestModifiedBy, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *MediaRequest) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given MediaRequestModifiedBy and assigns it to the ModifiedBy field.
func (o *MediaRequest) SetModifiedBy(v MediaRequestModifiedBy) {
	o.ModifiedBy = &v
}

// GetIs4k returns the Is4k field value if set, zero value otherwise.
func (o *MediaRequest) GetIs4k() bool {
	if o == nil || IsNil(o.Is4k) {
		var ret bool
		return ret
	}
	return *o.Is4k
}

// GetIs4kOk returns a tuple with the Is4k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetIs4kOk() (*bool, bool) {
	if o == nil || IsNil(o.Is4k) {
		return nil, false
	}
	return o.Is4k, true
}

// HasIs4k returns a boolean if a field has been set.
func (o *MediaRequest) HasIs4k() bool {
	if o != nil && !IsNil(o.Is4k) {
		return true
	}

	return false
}

// SetIs4k gets a reference to the given bool and assigns it to the Is4k field.
func (o *MediaRequest) SetIs4k(v bool) {
	o.Is4k = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *MediaRequest) GetServerId() float32 {
	if o == nil || IsNil(o.ServerId) {
		var ret float32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetServerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *MediaRequest) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given float32 and assigns it to the ServerId field.
func (o *MediaRequest) SetServerId(v float32) {
	o.ServerId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *MediaRequest) GetProfileId() float32 {
	if o == nil || IsNil(o.ProfileId) {
		var ret float32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *MediaRequest) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given float32 and assigns it to the ProfileId field.
func (o *MediaRequest) SetProfileId(v float32) {
	o.ProfileId = &v
}

// GetRootFolder returns the RootFolder field value if set, zero value otherwise.
func (o *MediaRequest) GetRootFolder() string {
	if o == nil || IsNil(o.RootFolder) {
		var ret string
		return ret
	}
	return *o.RootFolder
}

// GetRootFolderOk returns a tuple with the RootFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaRequest) GetRootFolderOk() (*string, bool) {
	if o == nil || IsNil(o.RootFolder) {
		return nil, false
	}
	return o.RootFolder, true
}

// HasRootFolder returns a boolean if a field has been set.
func (o *MediaRequest) HasRootFolder() bool {
	if o != nil && !IsNil(o.RootFolder) {
		return true
	}

	return false
}

// SetRootFolder gets a reference to the given string and assigns it to the RootFolder field.
func (o *MediaRequest) SetRootFolder(v string) {
	o.RootFolder = &v
}

func (o MediaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.RequestedBy) {
		toSerialize["requestedBy"] = o.RequestedBy
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !IsNil(o.Is4k) {
		toSerialize["is4k"] = o.Is4k
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.RootFolder) {
		toSerialize["rootFolder"] = o.RootFolder
	}
	return toSerialize, nil
}

func (o *MediaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varMediaRequest := _MediaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaRequest)

	if err != nil {
		return err
	}

	*o = MediaRequest(varMediaRequest)

	return err
}

type NullableMediaRequest struct {
	value *MediaRequest
	isSet bool
}

func (v NullableMediaRequest) Get() *MediaRequest {
	return v.value
}

func (v *NullableMediaRequest) Set(val *MediaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaRequest(val *MediaRequest) *NullableMediaRequest {
	return &NullableMediaRequest{value: val, isSet: true}
}

func (v NullableMediaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


