package store

import "okki.hu/konyvmuzeum/model"

type Mem struct {
	Data []model.Book
}

func (m *Mem) List() []model.Book {
	return m.Data
}

func (m *Mem) Search(q string) []model.Book {
	return m.Data
}
