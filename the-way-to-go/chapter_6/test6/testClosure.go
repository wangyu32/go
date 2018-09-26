package main

import "fmt"

func main() {

	sum := 0
	sum = func() int {
		s := 0
		for i := 1; i <= 1e6; i++ {
			s += i
		}
		return s
	}()

	fmt.Printf("sum=%d\n", sum)

}
