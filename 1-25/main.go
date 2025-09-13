package main

import (
	"fmt"
	"time"
)

func Sleeep(arg time.Duration) {
	t := time.NewTimer(arg)
	<-t.C
}

func main() {
	fmt.Println("Program started")
	fmt.Println("waiting for 3 seconds...")
	Sleeep(3 * time.Second)
	fmt.Println("Program ended")
}
