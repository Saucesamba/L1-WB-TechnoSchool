package main

import (
	"fmt"
	"time"
)

func producer() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}

	}()
	return ch
}

func square(ch chan int) chan int {
	cha := make(chan int)
	go func() {
		defer close(cha)
		for val := range ch {
			cha <- val * val
			time.Sleep(time.Second)
		}
	}()
	return cha
}

func reader(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
func main() {
	reader(square(producer()))
}
