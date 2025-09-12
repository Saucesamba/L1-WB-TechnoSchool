package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("1000000000000000000000", 10)

	b := new(big.Int)
	b.SetString("500000000000000000000", 10)
	res := new(big.Int)

	res.Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a, b, res)

	res.Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a, b, res)

	res.Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a, b, res)

	res.Div(a, b)
	fmt.Printf("Деление: %s / %s = %s\n", a, b, res)
}
