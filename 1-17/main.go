package main

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, item int) int {
	if len(arr) <= 1 && arr[0] != item {
		return -1
	}
	mid := len(arr) / 2
	if arr[mid] == item {
		return mid
	} else if item < arr[mid] {
		return BinarySearch(arr[:mid], item)
	} else {
		return BinarySearch(arr[mid+1:], item)
	}
}
func main() {
	arr := []int{1, 52, 10, 0, 4, 33, 6, 3, 134, 1356, 36, 7}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(BinarySearch(arr, 3))
}
