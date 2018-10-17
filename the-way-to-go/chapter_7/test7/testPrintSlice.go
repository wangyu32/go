package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	PrintSlice(arr)
}

func PrintSlice(slice1 [5]int) {
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}
