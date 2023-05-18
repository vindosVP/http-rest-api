package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_HandleHeartbeat(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
	s.handleHeartbeat().ServeHTTP(rec, req)
	assert.NotEmpty(t, rec.Body.String(), nil)
}
