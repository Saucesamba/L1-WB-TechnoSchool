package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	arr := strings.Split(str, " ")
	sb := strings.Builder{}
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	for i := 0; i < n; i++ {
		sb.WriteString(arr[i])
		sb.WriteString(" ")
	}

	fmt.Println(sb.String())
}
