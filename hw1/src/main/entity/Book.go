package entity

type Book struct {
	Author    string
	Title     string
	PageCount int16
}

func NewBook(title, author string, pageCount int16) *Book {
	return &Book{
		Title:     title,
		Author:    author,
		PageCount: pageCount,
	}
}
