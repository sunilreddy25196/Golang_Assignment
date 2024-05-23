package main

import (
	"fmt"
	"time"
)

func Produce(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func Consume(ch chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	ch := make(chan int)
	go Consume(ch)
	go Produce(ch)
	time.Sleep(time.Second * 2)
}
