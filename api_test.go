package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {
	response := callApi(http.MethodPost, "/api/deck/new", `{"shuffled": false}`)

	assert.Equal(t, http.StatusOK, response.Code)

	bodyRes := convertBodyToMap(response.Body.String())

	assert.NotNil(t, bodyRes["deck_id"])
	assert.Equal(t, false, bodyRes["shuffled"])
	assert.Equal(t, float64(52), bodyRes["remaining"])
}

func TestCreateDeckWithShuffledTrueAndCards(t *testing.T) {
	response := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": true, "cards": ["AS", "KD", "AC", "2C", "KH"]}`)

	assert.Equal(t, http.StatusOK, response.Code)

	bodyRes := convertBodyToMap(response.Body.String())

	assert.NotNil(t, bodyRes["deck_id"])
	assert.Equal(t, true, bodyRes["shuffled"])
	assert.Equal(t, float64(5), bodyRes["remaining"])
}

func TestCreateDeckWithInvalidCards(t *testing.T) {
	response := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["NOK"]}`)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	bodyRes := convertBodyToMap(response.Body.String())

	assert.Equal(t, "All the cards must be valid", bodyRes["error"])
}

func setupTestApi() (*gin.Engine, *httptest.ResponseRecorder) {
	router := setupRouter()
	w := httptest.NewRecorder()

	return router, w
}

func makeBodyReq(param string) *bytes.Reader {
	return bytes.NewReader([]byte(param))
}

func callApi(method, uri, body string) *httptest.ResponseRecorder {
	router, responseRecorder := setupTestApi()

	bodyReader := makeBodyReq(body)
	req, _ := http.NewRequest(method, uri, bodyReader)

	router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func convertBodyToMap(bodyString string) map[string]interface{} {
	bodyRes := map[string]interface{}{}
	json.Unmarshal([]byte(bodyString), &bodyRes)

	return bodyRes
}
