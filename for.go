package main

import "fmt"

func main() {

	//简单循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//临时变量，跳出判断，自增
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//break
	for {
		fmt.Println("loop")
		break
	}
}
