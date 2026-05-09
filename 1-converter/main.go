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
	// вызов функции ввода кол-ва валюты

	// вызов функции ввода целевой валюты
	

	// вызов функции для расчета
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

func inputValue() (float64, error) {

	return 0, nil
}