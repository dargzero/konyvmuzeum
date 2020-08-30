package store

import "okki.hu/konyvmuzeum/model"

// Books can fetch and manipulate book data
type Books interface {
	// List all books
	List() []model.Book

	// Search using a given query
	Search(q string) []model.Book
}

var BookStore = &Mem{
	Data: []model.Book{
		{
			Title: "foo",
		},
		{
			Title: "foo",
		},
	},
}
