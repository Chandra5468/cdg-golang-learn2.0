package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task ", id)
}

func main() {
	for i := range 10 {
		go task(i)
	}

	time.Sleep(2 * time.Second)
}
