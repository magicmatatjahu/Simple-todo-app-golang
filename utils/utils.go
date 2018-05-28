package utils

import (
	"net/http"
	"encoding/json"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {

	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response, _ := json.Marshal(payload)
	w.Write(response)
}