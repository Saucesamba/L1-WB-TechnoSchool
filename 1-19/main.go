package main

import "fmt"

func reverce(str string) string {
	st := []rune(str)
	for i := 0; i < len(st)/2; i++ {
		st[i], st[len(st)-1-i] = st[len(st)-1-i], st[i]
	}
	return string(st)
}
func main() {
	fmt.Println(reverce("главрыба"))
}
