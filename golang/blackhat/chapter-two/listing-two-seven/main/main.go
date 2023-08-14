package main

import (
	"fmt"
	"sync"
)

// woker function's channel will be used to receive work,
// and a pointer of WaitGroup will be used to track when
// a single work item has been completed.
func worker(ports chan int, waitGroup *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		waitGroup.Done()
	}
}

func main() {
	// This allows the channel to be buffered, which means
	// you can send it an item without waiting for a receiver
	// to read the item.
	ports := make(chan int, 100)
	var waitGroup sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &waitGroup)
	}
	for i := 1; i <= 1024; i++ {
		waitGroup.Add(1)
		ports <- i
	}
	waitGroup.Wait()
	close(ports)
}

// Output: numbers are printed in no paticular order.
