package main

import (
	"fmt"
	"sync"
)

func Square(x int, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Printf("square(%d)\n", x*x)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{2, 4, 6, 8, 10}
	for _, x := range arr {
		go Square(x, wg)
	}
	wg.Wait()
}
