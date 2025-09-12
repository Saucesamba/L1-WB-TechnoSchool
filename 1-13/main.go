package main

import "fmt"

func main() {
	a := 13
	b := 7
	fmt.Printf("a=%d, b=%d\n", a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a=%d, b=%d\n", a, b)
}
