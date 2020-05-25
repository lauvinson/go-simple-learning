package main

import (
	"fmt"
	"time"
)

func main() {
	//超时

	c1 := make(chan string, 1)
	go func() {
		//等待两秒
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	//一秒钟后执行
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		//等待两秒钟
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	//三秒种后执行
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
