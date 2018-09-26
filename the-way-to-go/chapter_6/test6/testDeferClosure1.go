package main

//学习地址 defer、return、返回值之间执行顺序的坑
// https://blog.csdn.net/qq_22063697/article/details/74892728
import (
	"fmt"
)

func main() {
	fmt.Println("return:", a()) // 打印结果为 return: 0
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}
