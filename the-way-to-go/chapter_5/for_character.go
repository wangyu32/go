package main

import "fmt"

func main() {
	num := 25
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("G")
		}
		fmt.Printf("\n")
	}
}
