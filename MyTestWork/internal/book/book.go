package book

import "errors"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	Borrowed bool
}

func New(title, author string, id, year int) (*Book, error) {
	if title == "" || author == "" || id <= 0 || year < 0 {
		return nil, errors.New("Ошибка")
	}
	return &Book{
		ID:       id,
		Title:    title,
		Year:     year,
		Borrowed: false,
	}, nil
}
