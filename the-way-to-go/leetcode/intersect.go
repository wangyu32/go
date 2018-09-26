package main

import "fmt"

func main() {

	fmt.Println("vim-go")

	printArray(intersect(nil, nil))
	printArray(intersect([]int{1}, nil))
	printArray(intersect([]int{3}, nil))
	printArray(intersect([]int{1, 2, 3}, nil))
	printArray(intersect(nil, []int{4}))
	printArray(intersect(nil, []int{2, 3, 4}))

}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 1 {
		return nums1
	}
	if len(nums2) == 1 {
		return nums2
	}

	return nil
}

func printArray(array []int) {
	if array == nil {
		fmt.Println(nil)
		return
	}

	for i := 0; i < len(array); i++ {
		fmt.Printf("%d, ", array[i])
	}

	fmt.Println()
}
