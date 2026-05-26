package main

import "fmt"

type Book struct {
    ID    int
    Title string
}

func main() {
    books := []Book{
        {ID: 1, Title: "Go"},
        {ID: 2, Title: "Rust"},
        {ID: 3, Title: "Python"},
    }

    // Задача: изменить название книги по ID.
    fmt.Println(changeBook(books, 2, "penis")) // {2 Rust}
		fmt.Println(books)
    fmt.Println(changeBook(books, 1, "huy")) // {0 }
		fmt.Println(books)
}

func changeBook(books []Book, id int, newTitle string) Book {
    // задача
		for book := range books {
			if books[book].ID == id {
				books[book].Title = newTitle
				return books[book]
			}
		}
		return Book{}
}
