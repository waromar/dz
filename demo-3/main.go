package main

import "fmt"

func main() {
	fmt.Println(*nil)
	// fmt.Println("Приложение для закладок")
	// bookmarks := map[string]string{}
	// for {
	// 	variant := getMenu()
	// 	switch variant {
	// 	case 1:
	// 		printBookmar(bookmarks)
	// 	case 2:
	// 		bookmarks = addBookmark(bookmarks)
	// 	case 3:
	// 		bookmarks = deleteBookmark(bookmarks)
	// 	case 4:
	// 		break
	// 	}
	// }
}
func getMenu() int {
	var variant int
	fmt.Println("Выберете вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant

}
func printBookmar(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}

}

func addBookmark(bookmarks map[string]string) map[string]string {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Println("Введите название:")
	fmt.Scan(&newBookmarkKey)
	fmt.Println("Введите ссылку:")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue

	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var bookmarkKeyToDelete string
	fmt.Println("Введите название:")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
