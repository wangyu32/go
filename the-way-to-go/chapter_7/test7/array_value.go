package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)
	f1(arr)
	fmt.Println(arr)
	f2(&arr)
	fmt.Println(arr)
	fmt.Println("------------")
	main0()

}

func f1(a [5]int) {
	a[0] = 1

	fmt.Println(a)
}

func f2(a *[5]int) {

	a[0] = 2
	fmt.Println(a)
}

func main0() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	arr2 := arr1
	arr2[2] = 100

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array arr1 at index %d is %d\n", i, arr1[i])
	}
	fmt.Println()
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Array arr2 at index %d is %d\n", i, arr2[i])
	}
}
