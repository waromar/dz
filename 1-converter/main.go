package main
import ( 
	"fmt" 
)
func main() {
	const USDinEUR float64 = 0.8497
	const USDinRUB float64 = 75.22
	const EURinRUB float64 = USDinEUR * USDinRUB
	
	currency := inputCurrency()
	value := inputValue()
	goalCurency := inputGoalCurrency(currency)
	fmt.Println(currency, value, goalCurency)
}
func inputCurrency() string {
	fmt.Println("Введите исходную валюту: RUB, EUR, USD")
	var currency string
	fmt.Scan(&currency)
	if currency != "RUB" && currency != "EUR" && currency != "USD" {
		fmt.Println("Ошибка ввода валюты")
		return inputCurrency()
	}
	return currency
}

func inputGoalCurrency(currency string) string {
	var goalCurrency string

	if currency == "RUB" {
		fmt.Println("Введите целевую валюту: EUR, USD")
	} else if currency == "EUR" {
		fmt.Println("Введите целевую валюту: RUB, USD")
	} else if currency == "USD" {
		fmt.Println("Введите целевую валюту: EUR, RUB")
	} else {
		fmt.Println("Ошибка: Не существующая валюта")
	}
	fmt.Scan(&goalCurrency)
	if currency == goalCurrency {
		fmt.Println("Ошибка: Валюты одинаковые")
		return inputGoalCurrency(currency)
	} 
	return goalCurrency
}
func inputValue() float64 {
	var value float64
	fmt.Println("Введите колличество валюты")
	fmt.Scan(&value)
	return value
}




	
//func scanNumber(number *float64) {
//	fmt.Scan(number)
//}
//func calculate(num float64, currency1, currency2 string){

//}



	// if currency != "RUB" {
	// 	fmt.Println("Ошибка ввода валюты")
	// 	return inputCurrency1()
	// } else if currency != "EUR"{
	// 	fmt.Println("Ошибка ввода валюты")
	// 	return inputCurrency1()
	// } else if currency != "USD" {
	// 	fmt.Println("Ошибка ввода валюты")
	// 	return inputCurrency1()
	// }