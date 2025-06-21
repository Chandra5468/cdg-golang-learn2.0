package main

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("Hello world")
}

func cooridnate() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Go routine %d working ---\n", id)
		}(i)

	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
