package api

import (
	"net/http"

	"okki.hu/konyvmuzeum/store"
)

func BooksListHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, store.BookStore.List())
}

func BooksSearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	writeResponse(w, store.BookStore.Search(q))
}
