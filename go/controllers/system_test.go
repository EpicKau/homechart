package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/candiddev/homechart/go/models"
	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestSystemInfoRead(t *testing.T) {
	logger.UseTestLogger(t)

	var i []Info

	r := request{
		method:       "GET",
		responseType: &i,
		uri:          "/api",
	}

	msg := r.do()

	assert.Equal(t, msg.Status, 200)
	assert.Equal(t, i[0].Version, info.Version)
}

func TestSystem(t *testing.T) {
	logger.UseTestLogger(t)

	var sc systemClient

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/v1/system/client?key=client", nil)
	r, _ := client.Do(req)
	json.NewDecoder(r.Body).Decode(&sc)

	assert.Equal(t, r.StatusCode, 200)
	assert.Equal(t, len(sc.Headers), 2)
	assert.Contains(t, sc.RemoteAddr, "127.0.0.1")

	var h models.Health

	client = &http.Client{}
	req, _ = http.NewRequest(http.MethodGet, ts.URL+"/api/v1/system/health?key=health", nil)
	r, _ = client.Do(req)
	json.NewDecoder(r.Body).Decode(&h)

	assert.Equal(t, r.StatusCode, 200)
	assert.Equal(t, h.DB, true)

	req, _ = http.NewRequest(http.MethodGet, ts.URL+"/api/v1/system/health", nil)
	r, _ = client.Do(req)

	assert.Equal(t, r.StatusCode, 403)

	req, _ = http.NewRequest(http.MethodGet, ts.URL+"/api/v1/system/metrics", nil)
	r, _ = client.Do(req)

	assert.Equal(t, r.StatusCode, 403)

	req, _ = http.NewRequest(http.MethodGet, ts.URL+"/api/v1/system/pprof/goroutines", nil)
	r, _ = client.Do(req)

	assert.Equal(t, r.StatusCode, 403)

	req, _ = http.NewRequest(http.MethodGet, ts.URL+"/api/v1/system/stop", nil)
	r, _ = client.Do(req)

	assert.Equal(t, r.StatusCode, 403)
}
