package gobase64

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoBase64_HandleEncode(t *testing.T) {
	s := New(NewConfig())

	testCases := []struct {
		payload url.Values
		code    int
	}{
		{payload: url.Values{"text": {"And"}}, code: http.StatusOK},
		{payload: url.Values{"text": {"Привет"}}, code: http.StatusOK},

		{payload: url.Values{"test": {"And"}}, code: http.StatusBadRequest},
		{payload: url.Values{}, code: http.StatusBadRequest},
	}

	for _, cs := range testCases {
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/encode", nil)
		req.URL.RawQuery = cs.payload.Encode()

		s.HandlerEncode().ServeHTTP(rec, req)

		assert.Equal(t, cs.code, rec.Code)
	}
}
