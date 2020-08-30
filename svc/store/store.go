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
			Title:     "The Book of Foo",
			Author:    "Mr. Big",
			ISBN:      "ABCDEF12345",
			Language:  "english",
			Pages:     160,
			Publisher: "The Publisher",
			Year:      2005,
		},
		{
			Title:     "The Book of Bar",
			Author:    "Ms. Small",
			ISBN:      "ABCDEF45365",
			Language:  "english",
			Pages:     75,
			Publisher: "The Publisher",
			Year:      2001,
		},
		{
			Title:     "The Book of Bar",
			Author:    "Mrs. Medium",
			ISBN:      "ABCDEF1D34G",
			Language:  "swedish",
			Pages:     240,
			Publisher: "The Publisher",
			Year:      2009,
		},
	},
}
