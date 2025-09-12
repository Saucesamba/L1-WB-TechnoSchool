package main

import (
	"flag"
	"fmt"
	"time"
)

func Worker(i int, ch <-chan int) {
	for j := range ch {
		fmt.Printf("Print value %d from %d gorutine \n", j, i)
	}

}
func main() {
	workerNum := flag.Int("wn", 10, "worker number")
	flag.Parse()
	ch := make(chan int, 10)

	for i := 0; i < *workerNum; i++ {
		go Worker(i, ch)
	}
	for i := 0; i < 1000; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
