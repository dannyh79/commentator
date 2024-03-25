package testutils_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertEqual(t *testing.T) func(got any, want any) {
	return func(got any, want any) {
		t.Helper()
		if !cmp.Equal(got, want) {
			t.Errorf(cmp.Diff(want, got))
		}
	}
}

func AssertJsonHeader(t *testing.T) func(rr *httptest.ResponseRecorder) {
	return func(rr *httptest.ResponseRecorder) {
		t.Helper()
		want := "application/json"
		if got := rr.Header().Get("Content-Type"); !strings.Contains(got, want) {
			t.Errorf("got HTTP status %v, want %v", got, want)
		}
	}
}

func AssertHttpStatus(t *testing.T) func(rr *httptest.ResponseRecorder, want int) {
	return func(rr *httptest.ResponseRecorder, want int) {
		t.Helper()
		if got := rr.Result().StatusCode; got != want {
			t.Errorf("got HTTP status %v, want %v", got, want)
		}
	}
}
