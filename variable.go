package main

import "fmt"

func main() {
	// 声明 赋值 字符串
	var a string = "initial"
	fmt.Println(a)

	// 同时声明 赋值
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明 赋值 布尔
	var d = true
	fmt.Println(d)

	// 声明 类型默认值
	var e int
	fmt.Println(e)

	// := 声明并赋值，系统自动推断类型
	f := "short"
	fmt.Println(f)

	// 同时声明 赋值 自动类型推断
	g, h, i := 1, 2, 3
	fmt.Println(g, h, i)
}