package main

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, item int) int {
	return binarySearchRecursive(arr, item, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, item, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == item {
		return mid
	} else if item < arr[mid] {
		return binarySearchRecursive(arr, item, low, mid-1)
	} else {
		return binarySearchRecursive(arr, item, mid+1, high)
	}
}

func main() {
	arr := []int{1, 52, 10, 0, 4, 33, 6, 3, 134, 1356, 36, 7}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(BinarySearch(arr, 3))
}
