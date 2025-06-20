package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// RESOURCE : https://learn.innoskrit.in/blog/print-odd-even-numbers-synchronously-using-goroutines-and-channels-in-go/

type Notify struct{}

func PrintOddEvenSync() {

	printOdd := make(chan Notify)

	printEven := make(chan Notify)

	closer := make(chan Notify)

	go func() {
		start := 1

		for {
			select {
			case <-printOdd:
				time.Sleep(time.Millisecond * 500)
				fmt.Println(start)

				start = start + 2

				printEven <- Notify{}

			case <-closer:
				return
			}
		}
	}()

	go func() {
		start := 2

		for {
			select {
			case <-printEven:
				time.Sleep(time.Millisecond * 500)
				fmt.Println(start)
				start = start + 2

				printOdd <- Notify{}
			case <-closer:
				return
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press enter to quit")

	printOdd <- Notify{}

	_, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("error while reading the console")
	}

	fmt.Println("Finished")
	closer <- Notify{}

	close(closer)
	close(printEven)
	close(printOdd)
}

func main() {
	PrintOddEvenSync()
}
