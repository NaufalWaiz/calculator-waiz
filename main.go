package main

import (
	"fmt"
	"calculator"
)

func main() {
	a, b := 10.0, 5.0

	fmt.println("penjumlahan :", calculator.Add(a, b))
	fmt.println("pengurangan :", calculator.Subtract(a, b))
	fmt.println("perkalian :", calculator.Multiply(a, b))
	fmt.println("pembagian :", calculator.Divide(a, b))

	result, err := calculator.Divide(a, 0)
	if err != nil {
		fmt.println("error:", err)
	} else {
		fmt.println("pembagian :", result)
	}
}
