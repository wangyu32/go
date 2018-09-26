package main

import "fmt"

func main() {
	var num1 = 15
	var num2 = 3
	var sum, pro, diff = fun1(num1, num2)

	fmt.Printf("num1 = %d,  num2 = %d, sum = %d, product = %d, diif = %d \n", num1, num2, sum, pro, diff)

	num1 = 4
	sum, pro, diff = fun2(num1, num2)

	fmt.Printf("num1 = %d,  num2 = %d, sum = %d, product = %d, diif = %d \n", num1, num2, sum, pro, diff)

}

func fun1(num1 int, num2 int) (int, int, int) {
	return num1 + num2, num1 * num2, num1 - num2
}

func fun2(num1 int, num2 int) (sum int, product int, diff int) {
	sum = num1 + num2
	product = num1 * num2
	diff = num1 - num2
	return
}
