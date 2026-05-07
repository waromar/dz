package main
import "fmt"
const (
	USDinEUR float64 = 0.8497
	USDinRUB float64 = 75.22
)
func main() {
	//Рассчитать EUR в RUB на основании первых двух
	var EUR float64 = 100
	EURinRUB := EUR / USDinEUR * USDinRUB
	fmt.Println(EURinRUB)
}
