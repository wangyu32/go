package main

/**
问题 7.5** 假设我们有如下数组：`items := [...]int{10, 20, 30, 40, 50}`

a) 如果我们写了如下的 for 循环，那么执行完 for 循环后的 `items` 的值是多少？如果你不确定的话可以测试一下:)

```go
for _, item := range items {
		item *= 2
	}
	```

	b) 如果 a) 无法正常工作，写一个 for 循环让值可以 double。

*/

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		fmt.Println(item)

	}

	fmt.Println("-------------")
	for _, item := range items {
		fmt.Println(item)
	}
}
