package main

import "fmt"

func main() {
	var arr = [4]byte{0, 1, 1, 0}
	var slice1 []byte = arr[2:]
	var slice2 []byte = arr[0:1]
	Append(slice1, slice2)
	PrintSlice(slice1)
}

func Append(slice, data []byte) []byte {

}

func PrintSlice(data []byte) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d:%d", i, data[i])
	}
}
