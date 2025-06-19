package main

import (
	"fmt"
	"maps"
)

// Similar to objects in js
func main() {
	map1 := make(map[string]string)

	map1["abx"] = "xyz"
	map1["aba"] = "daba"

	for x := range map1 {
		fmt.Println(x)
	}

	fmt.Println(map1["aba"])

	// If key does not exists in map -> It return zero value, empty string for string, 0 for int, false for boolean
	fmt.Println(map1["annana"])

	// Deleting a key

	delete(map1, "aba")

	fmt.Println(map1)

	k, ok := map1["aba"]

	if ok {
		fmt.Print("All okay this item exists ", k)
	} else {
		fmt.Println("Element not found", ok)
	}

	// Comparing two maps
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 4}

	fmt.Println(maps.Equal(m1, m2))

}
