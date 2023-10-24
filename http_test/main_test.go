package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	r := httptest.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, r)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, []byte("OK\n"), data)
}
