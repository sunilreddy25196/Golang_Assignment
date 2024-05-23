package main

import (
	"fmt"
	"time"
)

func task(input int) {
	time.Sleep(time.Duration(input) * time.Second)
	fmt.Println("Task completed")
}

func main() {
	done := make(chan struct{})
	go func() {
		task(1)
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Task completed in 3 seconds")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occured")
	}
}
