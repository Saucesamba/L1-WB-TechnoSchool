package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	res := make([]int, 0)
	sort.Ints(a)
	sort.Ints(b)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	fmt.Println(res)
}
