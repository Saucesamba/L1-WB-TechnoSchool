package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	value int
}

func increment(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	defer c.Unlock()
	c.value++
}

func main() {
	wg := &sync.WaitGroup{}
	c := Counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&c, wg)
	}
	wg.Wait()
	fmt.Println(c.value)
}
