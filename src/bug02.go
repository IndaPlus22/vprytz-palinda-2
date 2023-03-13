package main

import (
	"fmt"
	"sync"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	// create wait group
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// create channel
	ch := make(chan int)

	go Print(ch, wg)
	for i := 1; i <= 11; i++ {
		ch <- i
	}

	// close the channel
	close(ch)

	// then wait for all goroutines to finish
	wg.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}
