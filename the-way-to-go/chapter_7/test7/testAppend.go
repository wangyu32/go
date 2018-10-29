package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	//1. 将切片 b 的元素追加到切片 a 之后：`a = append(a, b...)`
	c := append(a, b...)
	p(c)
	//2. 复制切片 a 的元素到新的切片 b 上：
	//
	//   ```go
	//      b = make([]T, len(a))
	//         copy(b, a)
	//	    ```
	d := make([]int, len(a))
	copy(d, a)
	p(d)

	//3. 删除位于索引 i 的元素：`a = append(a[:i], a[i+1:]...)`
	e := append(a[:3], a[4])
	p(e)

	//4. 切除切片 a 中从索引 i 至 j 位置的元素：`a = append(a[:i], a[j:]...)`
	f := append(c[:3], c[6:]...)
	p(f)

	//5. 为切片 a 扩展 j 个元素长度：`a = append(a, make([]T, j)...)`
	g := make([]int, len(a))
	copy(g, a)
	p(g)
	fmt.Println(len(g))
	g = append(g, make([]int, 4)...)
	fmt.Println(len(g))
	p(g)
	//6. 在索引 i 的位置插入元素 x：`a = append(a[:i], append([]T{x}, a[i:]...)...)`

	//7. 在索引 i 的位置插入长度为 j 的新切片：`a = append(a[:i], append(make([]T, j), a[i:]...)...)`

	//8. 在索引 i 的位置插入切片 b 的所有元素：`a = append(a[:i], append(b, a[i:]...)...)`

	//9. 取出位于切片 a 最末尾的元素 x：`x, a = a[len(a)-1], a[:len(a)-1]`

	//10. 将元素 x 追加到切片 a：`a = append(a, x)`

}

func p(a []int) {
	fmt.Println(a)
}
