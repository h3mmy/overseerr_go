/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// RequestPostRequestSeasons - struct for RequestPostRequestSeasons
type RequestPostRequestSeasons struct {
	ArrayOfFloat32 *[]float32
	String *string
}

// []float32AsRequestPostRequestSeasons is a convenience function that returns []float32 wrapped in RequestPostRequestSeasons
func ArrayOfFloat32AsRequestPostRequestSeasons(v *[]float32) RequestPostRequestSeasons {
	return RequestPostRequestSeasons{
		ArrayOfFloat32: v,
	}
}

// stringAsRequestPostRequestSeasons is a convenience function that returns string wrapped in RequestPostRequestSeasons
func StringAsRequestPostRequestSeasons(v *string) RequestPostRequestSeasons {
	return RequestPostRequestSeasons{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestPostRequestSeasons) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfFloat32
	err = json.Unmarshal(data, &dst.ArrayOfFloat32)
	if err == nil {
		v, _ := json.Marshal(dst.ArrayOfFloat32)
		if string(v) == "{}" { // empty struct
			dst.ArrayOfFloat32 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfFloat32 = nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.String)
		if string(jsonstring) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfFloat32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RequestPostRequestSeasons)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestPostRequestSeasons)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestPostRequestSeasons) MarshalJSON() ([]byte, error) {
	if src.ArrayOfFloat32 != nil {
		return json.Marshal(&src.ArrayOfFloat32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestPostRequestSeasons) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfFloat32 != nil {
		return obj.ArrayOfFloat32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRequestPostRequestSeasons struct {
	value *RequestPostRequestSeasons
	isSet bool
}

func (v NullableRequestPostRequestSeasons) Get() *RequestPostRequestSeasons {
	return v.value
}

func (v *NullableRequestPostRequestSeasons) Set(val *RequestPostRequestSeasons) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPostRequestSeasons) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPostRequestSeasons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPostRequestSeasons(val *RequestPostRequestSeasons) *NullableRequestPostRequestSeasons {
	return &NullableRequestPostRequestSeasons{value: val, isSet: true}
}

func (v NullableRequestPostRequestSeasons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPostRequestSeasons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
