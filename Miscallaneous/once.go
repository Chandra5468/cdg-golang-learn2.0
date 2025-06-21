package main

import (
	"fmt"
	"sync"
)

/*
Ensuring One-Time Execution with sync.Once
Sometimes, you might have a piece of code that you only want to run once,
no matter how many goroutines are involved.
This is common with setup functions or tasks that should happen only once,
like initializing a resource. Go’s sync.Once provides a simple way to ensure that a function is executed only once,
even if multiple goroutines try to run it at the same time.
*/

func execOnce() {
	once := &sync.Once{}
	wg := &sync.WaitGroup{}

	setup := func() {
		fmt.Println("Initializing configuration....")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Go routine %d is calling the setup \n", id)
			once.Do(setup)
		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines finished...")
}
