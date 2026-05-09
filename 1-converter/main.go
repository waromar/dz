package main

import ( 
	"fmt" 
)
const USDinEUR float64 = 0.8497
const USDinRUB float64 = 75.22
const EURinRUB float64 = USDinRUB / USDinEUR

func main() {
	currency, err := inputCurrency()
	if err != nil {
		fmt.Println(err)
		currency, _ = inputCurrency()
	}
	fmt.Println(currency)
	value, err := inputValue()
	if err != nil {
		fmt.Println(err)
		value, _ = inputValue()
	}
	fmt.Println(value)
	// вызов функции ввода кол-ва валюты
	// вызов функции ввода целевой валюты
	
	// вызов функции для расчета
}
func inputValue() (float64, error) {
fmt.Println("Введите количество валюты:")
var value float64
fmt.Scan(&value)
if value < 0 {
	return 0, fmt.Errorf("Колличество валюты отрицательное")
} else {
	return 0, fmt.Errorf("Некорректный ввод")
}
	return 0, nil
}
// пример как надо сделать
func inputCurrency() (string, error) {
	fmt.Println("Введите исходную валюту: RUB, EUR, USD")
	var currency string
	fmt.Scan(&currency)
	if currency != "RUB" && currency != "EUR" && currency != "USD" {
		return "", fmt.Errorf("Ошибка ввода валюты")
	}
	return currency, nil
}

