package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	jsonBody := []byte(`{"shuffled": false}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest(http.MethodPost, "/api/deck/new", bodyReader)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	jsonStr := w.Body.String()
	bodyRes := map[string]interface{}{}
	json.Unmarshal([]byte(jsonStr), &bodyRes)

	assert.NotNil(t, bodyRes["deck_id"])
	assert.Equal(t, false, bodyRes["shuffled"])
	assert.Equal(t, float64(52), bodyRes["remaining"])
}
