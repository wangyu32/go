package main

import "fmt"
import "math"

func main() {
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)

	//fmt.Printf("16 bit int is: %d %s \n", Uint8FromInt(127))
	//fmt.Printf("16 bit int is: %d %s \n", Uint8FromInt(129))
	println(math.MaxUint8)

	a3, a4 := Uint8FromInt(256)
	println(a3)
	println(a4)
	println(a4.Error())

	a1, a2 := Uint8FromInt(120)
	println(a1)
	println(a2)
	//	println(a2.Error())

}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 { // conversion is safe
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}
