/*
Overseerr API

Testing TvAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package overseerr_go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_overseerr_go_TvAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TvAPIService TvTvIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tvId float32

		resp, httpRes, err := apiClient.TvAPI.TvTvIdGet(context.Background(), tvId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TvAPIService TvTvIdRatingsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tvId float32

		resp, httpRes, err := apiClient.TvAPI.TvTvIdRatingsGet(context.Background(), tvId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TvAPIService TvTvIdRecommendationsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tvId float32

		resp, httpRes, err := apiClient.TvAPI.TvTvIdRecommendationsGet(context.Background(), tvId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TvAPIService TvTvIdSeasonSeasonIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tvId float32
		var seasonId float32

		resp, httpRes, err := apiClient.TvAPI.TvTvIdSeasonSeasonIdGet(context.Background(), tvId, seasonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TvAPIService TvTvIdSimilarGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tvId float32

		resp, httpRes, err := apiClient.TvAPI.TvTvIdSimilarGet(context.Background(), tvId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
