package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// waitGroup acts as a synchronized counter.
	var waitGroup sync.WaitGroup
	for i := 1; i < 1024; i++ {
		// Increments the counter through waitGroup.Add(1)
		// each time you create a goroutine to scan a port.
		waitGroup.Add(1)
		go func(j int) {
			// Decrements the counter whenever
			// one unit of work has been performed.
			defer waitGroup.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	// Blocks until all the work has been done
	// and your counter has returned to zero.
	waitGroup.Wait()
}

// Better approach than listing-two-four, but still incorrect
// This might give inconsistent results.
//
// Scanning an excessive number of hosts or ports simultaneously
// may cause network or system limitations to skew your results.
