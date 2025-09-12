package main

import (
	"fmt"
	"sync"
)

func worker(x map[int]int, ind int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 25; i++ {
		x[ind]++
	}
}

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	mm := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(mm, i, wg, mu)
	}
	wg.Wait()
	fmt.Println(mm)
}
