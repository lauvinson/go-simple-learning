package main

import (
	"fmt"
	"time"
)

func main() {
	//使用make创建channel
	message := make(chan string)

	//channel <- value 发送
	go func() {message <- "ping"}()

	//variable <- channel 接收
	msg := <- message
	fmt.Println(msg)

	//channel buffering buffer可滞后接收
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel1"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//channel synchronization 同步通道
	done := make(chan bool, 1)
	go worker(done)
	<-done

	//channel directions 通道方向
	//当channel作为函数参数时，可以指定通道输入还是输出

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	input(pings, "hello")
	output(pings, pongs)
	fmt.Println(<-pongs)
}

//chan type means 接受普通通道
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

//chan <- string means 接受string到通道
func input(pings chan <- string, msg string) {
	pings <- msg
}
//<-chan string means 从通道输出string
func output(pings <-chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}