package main

import (
	"net/http"
	"testing"

	"snippetbox.reandov.com/internal/assert"
)

func TestHealth(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/health")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
