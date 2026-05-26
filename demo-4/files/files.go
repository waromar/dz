package files

import (
	"os"
	"fmt"
)

func ReadFile() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content string, name string) {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
		fmt.Println("Запись успешна")
		file.Close()
}