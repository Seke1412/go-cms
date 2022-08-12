package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSamples(t *testing.T) {
	router := setupRouter()
	request, _ := http.NewRequest("GET", "http://localhost:8080/samples", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("cache-control", "no-cache")

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "Response code should be 200")
	assert.NotEmpty(t, response.Body, "tokenString response should not be empty")

}
