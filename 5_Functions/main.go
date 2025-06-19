package main

import (
	"fmt"
	"math/rand"
)

func add(a int, b int) int {
	return a + b
}

func proceesIt(fn func(int) int) {
	fmt.Println(fn(1))
}

func retFn() func(a int) int {
	return func(a int) int {
		return 100
	}
}

func variadicFunction(nums ...int) int {
	sum := 0
	for _, x := range nums {
		sum = sum + x
	}

	return sum
}

func closu() func() int {
	var count int = rand.Int()

	return func() int {
		return count + 1
	}
}

func main() {
	fmt.Println(add(2, 3))

	fn := func(int) int {
		return 101
	}

	proceesIt(fn)

	x := retFn()

	fmt.Println(x(50))

	// Calling variadic function

	fmt.Println(variadicFunction(1, 3, 2, 1, 3, 2, 1, 32313))

	slicer := []int{2, 14, 244, 949, 293932}

	fmt.Println(variadicFunction(slicer...))

	// CLOSURES
	c := closu()
	fmt.Println(c())
}
