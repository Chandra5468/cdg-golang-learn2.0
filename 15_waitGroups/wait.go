package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("This is task id ", id)
}

func main() {

	wg := &sync.WaitGroup{}

	for i := range 100 {
		wg.Add(1)
		go task(i, wg)
	}

	wg.Wait()

}
