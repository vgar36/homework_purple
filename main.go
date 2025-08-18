package main

import "fmt"

const USDToEUR = 0.92
const USDToRUB = 92.50
const EURToRUB = USDToRUB / USDToEUR

func getInput() (amount float64, from string, to string) {
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&from)
	fmt.Print("Введите целевую валюту (USD, EUR, RUB): ")
	fmt.Scan(&to)

	return amount, from, to
}

func convert() (amount float64, from string, to string) {

	return result
}

func main() {

	fmt.Printf("1 EUR = %.2f", EURToRUB)
}
