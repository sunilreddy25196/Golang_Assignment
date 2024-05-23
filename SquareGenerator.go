package main

import (
	"fmt"
)

func Generator(out chan int, numbers []int) {
	for _, num := range numbers {
		out <- num
	}
	close(out)
}

func Square(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func Print(in chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	nums := []int{100, 30, 46}
	ch := make(chan int)
	square := make(chan int)
	go Generator(ch, nums)
	go Square(ch, square)
	Print(square)
}
