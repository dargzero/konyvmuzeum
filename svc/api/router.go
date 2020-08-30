package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	booksApi(router)
	return router
}

func booksApi(r *mux.Router) {
	r.HandleFunc("/books", BooksHandler)
}
