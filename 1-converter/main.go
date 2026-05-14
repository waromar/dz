package main

import (
	"fmt"
)

const RUBinEUR float64 = USDinEUR / USDinRUB
const EURinUSD float64 = 1 / USDinEUR
const USDinEUR float64 = 0.8497
const USDinRUB float64 = 75.22
const EURinRUB float64 = USDinRUB / USDinEUR
const RUBinUSD float64 = 1 / USDinRUB

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
		goalCurrency, _ = inputGoalCurrency(currency)
	}
	
	// вызов функции для расчета
	result := calculate(value, currency, goalCurrency)
	fmt.Println("Вывод:", result, goalCurrency)
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
	var currency string
	fmt.Println("Введите исходную валюту: RUB, EUR, USD")

	_, err := fmt.Scan(&currency) 
	if err != nil{
	return "" ,fmt.Errorf("Ошибка ввода валюты") 
}
return currency, nil
}
func inputGoalCurrency(currency string) (string, error) {
	switch currency {
	case "RUB":
		return selectGoalCurrency("EUR", "USD")
	case "USD":
		return selectGoalCurrency("RUB", "EUR")
	case "EUR":
		return selectGoalCurrency("RUB", "USD")
	default: 
		return "", fmt.Errorf("Неизвестная бурмалда")
	}
	

}

func selectGoalCurrency(currency1, currency2 string) (string, error) {
	var goalCurrency string
		fmt.Printf("Введите целевую валюту: %v, %v\n", currency1, currency2)
		_, err := fmt.Scan(&goalCurrency)
		if err != nil {
			return "", fmt.Errorf("Неизвестная валюта")
		}
		if goalCurrency != currency1 && goalCurrency != currency2 {
			return "", 	fmt.Errorf("Введена бурмалда")
		}
	return goalCurrency, nil
}


func calculate(value float64, currency, goalCurrency string) float64 {
	switch {
	case currency == "RUB" && goalCurrency == "EUR":
		return value * RUBinEUR
	case currency == "RUB" && goalCurrency == "USD":
		return value * RUBinUSD
	case currency == "USD" && goalCurrency == "EUR":
		return value * USDinEUR
	case currency == "USD" && goalCurrency == "RUB":
		return value * USDinRUB
	case currency == "EUR" && goalCurrency == "USD":
		return value * EURinUSD
	case currency == "EUR" && goalCurrency == "RUB":
		return value * EURinRUB
	default:
		return 0
	}
}
