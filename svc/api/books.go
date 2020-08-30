package api

import (
	"encoding/json"
	"net/http"

	"okki.hu/konyvmuzeum/store"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(store.BookStore)
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
