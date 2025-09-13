package main

import "fmt"

func RemoveEl(arr []int, ind int) []int {
	copy(arr[ind:], arr[ind+1:])
	arr[len(arr)-1] = 0
	return arr[:len(arr)-1]

}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	res := RemoveEl(arr, 7)
	fmt.Println(res)
}
