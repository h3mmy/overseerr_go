/*
Overseerr API

Testing ServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package overseerr_go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/h3mmy/overseerr_go"
)

func Test_overseerr_go_ServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceAPIService ServiceRadarrGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceAPI.ServiceRadarrGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ServiceRadarrRadarrIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var radarrId float32

		resp, httpRes, err := apiClient.ServiceAPI.ServiceRadarrRadarrIdGet(context.Background(), radarrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ServiceSonarrGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceAPI.ServiceSonarrGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ServiceSonarrLookupTmdbIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tmdbId float32

		resp, httpRes, err := apiClient.ServiceAPI.ServiceSonarrLookupTmdbIdGet(context.Background(), tmdbId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ServiceSonarrSonarrIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sonarrId float32

		resp, httpRes, err := apiClient.ServiceAPI.ServiceSonarrSonarrIdGet(context.Background(), sonarrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
