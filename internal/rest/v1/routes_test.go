package routes_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	routes "github.com/dannyh79/commentator/internal/rest/v1"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
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

			assertJsonHeader(t)(rr)
			assertHttpStatus(t)(rr, tc.expectedStatus)
			assertEqual(t)(rr.Body.String(), tc.expectedData)
		})
	}
}

type TestSuite struct {
	Engine *gin.Engine
}

func newTestSuite() *TestSuite {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	routes.AddRoutes(engine)
	return &TestSuite{engine}
}

func assertEqual(t *testing.T) func(got any, want any) {
	return func(got any, want any) {
		t.Helper()
		if !cmp.Equal(got, want) {
			t.Errorf(cmp.Diff(want, got))
		}
	}
}

func assertJsonHeader(t *testing.T) func(rr *httptest.ResponseRecorder) {
	return func(rr *httptest.ResponseRecorder) {
		t.Helper()
		want := "application/json"
		if got := rr.Header().Get("Content-Type"); !strings.Contains(got, want) {
			t.Errorf("got HTTP status %v, want %v", got, want)
		}
	}
}

func assertHttpStatus(t *testing.T) func(rr *httptest.ResponseRecorder, want int) {
	return func(rr *httptest.ResponseRecorder, want int) {
		t.Helper()
		if got := rr.Result().StatusCode; got != want {
			t.Errorf("got HTTP status %v, want %v", got, want)
		}
	}
}
