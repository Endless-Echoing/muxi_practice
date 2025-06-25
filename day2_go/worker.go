package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 100; j++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
	fmt.Printf("Worker %d finished\n", id)
}
