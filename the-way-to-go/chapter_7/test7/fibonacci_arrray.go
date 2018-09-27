package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var arr [50]int
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("i=%d num=%d \n", (i + 1), arr[i])
	}

}
