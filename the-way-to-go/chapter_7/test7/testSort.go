package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{2, 1, 3, 4, 5, 6}
	c := []int{2, 1, 0}
	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(sort.IntsAreSorted(b))
	fmt.Println(sort.IntsAreSorted(c))

	fmt.Println(a)
	fmt.Println(sort.SearchInts(a, 0))
	fmt.Println(sort.SearchInts(a, 1))
	fmt.Println(sort.SearchInts(a, 2))
	fmt.Println(sort.SearchInts(a, 3))
	fmt.Println(sort.SearchInts(a, 4))
	fmt.Println(sort.SearchInts(a, 5))
	fmt.Println(sort.SearchInts(a, 6))
	fmt.Println(sort.SearchInts(a, 7))
	fmt.Println(c)
	sort.Ints(c)
	fmt.Println(c)

}
