package main

import "fmt"

func chann() {
	counter := make(chan int)

	go func() {
		count := 0
		for i := 0; i < 5; i++ {
			count++
		}
		counter <- count
	}()

	fmt.Println("Final counter : ", <-counter)
}
