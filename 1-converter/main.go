package main

import (
	"fmt"
)

const USDinEUR float64 = 0.8497
const USDinRUB float64 = 75.22
const EURinRUB float64 = USDinRUB / USDinEUR

func main() {
	// вызов функции ввода валюты
	currency, err := inputCurrency()
	if err != nil {
		fmt.Println(err)
		currency, _ = inputCurrency()
	}
	// вызов функции ввода кол-ва валюты
	value, err := inputValue()
	if err != nil {
		fmt.Println(err)
		value, _ = inputValue()
	}
	// вызов функции ввода целевой валюты
	goalCurrency, err := inputGoalCurrency(currency)
	if err != nil {
		fmt.Println(err)
		goalCurrency,_ = inputGoalCurrency()
	}
	// вызов функции для расчета
	fmt.Println(currency,value,goalCurrency)
}
func inputValue() (float64, error) {
	fmt.Println("Введите количество валюты:")
	var value float64
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, fmt.Errorf("Вводите цифры")
	}
	if value < 0 {
		return 0, fmt.Errorf("Колличество валюты отрицательное")
	}
	return value, nil
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
func inputGoalCurrency(currency string) (string, error) {
	var goalCurrency string
	
	switch currency {
	case "RUB":
		fmt.Println("Введите целевую валюту: EUR, USD")
	case "EUR":
		fmt.Println("Введите целевую валюту: RUB, USD")
	case "USD":
		fmt.Println("Введите целевую валюту: RUB, EUR")
	}
	_, err := fmt.Scan(&goalCurrency)
	if err != nil {
		return "", fmt.Errorf("Ошибка валюты")
	}
	return goalCurrency, nil

}