package reader

import (
	"errors"
	"test_project/internal/book"
)



type Reader struct {
	ID int
	Name string
	Books []*book.Book
}

func New(id int, name string) (*Reader, error) {
	if id <= 0 || name == "" {
		return nil, errors.New("Ошибка")
	}
	return &Reader{
		ID: id,
		Name: name,
	}, nil
}