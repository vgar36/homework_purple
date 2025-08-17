package main

import "fmt"

func main() {
	const USDToEUR = 0.92
	const USDToRUB = 92.50
	const EURToRUB = USDToRUB / USDToEUR
	fmt.Printf("1 EUR = %.2f", EURToRUB)
}
