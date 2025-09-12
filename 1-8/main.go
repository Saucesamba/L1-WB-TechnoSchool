package main

import "fmt"

func main() {
	var val, i, r int64
	fmt.Println("Input number and bit number to invert")
	fmt.Scanln(&val, &i)
	r = 1 << (i - 1)
	fmt.Println(r ^ val)
}
