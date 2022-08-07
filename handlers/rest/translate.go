package rest

import (
	"encoding/json"
	"net/http"
)


type Resp struct {
	Language string `json:"language"`
	Translation string `json:"translate"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	resp := Resp{
		Language: "english",
		Translation: "hello",
	}

	if err := enc.Encode(resp); err !=nil {
		panic("unable to encode response")
	}
}