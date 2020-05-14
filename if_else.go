package main

import "fmt"

func main() {
	//位运算取模
	if 7&(1) == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//没有else
	if 8&(4-1) == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 可以在if前声明变量 不需要括号
	if num := 9; num < 0 {
		fmt.Println(num, "is negarive")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}else {
		fmt.Println(num, "has multiple digits")
	}
}
