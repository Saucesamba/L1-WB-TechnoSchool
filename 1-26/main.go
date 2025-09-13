package main

import (
	"fmt"
	"strings"
)

func CheckString(str string) bool {
	m := make(map[string]int)
	str = strings.ToLower(str)
	for _, v := range str {
		if m[string(v)] == 0 {
			m[string(v)] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckString("abcd"))
	fmt.Println(CheckString("abCdefAaf"))
}
