package main

import "fmt"

//  import "arrayUtil"

func main() {
	fmt.Println("hello world")
	var array = []int{1, 2, 3, 4}
	//	arrayUtil.printArray(array)
	printArray(array)
	headToTail(array)

	printArray(array)
}

func headToTail(arr []int) {
	l := len(arr)
	temp := arr[0]
	arr[0] = arr[l-1]
	arr[l-1] = temp
}
