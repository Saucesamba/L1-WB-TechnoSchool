package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//завершение по условию
func workerIf() {
	i := 0
	for i < 100 {
		i++
		if i == 50 {
			return
		}
		time.Sleep(30 * time.Millisecond)
	}
}

//котекст с отменой
func workerCtxCancel(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work stopped")
			return
		default:
			i++
			time.Sleep(time.Second)
		}
	}
}

//контекст с таймаутом
func workerCtxTimeOut(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work stopped")
			return
		default:
			i++
			time.Sleep(time.Second)
		}
	}
}

//контекст с дедлайном
func workerCtxDeadline(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work stopped")
			return
		default:
			i++
			time.Sleep(time.Second)
		}
	}
}

//завершение по сигналу с канала
func workerCh(ch <-chan int) {
	select {
	case v := <-ch:
		fmt.Println("Worker stopped", v)
		return
	default:
		fmt.Println("Worker running")
		time.Sleep(time.Second)
	}
}

//завершение по goex
func workerGoex() {
	defer fmt.Println("Goex stopped")
	time.Sleep(time.Second * 10)
	runtime.Goexit()
}

func main() {
}
