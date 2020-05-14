package main

import "fmt"

func main() {
	//变长参函数，fmtprintln就是一个
	a := []int{1,2,3,4}
	//已有slice，使用...
	fmt.Println(p(a...))
}

func p(a ...int) int {
	j := 0
	for i := range a {
		j += i
	}
	return j
}
