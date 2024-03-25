package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	repo "github.com/dannyh79/commentator/internal/repo"
	routes "github.com/dannyh79/commentator/internal/rest/v1"
	"github.com/dannyh79/commentator/internal/shows"
	utils "github.com/dannyh79/commentator/internal/testutils"
	"github.com/gin-gonic/gin"
)

func Test_POST_Shows_Comments(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		payload        string
		expectedStatus int
		expectedData   string
	}{
		{
			name:           "returns status 201 with created comment",
			path:           "/v1/shows/1/comments",
			payload:        `{"user_id":2,"comment":"some comments"}`,
			expectedStatus: http.StatusCreated,
			expectedData:   `{"result":{"user_id":2,"comment":"some comments","upvote":0}}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			suite := newTestSuite()

			req := httptest.NewRequest(http.MethodPost, tc.path, nil)
			rr := httptest.NewRecorder()

			suite.Engine.ServeHTTP(rr, req)

			utils.AssertJsonHeader(t)(rr)
			utils.AssertHttpStatus(t)(rr, tc.expectedStatus)
			utils.AssertEqual(t)(rr.Body.String(), tc.expectedData)
		})
	}
}

type TestSuite struct {
	Engine *gin.Engine
}

func newTestSuite() *TestSuite {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	r := repo.NewInMemoryCommentRepo()
	usecase := shows.NewShowComments(r)
	routes.AddRoutes(engine, usecase)
	return &TestSuite{engine}
}
