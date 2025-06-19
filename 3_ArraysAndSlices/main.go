package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Arrays and slices")

	var nums = make([]int, 4, 10)

	// 4 -> Length
	// 10 -> Capacity

	// fmt.Println("This is nums", nums)

	nums = append(nums, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("This is nums", nums, len(nums), cap(nums))
	// Now capacity got increased to 20. As we have 11 elements.

	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4
	// nums[4] = 5
	// nums[5] = 1
	// nums[6] = 1
	// nums[7] = 2
	// nums[8] = 3
	// nums[9] = 4
	// nums[10] = 5
	// nums[11] = 1
	// fmt.Println(nums, len(nums), cap(nums))
	// This will give index out of range with length 4

	// copy() USING COPY Function
	nums2 := make([]int, 4, 10)
	copy(nums2, nums)

	fmt.Println("This is nums2 ", nums2, len(nums2), cap(nums2))
	// THis will copy only 4 elements as length is 4

	fmt.Println(nums[4:7])

	// COMPARING TWO SLICES

	fmt.Println("We are comparing two slices .... ", slices.Equal(nums, nums2))
}
