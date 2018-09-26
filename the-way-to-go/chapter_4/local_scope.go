package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	a := "O" //只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
	print(a)
}
