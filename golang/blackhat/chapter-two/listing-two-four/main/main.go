package main

import (
	"fmt"
	"net"
)

// Launches a single goroutine per connection
//
// The main goroutine doesn't know to wait for
// the connection to take place.
//
// Code completes and exits as soon as the for
// loop finishes it iterations
func main() {
	for i := 1; i < 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
