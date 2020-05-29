package main

import "fmt"

var (
	CLOSED = make(chan int)
)

func init() {
	close(CLOSED)
}

func a() chan int {
	return CLOSED
}

func main() {
	fmt.Println("start")
	select {
	//case 必须有IO流动，返回nil会遭到阻塞
	case <-a():
		fmt.Println("CLOSED")
	default:
		fmt.Println("default")
	}
	fmt.Println("end")
}
