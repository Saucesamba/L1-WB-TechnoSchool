package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d interupted \n", id)
			return
		default:
			fmt.Println("doing some work")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, ctx, wg)
	}

	go func() {
		sig := <-ch
		fmt.Printf("Received signal: %v\n", sig)
		fmt.Println("Cancelling all work...")
		cancel()
	}()
	wg.Wait()
	fmt.Println("All work terminated")
}
