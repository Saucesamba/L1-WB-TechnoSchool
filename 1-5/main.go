package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sender(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Sender is done")
			return
		default:
			ch <- i
			fmt.Println("Sender is send", i)
			time.Sleep(1 * time.Second)
		}
	}
}

func receiver(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Receiver is done")
			return
		case value, ok := <-ch:
			if !ok {
				fmt.Println("Receiver: channel closed")
				return
			}
			fmt.Printf("Receiver: received %d\n", value)
		}
	}

}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go sender(ctx, ch, wg)
	go receiver(ctx, ch, wg)

	go func() {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			cancel()

		}
	}()
	wg.Wait()
}
