package main

import (
	"fmt"
)

// Channels are used for communicating between go routines as we cannot use return values here. As return is a blocking statement.

// func task(s chan string) {
// 	fmt.Println(<-s)
// }

// func processNum(n chan int) {
// 	fmt.Println(<-n)
// }

// func sum(result chan int, num1, num2 int) {
// 	result <- num1 + num2
// 	close(result)
// }

// Channels are blocking in nature
func main() {
	// messageChan := make(chan string)
	// go task(messageChan)
	// messageChan <- "Namaste" // channels are blocking in nature. So, using them after go routine call
	// close(messageChan)

	// // The send will block until another goroutine is ready to receive, and vice versa
	// n := make(chan int)

	// for i := range 100 {
	// 	go processNum(n)
	// 	n <- i
	// }

	// close(n)

	// result := make(chan int)

	// go sum(result, 3, 5)

	// fmt.Println(<-result)

	// Buffered channel will have space. So it will not block or give deadlock as receiver is not needed until sender is full.
	email := make(chan string, 100) // You can pass struct datatype here as well
	done := make(chan bool)
	// defer close(done)
	go emailSenderQueue(email, done)
	for i := 0; i < 1000; i++ {
		email <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(email)
	<-done

	close(done)

	chan1 := make(chan string)
	chan2 := make(chan int)

	go func() {
		chan1 <- "Namaste Bharat"
		close(chan1)
	}()

	go func() {
		chan2 <- 22
		close(chan2)
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Channel 1 received first", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Channel 2 received first", chan2Val)
		}
	}

}

func emailSenderQueue(email chan string, done chan bool) {
	for em := range email {
		fmt.Println("Sending email to ", em)
	}
	// defer close(done)
	done <- true

}
