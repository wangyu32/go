package main

import "fmt"

//练习7.11
func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringInToSlice(s, in, 0)
	fmt.Println(res)

	array := []int{1, 2, 3, 4, 5, 6}

}

func InsertStringInToSlice(slice []string, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	//func copy(dst, src []T) int` copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
