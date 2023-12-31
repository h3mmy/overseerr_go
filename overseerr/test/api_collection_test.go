/*
Overseerr API

Testing CollectionAPIService

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

func Test_overseerr_go_CollectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CollectionAPIService CollectionCollectionIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId float32

		resp, httpRes, err := apiClient.CollectionAPI.CollectionCollectionIdGet(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
