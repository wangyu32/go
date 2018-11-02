package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	p(a)
	e := append(a[:3], a[4:]...)
	p(a)
	p(e)

}

func p(a []int) {
	fmt.Println(a)
}
