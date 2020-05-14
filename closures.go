package main

import (
	"fmt"
)

/*
 * 闭包
 */

//这个函数返回的函数，对其内部变量持有‘副本’
func intSeq() func(int) int {
	i := 0
	return func(a int) int {
		i += a
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt(0))
	fmt.Println(nextInt(1))
	fmt.Println(nextInt(2))

	newInts := intSeq()
	fmt.Println(newInts(0))
}
