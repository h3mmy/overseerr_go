# \UsersAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthMeGet**](UsersAPI.md#AuthMeGet) | **Get** /auth/me | Get logged-in user
[**AuthResetPasswordGuidPost**](UsersAPI.md#AuthResetPasswordGuidPost) | **Post** /auth/reset-password/{guid} | Reset the password for a user
[**AuthResetPasswordPost**](UsersAPI.md#AuthResetPasswordPost) | **Post** /auth/reset-password | Send a reset password email
[**SettingsPlexUsersGet**](UsersAPI.md#SettingsPlexUsersGet) | **Get** /settings/plex/users | Get Plex users
[**UserGet**](UsersAPI.md#UserGet) | **Get** /user | Get all users
[**UserImportFromPlexPost**](UsersAPI.md#UserImportFromPlexPost) | **Post** /user/import-from-plex | Import all users from Plex
[**UserPost**](UsersAPI.md#UserPost) | **Post** /user | Create new user
[**UserPut**](UsersAPI.md#UserPut) | **Put** /user | Update batch of users
[**UserRegisterPushSubscriptionPost**](UsersAPI.md#UserRegisterPushSubscriptionPost) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
[**UserUserIdDelete**](UsersAPI.md#UserUserIdDelete) | **Delete** /user/{userId} | Delete user by ID
[**UserUserIdGet**](UsersAPI.md#UserUserIdGet) | **Get** /user/{userId} | Get user by ID
[**UserUserIdPut**](UsersAPI.md#UserUserIdPut) | **Put** /user/{userId} | Update a user by user ID
[**UserUserIdQuotaGet**](UsersAPI.md#UserUserIdQuotaGet) | **Get** /user/{userId}/quota | Get quotas for a specific user
[**UserUserIdRequestsGet**](UsersAPI.md#UserUserIdRequestsGet) | **Get** /user/{userId}/requests | Get requests for a specific user
[**UserUserIdSettingsMainGet**](UsersAPI.md#UserUserIdSettingsMainGet) | **Get** /user/{userId}/settings/main | Get general settings for a user
[**UserUserIdSettingsMainPost**](UsersAPI.md#UserUserIdSettingsMainPost) | **Post** /user/{userId}/settings/main | Update general settings for a user
[**UserUserIdSettingsNotificationsGet**](UsersAPI.md#UserUserIdSettingsNotificationsGet) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
[**UserUserIdSettingsNotificationsPost**](UsersAPI.md#UserUserIdSettingsNotificationsPost) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
[**UserUserIdSettingsPasswordGet**](UsersAPI.md#UserUserIdSettingsPasswordGet) | **Get** /user/{userId}/settings/password | Get password page informatiom
[**UserUserIdSettingsPasswordPost**](UsersAPI.md#UserUserIdSettingsPasswordPost) | **Post** /user/{userId}/settings/password | Update password for a user
[**UserUserIdSettingsPermissionsGet**](UsersAPI.md#UserUserIdSettingsPermissionsGet) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
[**UserUserIdSettingsPermissionsPost**](UsersAPI.md#UserUserIdSettingsPermissionsPost) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
[**UserUserIdWatchDataGet**](UsersAPI.md#UserUserIdWatchDataGet) | **Get** /user/{userId}/watch_data | Get watch data
[**UserUserIdWatchlistGet**](UsersAPI.md#UserUserIdWatchlistGet) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user



## AuthMeGet

> User AuthMeGet(ctx).Execute()

Get logged-in user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AuthMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AuthMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthMeGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AuthMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthMeGetRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthResetPasswordGuidPost

> AuthLogoutPost200Response AuthResetPasswordGuidPost(ctx, guid).AuthResetPasswordGuidPostRequest(authResetPasswordGuidPostRequest).Execute()

Reset the password for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	guid := "9afef5a7-ec89-4d5f-9397-261e96970b50" // string | 
	authResetPasswordGuidPostRequest := *openapiclient.NewAuthResetPasswordGuidPostRequest("Password_example") // AuthResetPasswordGuidPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AuthResetPasswordGuidPost(context.Background(), guid).AuthResetPasswordGuidPostRequest(authResetPasswordGuidPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AuthResetPasswordGuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthResetPasswordGuidPost`: AuthLogoutPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AuthResetPasswordGuidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthResetPasswordGuidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authResetPasswordGuidPostRequest** | [**AuthResetPasswordGuidPostRequest**](AuthResetPasswordGuidPostRequest.md) |  | 

### Return type

[**AuthLogoutPost200Response**](AuthLogoutPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthResetPasswordPost

> AuthLogoutPost200Response AuthResetPasswordPost(ctx).AuthResetPasswordPostRequest(authResetPasswordPostRequest).Execute()

Send a reset password email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	authResetPasswordPostRequest := *openapiclient.NewAuthResetPasswordPostRequest("Email_example") // AuthResetPasswordPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AuthResetPasswordPost(context.Background()).AuthResetPasswordPostRequest(authResetPasswordPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AuthResetPasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthResetPasswordPost`: AuthLogoutPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AuthResetPasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthResetPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authResetPasswordPostRequest** | [**AuthResetPasswordPostRequest**](AuthResetPasswordPostRequest.md) |  | 

### Return type

[**AuthLogoutPost200Response**](AuthLogoutPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsPlexUsersGet

> []SettingsPlexUsersGet200ResponseInner SettingsPlexUsersGet(ctx).Execute()

Get Plex users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SettingsPlexUsersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SettingsPlexUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexUsersGet`: []SettingsPlexUsersGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SettingsPlexUsersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexUsersGetRequest struct via the builder pattern


### Return type

[**[]SettingsPlexUsersGet200ResponseInner**](SettingsPlexUsersGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> UserGet200Response UserGet(ctx).Take(take).Skip(skip).Sort(sort).Execute()

Get all users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	take := float32(20) // float32 |  (optional)
	skip := float32(0) // float32 |  (optional)
	sort := "sort_example" // string |  (optional) (default to "created")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserGet(context.Background()).Take(take).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGet`: UserGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **sort** | **string** |  | [default to &quot;created&quot;]

### Return type

[**UserGet200Response**](UserGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserImportFromPlexPost

> []User UserImportFromPlexPost(ctx).UserImportFromPlexPostRequest(userImportFromPlexPostRequest).Execute()

Import all users from Plex



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userImportFromPlexPostRequest := *openapiclient.NewUserImportFromPlexPostRequest() // UserImportFromPlexPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserImportFromPlexPost(context.Background()).UserImportFromPlexPostRequest(userImportFromPlexPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserImportFromPlexPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserImportFromPlexPost`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserImportFromPlexPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserImportFromPlexPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userImportFromPlexPostRequest** | [**UserImportFromPlexPostRequest**](UserImportFromPlexPostRequest.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPost

> User UserPost(ctx).UserPostRequest(userPostRequest).Execute()

Create new user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userPostRequest := *openapiclient.NewUserPostRequest() // UserPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserPost(context.Background()).UserPostRequest(userPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPost`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPostRequest** | [**UserPostRequest**](UserPostRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPut

> []User UserPut(ctx).UserPutRequest(userPutRequest).Execute()

Update batch of users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userPutRequest := *openapiclient.NewUserPutRequest() // UserPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserPut(context.Background()).UserPutRequest(userPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPut`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPutRequest** | [**UserPutRequest**](UserPutRequest.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRegisterPushSubscriptionPost

> UserRegisterPushSubscriptionPost(ctx).UserRegisterPushSubscriptionPostRequest(userRegisterPushSubscriptionPostRequest).Execute()

Register a web push /user/registerPushSubscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userRegisterPushSubscriptionPostRequest := *openapiclient.NewUserRegisterPushSubscriptionPostRequest("Endpoint_example", "Auth_example", "P256dh_example") // UserRegisterPushSubscriptionPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UserRegisterPushSubscriptionPost(context.Background()).UserRegisterPushSubscriptionPostRequest(userRegisterPushSubscriptionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserRegisterPushSubscriptionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserRegisterPushSubscriptionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRegisterPushSubscriptionPostRequest** | [**UserRegisterPushSubscriptionPostRequest**](UserRegisterPushSubscriptionPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdDelete

> User UserUserIdDelete(ctx, userId).Execute()

Delete user by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdDelete(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdDelete`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdGet

> User UserUserIdGet(ctx, userId).Execute()

Get user by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdPut

> User UserUserIdPut(ctx, userId).User(user).Execute()

Update a user by user ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	user := *openapiclient.NewUser(int32(1), "hey@itsme.com", "2020-09-02T05:02:23.000Z", "2020-09-02T05:02:23.000Z") // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdPut(context.Background(), userId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdPut`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdQuotaGet

> UserUserIdQuotaGet200Response UserUserIdQuotaGet(ctx, userId).Execute()

Get quotas for a specific user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdQuotaGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdQuotaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdQuotaGet`: UserUserIdQuotaGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdQuotaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserUserIdQuotaGet200Response**](UserUserIdQuotaGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdRequestsGet

> UserUserIdRequestsGet200Response UserUserIdRequestsGet(ctx, userId).Take(take).Skip(skip).Execute()

Get requests for a specific user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	take := float32(20) // float32 |  (optional)
	skip := float32(0) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdRequestsGet(context.Background(), userId).Take(take).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdRequestsGet`: UserUserIdRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **take** | **float32** |  | 
 **skip** | **float32** |  | 

### Return type

[**UserUserIdRequestsGet200Response**](UserUserIdRequestsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsMainGet

> UserUserIdSettingsMainGet200Response UserUserIdSettingsMainGet(ctx, userId).Execute()

Get general settings for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsMainGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsMainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsMainGet`: UserUserIdSettingsMainGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsMainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsMainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserUserIdSettingsMainGet200Response**](UserUserIdSettingsMainGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsMainPost

> UserUserIdSettingsMainGet200Response UserUserIdSettingsMainPost(ctx, userId).UserUserIdSettingsMainPostRequest(userUserIdSettingsMainPostRequest).Execute()

Update general settings for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	userUserIdSettingsMainPostRequest := *openapiclient.NewUserUserIdSettingsMainPostRequest() // UserUserIdSettingsMainPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsMainPost(context.Background(), userId).UserUserIdSettingsMainPostRequest(userUserIdSettingsMainPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsMainPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsMainPost`: UserUserIdSettingsMainGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsMainPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsMainPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUserIdSettingsMainPostRequest** | [**UserUserIdSettingsMainPostRequest**](UserUserIdSettingsMainPostRequest.md) |  | 

### Return type

[**UserUserIdSettingsMainGet200Response**](UserUserIdSettingsMainGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsNotificationsGet

> UserSettingsNotifications UserUserIdSettingsNotificationsGet(ctx, userId).Execute()

Get notification settings for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsNotificationsGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsNotificationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsNotificationsGet`: UserSettingsNotifications
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsNotificationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsNotificationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSettingsNotifications**](UserSettingsNotifications.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsNotificationsPost

> UserSettingsNotifications UserUserIdSettingsNotificationsPost(ctx, userId).UserSettingsNotifications(userSettingsNotifications).Execute()

Update notification settings for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	userSettingsNotifications := *openapiclient.NewUserSettingsNotifications() // UserSettingsNotifications | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsNotificationsPost(context.Background(), userId).UserSettingsNotifications(userSettingsNotifications).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsNotificationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsNotificationsPost`: UserSettingsNotifications
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsNotificationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsNotificationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSettingsNotifications** | [**UserSettingsNotifications**](UserSettingsNotifications.md) |  | 

### Return type

[**UserSettingsNotifications**](UserSettingsNotifications.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsPasswordGet

> UserUserIdSettingsPasswordGet200Response UserUserIdSettingsPasswordGet(ctx, userId).Execute()

Get password page informatiom



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsPasswordGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsPasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsPasswordGet`: UserUserIdSettingsPasswordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsPasswordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsPasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserUserIdSettingsPasswordGet200Response**](UserUserIdSettingsPasswordGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsPasswordPost

> UserUserIdSettingsPasswordPost(ctx, userId).UserUserIdSettingsPasswordPostRequest(userUserIdSettingsPasswordPostRequest).Execute()

Update password for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	userUserIdSettingsPasswordPostRequest := *openapiclient.NewUserUserIdSettingsPasswordPostRequest("NewPassword_example") // UserUserIdSettingsPasswordPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UserUserIdSettingsPasswordPost(context.Background(), userId).UserUserIdSettingsPasswordPostRequest(userUserIdSettingsPasswordPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsPasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUserIdSettingsPasswordPostRequest** | [**UserUserIdSettingsPasswordPostRequest**](UserUserIdSettingsPasswordPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsPermissionsGet

> UserUserIdSettingsPermissionsGet200Response UserUserIdSettingsPermissionsGet(ctx, userId).Execute()

Get permission settings for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsPermissionsGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsPermissionsGet`: UserUserIdSettingsPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserUserIdSettingsPermissionsGet200Response**](UserUserIdSettingsPermissionsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdSettingsPermissionsPost

> UserUserIdSettingsPermissionsGet200Response UserUserIdSettingsPermissionsPost(ctx, userId).UserUserIdSettingsPermissionsPostRequest(userUserIdSettingsPermissionsPostRequest).Execute()

Update permission settings for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	userUserIdSettingsPermissionsPostRequest := *openapiclient.NewUserUserIdSettingsPermissionsPostRequest(float32(123)) // UserUserIdSettingsPermissionsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdSettingsPermissionsPost(context.Background(), userId).UserUserIdSettingsPermissionsPostRequest(userUserIdSettingsPermissionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdSettingsPermissionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdSettingsPermissionsPost`: UserUserIdSettingsPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdSettingsPermissionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdSettingsPermissionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUserIdSettingsPermissionsPostRequest** | [**UserUserIdSettingsPermissionsPostRequest**](UserUserIdSettingsPermissionsPostRequest.md) |  | 

### Return type

[**UserUserIdSettingsPermissionsGet200Response**](UserUserIdSettingsPermissionsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdWatchDataGet

> UserUserIdWatchDataGet200Response UserUserIdWatchDataGet(ctx, userId).Execute()

Get watch data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdWatchDataGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdWatchDataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdWatchDataGet`: UserUserIdWatchDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdWatchDataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdWatchDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserUserIdWatchDataGet200Response**](UserUserIdWatchDataGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUserIdWatchlistGet

> UserUserIdWatchlistGet200Response UserUserIdWatchlistGet(ctx, userId).Page(page).Execute()

Get the Plex watchlist for a specific user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func main() {
	userId := float32(8.14) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUserIdWatchlistGet(context.Background(), userId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUserIdWatchlistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUserIdWatchlistGet`: UserUserIdWatchlistGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUserIdWatchlistGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserIdWatchlistGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]

### Return type

[**UserUserIdWatchlistGet200Response**](UserUserIdWatchlistGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

