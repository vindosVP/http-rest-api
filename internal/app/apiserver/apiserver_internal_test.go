package apiserver

import (
	"github.com/stretchr/testify/assert"
	"github.com/vindosVp/http-rest-api/internal/app/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_HandleHeartbeat(t *testing.T) {
	conf, err := config.NewConfig("../../../configs/apiserver_test.yaml")

	if err != nil {
		t.Fatal(err)
	}

	s := New(conf)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
	s.handleHeartbeat().ServeHTTP(rec, req)
	assert.NotEmpty(t, rec.Body.String(), nil)
}
