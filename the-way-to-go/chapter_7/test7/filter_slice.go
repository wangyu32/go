package main

import "fmt"

// 练习7.10
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := Filter(s, even)
	fmt.Println(s1)

	fmt.Println("--------------")

	s2 := Filter(s, odd)
	fmt.Println(s2)
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	//for _, i := range s {
	for i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

func even(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func odd(a int) bool {
	return !even(a)
}
