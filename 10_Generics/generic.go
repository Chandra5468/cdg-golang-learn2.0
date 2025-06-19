package main

import "fmt"

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	sum := []int{1, 2, 64, 25, 22}

	printSlice(sum)

	names := []string{"abc", "def", "xyz"}

	printSlice(names)

	myStack := stack[string]{
		elements: []string{"golang"},
	}

	fmt.Println(myStack)

	myStackNum := stack[int]{
		elements: []int{43, 393},
	}

	fmt.Print(myStackNum)
}
