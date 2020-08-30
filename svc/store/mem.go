package store

import (
	"strings"

	"okki.hu/konyvmuzeum/model"
)

type Mem struct {
	Data []model.Book
}

func (m *Mem) List() []model.Book {
	return m.Data
}

func (m *Mem) Search(q string) []model.Book {
	var result []model.Book
	for _, b := range m.Data {
		if matches(b, q) {
			result = append(result, b)
		}
	}
	return result
}

func matches(b model.Book, q string) bool {
	return contains(b.Author, q) ||
		contains(b.Title, q) ||
		contains(b.ISBN, q) ||
		contains(b.Publisher, q)
}

func contains(prop, q string) bool {
	return strings.Contains(
		strings.ToLower(prop),
		strings.ToLower(q))
}
