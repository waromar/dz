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

// func scanValue() float64 {
// 	var value float64
// 	fmt.Println("Введите число (0 для выхода)")
// 	fmt.Scan(&value)
// 	return value
// }
func scanOperations() []float64 {
	operations := []float64{}
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

// выполнить приказ 66
if len(sortedOperations) % 2 == 0 {
	return (sortedOperations[len(sortedOperations)/2 - 1] + sortedOperations[len(sortedOperations)/2]) / 2
} else {
	return sortedOperations[len(sortedOperations)/ 2]
}
}
func sortSlice(operations []float64) []float64 {
	// проверяем, если len = 1 то сортировать ничего не надо
if len(operations) < 2 {
	return operations
}
// задаем переменную pivot равную слайсу operapions, но с длинной в 2 раза меньше
pivot := operations[len(operations)/2]
// обьявляем переменные less, equal, greaterc как слайс с типом float64
var less, equal, greaterc []float64
// создаем цикл со switch который пинимает value
for _, value := range operations {
	switch {
// если value < pivot то в less добовляем value
	case value < pivot:
		less = append(less, value)
// если value == pivot то в equal добовляем value
	case value == pivot:
		equal = append(equal, value)
// "иначе" добовляем в greaterc, value
	default:
		greaterc = append(greaterc, value)
	}
}
// ретурним append append(a) и вызывам функцию sortSlice саму в себя с принимающей переменную greaterc less и equal
//  так же перепрезываем функцию sortSlice саму в себя с принимающей переменную greaterc
return append(append(sortSlice(less),equal...), sortSlice(greaterc)...)
}