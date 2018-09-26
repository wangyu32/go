package main

import "fmt"

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
