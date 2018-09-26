package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	a()
	f()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
