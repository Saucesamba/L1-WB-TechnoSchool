package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	arr := strings.Split(str, " ")
	sb := strings.Builder{}
	for i := 0; i < len(arr)-1; i++ {
		arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	}
	for i := 0; i < len(arr); i++ {
		sb.WriteString(arr[i])
		sb.WriteString(" ")
	}

	fmt.Println(sb.String())
}
