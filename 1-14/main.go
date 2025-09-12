package main

import (
	"fmt"
)

func resolver(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	}
}

func main() {
	arr := []interface{}{1, "cat", false, make(chan int)}
	for _, v := range arr {
		resolver(v)
	}
}
