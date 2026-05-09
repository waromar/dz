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
		goalCurrency,_ = inputGoalCurrency(currency)
	}
	// вызов функции для расчета
	sum := calculate(value, currency, goalCurrency)
	fmt.Println("Вывод:", sum)
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
		_, err := fmt.Scan(&goalCurrency)
		if err != nil {
			return "", fmt.Errorf("Ошибка валюты")
		}
		if goalCurrency != "EUR" && goalCurrency != "USD" {
			return "", fmt.Errorf("Введена бурда")
		 }

	case "EUR":
		fmt.Println("Введите целевую валюту: RUB, USD")
		_, err := fmt.Scan(&goalCurrency)
		if err != nil {
			return "", fmt.Errorf("Ошибка валюты")
		}
		if goalCurrency != "RUB" && goalCurrency != "USD" {
			return "", fmt.Errorf("Введена бурда")
		}
	case "USD":
		fmt.Println("Введите целевую валюту: RUB, EUR")
		_, err := fmt.Scan(&goalCurrency)
		if err != nil {
			return "", fmt.Errorf("Ошибка валюты")
			
		}
		if goalCurrency != "RUB" && goalCurrency != "EUR" {
			return "", fmt.Errorf("Введена бурда")
		}
	}


	return goalCurrency, nil
}
func calculate(value float64, currency string, goalCurrency string) float64 {
if currency == "RUB" && goalCurrency == "EUR"	{
 return value * RUBinEUR
} else if currency == "USD" && goalCurrency == "EUR" {
 return value * USDinEUR
} else if currency == "RUB" && goalCurrency == "USD" {
 return value * RUBinUSD
} else if currency == "USD" && goalCurrency == "RUB" {
 return value * USDinRUB 
} else if currency == "EUR" && goalCurrency == "RUB" {
 return value * EURinRUB
} else if currency == "EUR" && goalCurrency == "USD" {
 return value * EURinUSD
} 
return value
}

