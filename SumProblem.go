package main

import (
	"fmt"
)

func Calculatesum(n int, ch chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	ch <- sum
}
func main() {
	var input int
	fmt.Print("Enter Number:")
	fmt.Scan(&input)
	ch := make(chan int)
	go Calculatesum(input, ch)
	result := <-ch
	fmt.Println("Total sum:", result)
}
