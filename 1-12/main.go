package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make([]string, 0)
	mm := make(map[string]bool)
	for _, v := range arr {
		if _, ok := mm[v]; !ok {
			mm[v] = true
			res = append(res, v)
		}
	}
	fmt.Println(res)
}
