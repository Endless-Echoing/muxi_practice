package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("\nAll workers are done")
	fmt.Printf("Final counter value: %d\n\n", counter)

	letterCh := make(chan struct{})
	numberCh := make(chan struct{})
	done := make(chan struct{})

	go printletter(letterCh, numberCh, done)
	go printnum(letterCh, numberCh, done)

	letterCh <- struct{}{}

	<-done
	fmt.Println()
}
