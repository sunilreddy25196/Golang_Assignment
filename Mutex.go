package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
