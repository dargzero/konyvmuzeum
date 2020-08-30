package model

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
	Publisher string `json:"publisher"`
	Language  string `json:"language"`
	Pages     int    `json:"pages"`
}
