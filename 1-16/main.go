package main

import "fmt"

func QuickSort(x []int) []int {
	if len(x) <= 1 {
		return x
	}
	pivot := x[len(x)/2]
	mid := make([]int, 0)
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 1; i < len(x); i++ {
		if x[i] > pivot {
			right = append(right, x[i])
		} else if x[i] < pivot {
			left = append(left, x[i])
		} else {
			mid = append(mid, x[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)
	for _, v := range mid {
		left = append(left, v)
	}
	for _, v := range right {
		left = append(left, v)
	}
	return left
}

func main() {
	arr := []int{1, 9, 2, 3, 4, 6, 1, 9, 6, 2, 0, 0, 3, 5, 6, 3}
	arrSorted := QuickSort(arr)
	fmt.Println(arrSorted)
}
