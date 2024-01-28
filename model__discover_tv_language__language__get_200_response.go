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

// checks if the DiscoverTvLanguageLanguageGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoverTvLanguageLanguageGet200Response{}

// DiscoverTvLanguageLanguageGet200Response struct for DiscoverTvLanguageLanguageGet200Response
type DiscoverTvLanguageLanguageGet200Response struct {
	Page *float32 `json:"page,omitempty"`
	TotalPages *float32 `json:"totalPages,omitempty"`
	TotalResults *float32 `json:"totalResults,omitempty"`
	Language *SpokenLanguage `json:"language,omitempty"`
	Results []TvResult `json:"results,omitempty"`
}

// NewDiscoverTvLanguageLanguageGet200Response instantiates a new DiscoverTvLanguageLanguageGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoverTvLanguageLanguageGet200Response() *DiscoverTvLanguageLanguageGet200Response {
	this := DiscoverTvLanguageLanguageGet200Response{}
	return &this
}

// NewDiscoverTvLanguageLanguageGet200ResponseWithDefaults instantiates a new DiscoverTvLanguageLanguageGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoverTvLanguageLanguageGet200ResponseWithDefaults() *DiscoverTvLanguageLanguageGet200Response {
	this := DiscoverTvLanguageLanguageGet200Response{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DiscoverTvLanguageLanguageGet200Response) GetPage() float32 {
	if o == nil || IsNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) GetPageOk() (*float32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *DiscoverTvLanguageLanguageGet200Response) SetPage(v float32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *DiscoverTvLanguageLanguageGet200Response) GetTotalPages() float32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret float32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) GetTotalPagesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given float32 and assigns it to the TotalPages field.
func (o *DiscoverTvLanguageLanguageGet200Response) SetTotalPages(v float32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *DiscoverTvLanguageLanguageGet200Response) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *DiscoverTvLanguageLanguageGet200Response) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *DiscoverTvLanguageLanguageGet200Response) GetLanguage() SpokenLanguage {
	if o == nil || IsNil(o.Language) {
		var ret SpokenLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) GetLanguageOk() (*SpokenLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given SpokenLanguage and assigns it to the Language field.
func (o *DiscoverTvLanguageLanguageGet200Response) SetLanguage(v SpokenLanguage) {
	o.Language = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DiscoverTvLanguageLanguageGet200Response) GetResults() []TvResult {
	if o == nil || IsNil(o.Results) {
		var ret []TvResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) GetResultsOk() ([]TvResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DiscoverTvLanguageLanguageGet200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TvResult and assigns it to the Results field.
func (o *DiscoverTvLanguageLanguageGet200Response) SetResults(v []TvResult) {
	o.Results = v
}

func (o DiscoverTvLanguageLanguageGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoverTvLanguageLanguageGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDiscoverTvLanguageLanguageGet200Response struct {
	value *DiscoverTvLanguageLanguageGet200Response
	isSet bool
}

func (v NullableDiscoverTvLanguageLanguageGet200Response) Get() *DiscoverTvLanguageLanguageGet200Response {
	return v.value
}

func (v *NullableDiscoverTvLanguageLanguageGet200Response) Set(val *DiscoverTvLanguageLanguageGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoverTvLanguageLanguageGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoverTvLanguageLanguageGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoverTvLanguageLanguageGet200Response(val *DiscoverTvLanguageLanguageGet200Response) *NullableDiscoverTvLanguageLanguageGet200Response {
	return &NullableDiscoverTvLanguageLanguageGet200Response{value: val, isSet: true}
}

func (v NullableDiscoverTvLanguageLanguageGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoverTvLanguageLanguageGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


