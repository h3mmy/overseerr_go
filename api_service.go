/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type ServiceAPI interface {

	/*
	ServiceRadarrGet Get non-sensitive Radarr server list

	Returns a list of Radarr server IDs and names in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ServiceAPIServiceRadarrGetRequest
	*/
	ServiceRadarrGet(ctx context.Context) ServiceAPIServiceRadarrGetRequest

	// ServiceRadarrGetExecute executes the request
	//  @return []RadarrSettings
	ServiceRadarrGetExecute(r ServiceAPIServiceRadarrGetRequest) ([]RadarrSettings, *http.Response, error)

	/*
	ServiceRadarrRadarrIdGet Get Radarr server quality profiles and root folders

	Returns a Radarr server's quality profile and root folder details in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param radarrId
	@return ServiceAPIServiceRadarrRadarrIdGetRequest
	*/
	ServiceRadarrRadarrIdGet(ctx context.Context, radarrId float32) ServiceAPIServiceRadarrRadarrIdGetRequest

	// ServiceRadarrRadarrIdGetExecute executes the request
	//  @return ServiceRadarrRadarrIdGet200Response
	ServiceRadarrRadarrIdGetExecute(r ServiceAPIServiceRadarrRadarrIdGetRequest) (*ServiceRadarrRadarrIdGet200Response, *http.Response, error)

	/*
	ServiceSonarrGet Get non-sensitive Sonarr server list

	Returns a list of Sonarr server IDs and names in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ServiceAPIServiceSonarrGetRequest
	*/
	ServiceSonarrGet(ctx context.Context) ServiceAPIServiceSonarrGetRequest

	// ServiceSonarrGetExecute executes the request
	//  @return []SonarrSettings
	ServiceSonarrGetExecute(r ServiceAPIServiceSonarrGetRequest) ([]SonarrSettings, *http.Response, error)

	/*
	ServiceSonarrLookupTmdbIdGet Get series from Sonarr

	Returns a list of series returned by searching for the name in Sonarr.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tmdbId
	@return ServiceAPIServiceSonarrLookupTmdbIdGetRequest
	*/
	ServiceSonarrLookupTmdbIdGet(ctx context.Context, tmdbId float32) ServiceAPIServiceSonarrLookupTmdbIdGetRequest

	// ServiceSonarrLookupTmdbIdGetExecute executes the request
	//  @return []SonarrSeries
	ServiceSonarrLookupTmdbIdGetExecute(r ServiceAPIServiceSonarrLookupTmdbIdGetRequest) ([]SonarrSeries, *http.Response, error)

	/*
	ServiceSonarrSonarrIdGet Get Sonarr server quality profiles and root folders

	Returns a Sonarr server's quality profile and root folder details in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sonarrId
	@return ServiceAPIServiceSonarrSonarrIdGetRequest
	*/
	ServiceSonarrSonarrIdGet(ctx context.Context, sonarrId float32) ServiceAPIServiceSonarrSonarrIdGetRequest

	// ServiceSonarrSonarrIdGetExecute executes the request
	//  @return ServiceSonarrSonarrIdGet200Response
	ServiceSonarrSonarrIdGetExecute(r ServiceAPIServiceSonarrSonarrIdGetRequest) (*ServiceSonarrSonarrIdGet200Response, *http.Response, error)
}

// ServiceAPIService ServiceAPI service
type ServiceAPIService service

type ServiceAPIServiceRadarrGetRequest struct {
	ctx context.Context
	ApiService ServiceAPI
}

func (r ServiceAPIServiceRadarrGetRequest) Execute() ([]RadarrSettings, *http.Response, error) {
	return r.ApiService.ServiceRadarrGetExecute(r)
}

