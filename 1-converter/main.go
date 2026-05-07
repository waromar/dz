package main
import "fmt"
func main() {
	const USDinEUR float64 = 0.8497
	const USDinRUB float64 = 75.22
	const EURinRUB float64 = USDinEUR * USDinRUB

}
func scanNumber(number float64) {
	fmt.Scan(&number)
}
func calculate(num float64, currency1, currency2 string){

}