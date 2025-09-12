package main

import "fmt"

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mm := make(map[int][]float32)
	for i := 0; i < len(arr); i++ {
		if mm[int(arr[i])/10*10] == nil {
			mm[int(arr[i])/10*10] = make([]float32, 0)
			mm[int(arr[i])/10*10] = append(mm[int(arr[i])/10*10], arr[i])
		} else {
			mm[int(arr[i])/10*10] = append(mm[int(arr[i])/10*10], arr[i])
		}
	}
	fmt.Println(mm)
}
