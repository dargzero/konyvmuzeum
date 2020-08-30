package api

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, result interface{}) {
	bytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
