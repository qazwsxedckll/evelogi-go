package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	request := httptest.NewRequest("GET", "/ping", nil)
	response := httptest.NewRecorder()

	s := NewServer()
	s.ServeHTTP(response, request)

	require.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", response.Body.String())
}
