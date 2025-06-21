package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	mu := &sync.Mutex{}

	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter :", counter)
	hello()
	cooridnate()

	chann()

	execOnce()
}
