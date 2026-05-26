package library

import (
	"fmt"
	"test_project/internal/book"
	"test_project/internal/reader"
)

type Library struct {
	Books []book.Book
	Readers map[int]*reader.Reader
	NextBookID int
	NextReaderID int
}

func New() *Library {
	return &Library{
		Readers: make(map[int]*reader.Reader),
	}
}

func (l *Library) AddBook(title, author string, year int) *book.Book {
	// добовляем ++ к айди новой книги
	l.NextBookID++
	// вызываем функцию новой книги и проверяем на ошибку
  book, err := book.New(title, author, l.NextBookID, year)
	if err != nil {
		fmt.Println(err)
	}
	// апендим книгу в книги библеотеки
	l.Books = append(l.Books, *book)
	// возвращаем книгу
	return book
}
func (l *Library) AddReader(name string) *reader.Reader {
	// добовляем ++ к айди новому читателю
	l.NextReaderID++
	reader, err := reader.New(l.NextReaderID, name)
	if err != nil {
		fmt.Println(err)
	}
	// добавляем читателя в список библиотеки
	l.Readers[l.NextReaderID] = reader
	// ретурним читателя
	return reader
}

func (l *Library) BorrowBook(readerID, bookID int) error {
	// проверяем через ОК есть ли читатаель с "таким" айди если нет то возвращаем ошибку
   r, ok := l.Readers[readerID]
	 if !ok {
		return fmt.Errorf("reader %d not found", readerID)
	 }
  // нужно создать пермеменную с типом бук.бук
  var foundBook book.Book
	// находим книгу по ID через рендж 
 for book := range l.Books {
	if l.Books[book].ID == bookID {
		foundBook = l.Books[book]
		break
	} else {
		return fmt.Errorf("book %d not found", bookID)
	}
 }
	// проверить занятость книги через иф
   if foundBook.Borrowed {
        return fmt.Errorf("book %d already borrowed", bookID)
    }
	// изменить занятость книги
  foundBook.Borrowed = true
	// апендим в книги читателя нашу найденную книгу
   r.Books = append(r.Books, &foundBook)
	// возвращаем nil
  return nil
}