package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
		operations := scanOperations()
		result := calculateMED(operations)
		fmt.Println("Ваш ответ:", result, "MED")
	}
}
func scanOperations() []float64 {
	operations := []float64{}
	fmt.Println("Введите числа через запятую: Пример ->(14, 32, 56)")
	var value string
	fmt.Scan(&value)
	operationsStr := strings.Split(value, ",")

	for _, elem := range operationsStr {
		operation, err := strconv.ParseFloat(elem, 64)
		if err != nil {
			fmt.Println(err)
		}
		operations = append(operations, operation)
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

func calculateMED(operations []float64) float64 {
	sortedOperations := sortSlice(operations)
	if len(sortedOperations)%2 == 0 {
		return (sortedOperations[len(sortedOperations)/2-1] + sortedOperations[len(sortedOperations)/2]) / 2
	} else {
		return sortedOperations[len(sortedOperations)/2]
	}
}
func sortSlice(operations []float64) []float64 {
	if len(operations) < 2 {
		return operations
	}
	pivot := operations[len(operations)/2]
	var less, equal, greaterc []float64
	for _, value := range operations {
		switch {
		case value < pivot:
			less = append(less, value)
		case value == pivot:
			equal = append(equal, value)
		default:
			greaterc = append(greaterc, value)
		}
	}
	return append(append(sortSlice(less), equal...), sortSlice(greaterc)...)
}
