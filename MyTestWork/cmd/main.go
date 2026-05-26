package main

import (
	"fmt"
	"test_project/internal/library"
)



func main() {
    lib := library.New()
    lib.AddBook("Go Programming", "hui", 2015)
    alice := lib.AddReader("Alice")
    err := lib.BorrowBook(1, 1)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(alice, alice.Books)
}

