package main

import (
	"fmt"
)

//练习7.12
func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	res := RemoveStringSlice(s, 1, 3)
	fmt.Println(s)
	fmt.Println(s[:3])
	fmt.Println(s[3:])
	fmt.Println(res) //M Q R
}

//删除 [start, end)区间的数据
func RemoveStringSlice(slice []string, start int, end int) []string {
	result := make([]string, len(slice)+start-end-1)
	at := copy(result, slice[:start])
	at += copy(result[at:], slice[end+1:])
	return result
}
