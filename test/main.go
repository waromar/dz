package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// **Описание**: Создайте программу для проверки существования файла и создания пустого Vault при его отсутствии
//
// **Входные данные**: Имя файла "vault.json" для проверки существования
//
// **Выходные данные**:
// - Если файл существует: сообщение "Файл vault.json найден"
// - Если файл не существует: создание файла vault.json с пустым Vault и сообщение "Создан новый vault.json"
//
// **Ограничения**:
// - Используйте os.Stat для проверки существования файла
// - Используйте os.IsNotExist для определения отсутствия файла
// - При отсутствии файла создайте пустую структуру Vault с пустым слайсом Accounts
// - Установите поле Updated в значение "2024-01-22"
// - Сериализуйте пустой Vault в JSON и сохраните в файл
// - Обработайте возможные ошибки
//
// **Примеры**:
// Входные данные: проверка файла "vault.json" (файл отсутствует)
// Выходные данные:
// - Создан файл vault.json с содержимым:
// ```json
// {
//   "accounts": [],
//   "updated": "2024-01-22"
// }
// ```
// - Сообщение: "Создан новый vault.json"
//
// Входные данные: проверка файла "vault.json" (файл существует)
// Выходные данные: сообщение "Файл vault.json найден"

type Account struct {
	URL  string `json:"url"`
	User string `json:"username"`
	Pass string `json:"password"`
}

type Vault struct {
	Accounts []Account `json:"accounts"`
	Updated  string    `json:"updated"`
}

func main() {
	// Ваш код здесь
	_, err := os.Stat("vault.json")
	if err == nil {
		fmt.Println("Файл vault.json найден")
		return
	} else if os.IsNotExist(err) {
		v := Vault{Accounts: []Account{}, Updated: "2024-01-22"}
		bytes, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		file, err := os.Create("vault.json")
		
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		file.Write(bytes)
	} else {
		fmt.Println(err)
		return
	}

}
