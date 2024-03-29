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

// checks if the MovieMovieIdRatingsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieMovieIdRatingsGet200Response{}

// MovieMovieIdRatingsGet200Response struct for MovieMovieIdRatingsGet200Response
type MovieMovieIdRatingsGet200Response struct {
	Title *string `json:"title,omitempty"`
	Year *float32 `json:"year,omitempty"`
	Url *string `json:"url,omitempty"`
	CriticsScore *float32 `json:"criticsScore,omitempty"`
	CriticsRating *string `json:"criticsRating,omitempty"`
	AudienceScore *float32 `json:"audienceScore,omitempty"`
	AudienceRating *string `json:"audienceRating,omitempty"`
}

// NewMovieMovieIdRatingsGet200Response instantiates a new MovieMovieIdRatingsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieMovieIdRatingsGet200Response() *MovieMovieIdRatingsGet200Response {
	this := MovieMovieIdRatingsGet200Response{}
	return &this
}

// NewMovieMovieIdRatingsGet200ResponseWithDefaults instantiates a new MovieMovieIdRatingsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieMovieIdRatingsGet200ResponseWithDefaults() *MovieMovieIdRatingsGet200Response {
	this := MovieMovieIdRatingsGet200Response{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MovieMovieIdRatingsGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetYear() float32 {
	if o == nil || IsNil(o.Year) {
		var ret float32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetYearOk() (*float32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given float32 and assigns it to the Year field.
func (o *MovieMovieIdRatingsGet200Response) SetYear(v float32) {
	o.Year = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MovieMovieIdRatingsGet200Response) SetUrl(v string) {
	o.Url = &v
}

// GetCriticsScore returns the CriticsScore field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetCriticsScore() float32 {
	if o == nil || IsNil(o.CriticsScore) {
		var ret float32
		return ret
	}
	return *o.CriticsScore
}

// GetCriticsScoreOk returns a tuple with the CriticsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetCriticsScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.CriticsScore) {
		return nil, false
	}
	return o.CriticsScore, true
}

// HasCriticsScore returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasCriticsScore() bool {
	if o != nil && !IsNil(o.CriticsScore) {
		return true
	}

	return false
}

// SetCriticsScore gets a reference to the given float32 and assigns it to the CriticsScore field.
func (o *MovieMovieIdRatingsGet200Response) SetCriticsScore(v float32) {
	o.CriticsScore = &v
}

// GetCriticsRating returns the CriticsRating field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetCriticsRating() string {
	if o == nil || IsNil(o.CriticsRating) {
		var ret string
		return ret
	}
	return *o.CriticsRating
}

// GetCriticsRatingOk returns a tuple with the CriticsRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetCriticsRatingOk() (*string, bool) {
	if o == nil || IsNil(o.CriticsRating) {
		return nil, false
	}
	return o.CriticsRating, true
}

// HasCriticsRating returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasCriticsRating() bool {
	if o != nil && !IsNil(o.CriticsRating) {
		return true
	}

	return false
}

// SetCriticsRating gets a reference to the given string and assigns it to the CriticsRating field.
func (o *MovieMovieIdRatingsGet200Response) SetCriticsRating(v string) {
	o.CriticsRating = &v
}

// GetAudienceScore returns the AudienceScore field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetAudienceScore() float32 {
	if o == nil || IsNil(o.AudienceScore) {
		var ret float32
		return ret
	}
	return *o.AudienceScore
}

// GetAudienceScoreOk returns a tuple with the AudienceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetAudienceScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.AudienceScore) {
		return nil, false
	}
	return o.AudienceScore, true
}

// HasAudienceScore returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasAudienceScore() bool {
	if o != nil && !IsNil(o.AudienceScore) {
		return true
	}

	return false
}

// SetAudienceScore gets a reference to the given float32 and assigns it to the AudienceScore field.
func (o *MovieMovieIdRatingsGet200Response) SetAudienceScore(v float32) {
	o.AudienceScore = &v
}

// GetAudienceRating returns the AudienceRating field value if set, zero value otherwise.
func (o *MovieMovieIdRatingsGet200Response) GetAudienceRating() string {
	if o == nil || IsNil(o.AudienceRating) {
		var ret string
		return ret
	}
	return *o.AudienceRating
}

// GetAudienceRatingOk returns a tuple with the AudienceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingsGet200Response) GetAudienceRatingOk() (*string, bool) {
	if o == nil || IsNil(o.AudienceRating) {
		return nil, false
	}
	return o.AudienceRating, true
}

// HasAudienceRating returns a boolean if a field has been set.
func (o *MovieMovieIdRatingsGet200Response) HasAudienceRating() bool {
	if o != nil && !IsNil(o.AudienceRating) {
		return true
	}

	return false
}

// SetAudienceRating gets a reference to the given string and assigns it to the AudienceRating field.
func (o *MovieMovieIdRatingsGet200Response) SetAudienceRating(v string) {
	o.AudienceRating = &v
}

func (o MovieMovieIdRatingsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieMovieIdRatingsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CriticsScore) {
		toSerialize["criticsScore"] = o.CriticsScore
	}
	if !IsNil(o.CriticsRating) {
		toSerialize["criticsRating"] = o.CriticsRating
	}
	if !IsNil(o.AudienceScore) {
		toSerialize["audienceScore"] = o.AudienceScore
	}
	if !IsNil(o.AudienceRating) {
		toSerialize["audienceRating"] = o.AudienceRating
	}
	return toSerialize, nil
}

type NullableMovieMovieIdRatingsGet200Response struct {
	value *MovieMovieIdRatingsGet200Response
	isSet bool
}

func (v NullableMovieMovieIdRatingsGet200Response) Get() *MovieMovieIdRatingsGet200Response {
	return v.value
}

func (v *NullableMovieMovieIdRatingsGet200Response) Set(val *MovieMovieIdRatingsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieMovieIdRatingsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieMovieIdRatingsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieMovieIdRatingsGet200Response(val *MovieMovieIdRatingsGet200Response) *NullableMovieMovieIdRatingsGet200Response {
	return &NullableMovieMovieIdRatingsGet200Response{value: val, isSet: true}
}

func (v NullableMovieMovieIdRatingsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieMovieIdRatingsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


