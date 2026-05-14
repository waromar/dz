package main

import "fmt"

func main() {
	var operation string
	fmt.Println("Выберете операцию: AVG, SUM, MED")
	fmt.Scan(&operation)
	switch operation {
	case "AVG":
		operations := scanOperations()
		result := calculateAVG(operations)
		fmt.Println("Ваш ответ:", result, "AVG")
	case "SUM":
		operations := scanOperations()
		result := calculateSUM(operations)
		fmt.Println("Ваш ответ:", result, "SUM")
	case "MED":
		

	}
}

func scanValue() float64 {
	var value float64
	fmt.Println("Введите число (0 для выхода)")
	fmt.Scan(&value)
	return value
}
func scanOperations() []float64 {
	operations := []float64{}
		for {
			value := scanValue()
			if value == 0 {
				break
			}
			operations = append(operations, value)
		}
return operations

}

func calculateAVG(operations []float64) float64 {
	var sum, count float64
	for _, elem := range operations {
		sum += elem
		count++
	}
	result := sum / count
	return result
}

func calculateSUM(operations []float64) float64 {
var sum float64
	for _, elem := range operations {
		sum += elem
	}
	result := sum 
	return result
}