package app_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestRootHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)
	assert.Equal(t, response.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, string(body), "API")
}
func TestHelloHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello", nil)
	response := executeRequest(req)
	assert.Equal(t, response.Code, http.StatusOK)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	assert.Equal(t, m["greeting"], "Hello")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	testApp.Router.ServeHTTP(rr, req)

	return rr
}
