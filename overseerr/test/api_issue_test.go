/*
Overseerr API

Testing IssueAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_IssueAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IssueAPIService IssueCommentCommentIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var commentId string

		httpRes, err := apiClient.IssueAPI.IssueCommentCommentIdDelete(context.Background(), commentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueCommentCommentIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var commentId string

		resp, httpRes, err := apiClient.IssueAPI.IssueCommentCommentIdGet(context.Background(), commentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueCommentCommentIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var commentId string

		resp, httpRes, err := apiClient.IssueAPI.IssueCommentCommentIdPut(context.Background(), commentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IssueAPI.IssueCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IssueAPI.IssueGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueIssueIdCommentPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId float32

		resp, httpRes, err := apiClient.IssueAPI.IssueIssueIdCommentPost(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueIssueIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId string

		httpRes, err := apiClient.IssueAPI.IssueIssueIdDelete(context.Background(), issueId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueIssueIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId float32

		resp, httpRes, err := apiClient.IssueAPI.IssueIssueIdGet(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssueIssueIdStatusPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId string
		var status string

		resp, httpRes, err := apiClient.IssueAPI.IssueIssueIdStatusPost(context.Background(), issueId, status).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssueAPIService IssuePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IssueAPI.IssuePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}