/*
ServiceRadarrGet Get non-sensitive Radarr server list

Returns a list of Radarr server IDs and names in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ServiceAPIServiceRadarrGetRequest
*/
func (a *ServiceAPIService) ServiceRadarrGet(ctx context.Context) ServiceAPIServiceRadarrGetRequest {
	return ServiceAPIServiceRadarrGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RadarrSettings
func (a *ServiceAPIService) ServiceRadarrGetExecute(r ServiceAPIServiceRadarrGetRequest) ([]RadarrSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RadarrSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAPIService.ServiceRadarrGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/radarr"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ServiceAPIServiceRadarrRadarrIdGetRequest struct {
	ctx context.Context
	ApiService ServiceAPI
	radarrId float32
}

func (r ServiceAPIServiceRadarrRadarrIdGetRequest) Execute() (*ServiceRadarrRadarrIdGet200Response, *http.Response, error) {
	return r.ApiService.ServiceRadarrRadarrIdGetExecute(r)
}

/*
ServiceRadarrRadarrIdGet Get Radarr server quality profiles and root folders

Returns a Radarr server's quality profile and root folder details in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param radarrId
 @return ServiceAPIServiceRadarrRadarrIdGetRequest
*/
func (a *ServiceAPIService) ServiceRadarrRadarrIdGet(ctx context.Context, radarrId float32) ServiceAPIServiceRadarrRadarrIdGetRequest {
	return ServiceAPIServiceRadarrRadarrIdGetRequest{
		ApiService: a,
		ctx: ctx,
		radarrId: radarrId,
	}
}

// Execute executes the request
//  @return ServiceRadarrRadarrIdGet200Response
func (a *ServiceAPIService) ServiceRadarrRadarrIdGetExecute(r ServiceAPIServiceRadarrRadarrIdGetRequest) (*ServiceRadarrRadarrIdGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceRadarrRadarrIdGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAPIService.ServiceRadarrRadarrIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/radarr/{radarrId}"
	localVarPath = strings.Replace(localVarPath, "{"+"radarrId"+"}", url.PathEscape(parameterValueToString(r.radarrId, "radarrId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ServiceAPIServiceSonarrGetRequest struct {
	ctx context.Context
	ApiService ServiceAPI
}

func (r ServiceAPIServiceSonarrGetRequest) Execute() ([]SonarrSettings, *http.Response, error) {
	return r.ApiService.ServiceSonarrGetExecute(r)
}

/*
ServiceSonarrGet Get non-sensitive Sonarr server list

Returns a list of Sonarr server IDs and names in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ServiceAPIServiceSonarrGetRequest
*/
func (a *ServiceAPIService) ServiceSonarrGet(ctx context.Context) ServiceAPIServiceSonarrGetRequest {
	return ServiceAPIServiceSonarrGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SonarrSettings
func (a *ServiceAPIService) ServiceSonarrGetExecute(r ServiceAPIServiceSonarrGetRequest) ([]SonarrSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SonarrSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAPIService.ServiceSonarrGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/sonarr"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ServiceAPIServiceSonarrLookupTmdbIdGetRequest struct {
	ctx context.Context
	ApiService ServiceAPI
	tmdbId float32
}

func (r ServiceAPIServiceSonarrLookupTmdbIdGetRequest) Execute() ([]SonarrSeries, *http.Response, error) {
	return r.ApiService.ServiceSonarrLookupTmdbIdGetExecute(r)
}

/*
ServiceSonarrLookupTmdbIdGet Get series from Sonarr

Returns a list of series returned by searching for the name in Sonarr.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tmdbId
 @return ServiceAPIServiceSonarrLookupTmdbIdGetRequest
*/
func (a *ServiceAPIService) ServiceSonarrLookupTmdbIdGet(ctx context.Context, tmdbId float32) ServiceAPIServiceSonarrLookupTmdbIdGetRequest {
	return ServiceAPIServiceSonarrLookupTmdbIdGetRequest{
		ApiService: a,
		ctx: ctx,
		tmdbId: tmdbId,
	}
}

// Execute executes the request
//  @return []SonarrSeries
func (a *ServiceAPIService) ServiceSonarrLookupTmdbIdGetExecute(r ServiceAPIServiceSonarrLookupTmdbIdGetRequest) ([]SonarrSeries, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SonarrSeries
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAPIService.ServiceSonarrLookupTmdbIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/sonarr/lookup/{tmdbId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tmdbId"+"}", url.PathEscape(parameterValueToString(r.tmdbId, "tmdbId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ServiceAPIServiceSonarrSonarrIdGetRequest struct {
	ctx context.Context
	ApiService ServiceAPI
	sonarrId float32
}

func (r ServiceAPIServiceSonarrSonarrIdGetRequest) Execute() (*ServiceSonarrSonarrIdGet200Response, *http.Response, error) {
	return r.ApiService.ServiceSonarrSonarrIdGetExecute(r)
}

/*
ServiceSonarrSonarrIdGet Get Sonarr server quality profiles and root folders

Returns a Sonarr server's quality profile and root folder details in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sonarrId
 @return ServiceAPIServiceSonarrSonarrIdGetRequest
*/
func (a *ServiceAPIService) ServiceSonarrSonarrIdGet(ctx context.Context, sonarrId float32) ServiceAPIServiceSonarrSonarrIdGetRequest {
	return ServiceAPIServiceSonarrSonarrIdGetRequest{
		ApiService: a,
		ctx: ctx,
		sonarrId: sonarrId,
	}
}

// Execute executes the request
//  @return ServiceSonarrSonarrIdGet200Response
func (a *ServiceAPIService) ServiceSonarrSonarrIdGetExecute(r ServiceAPIServiceSonarrSonarrIdGetRequest) (*ServiceSonarrSonarrIdGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceSonarrSonarrIdGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAPIService.ServiceSonarrSonarrIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/sonarr/{sonarrId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sonarrId"+"}", url.PathEscape(parameterValueToString(r.sonarrId, "sonarrId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
