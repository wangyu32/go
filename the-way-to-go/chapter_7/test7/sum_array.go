package main

import "fmt"

func main() {
	// array := [4]float32{1.1, 2.2, 3.3, 4.0}  // 数组
	array := []float32{1.1, 2.2, 3.3, 4.0} //切片
	x := Sum(array)
	fmt.Println(x)
}

/*
数组
func Sum(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}
*/
func Sum(a []float32) (sum float32) {
	for _, v := range a {
		sum += v
	}
	return sum
}
