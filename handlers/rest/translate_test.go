package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestTranslateHandler(t *testing.T) {
	response := httptest.NewRecorder()
	resquest := httptest.NewRequest(http.MethodGet, "/hello", nil)

	handler := http.HandlerFunc(TranslateHandler)

	handler.ServeHTTP(response, resquest)

	if response.Code != http.StatusOK {
		t.Errorf("want status 200, got %d", response.Code)
	}

	var resp Resp 
	json.Unmarshal(response.Body.Bytes(), &resp)

	if resp.Language != "english" {
		t.Errorf(`want language "english" but got %s`, resp.Language)
	}

	if resp.Translation != "hello" {
		t.Errorf(`want Translation "hello" but got %s`, resp.Translation)
	}
}