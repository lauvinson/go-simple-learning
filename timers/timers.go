package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	timer1 := time.NewTimer(2 * time.Second)

	//等待timer的C chan
	<-timer1.C
	fmt.Println("Timer 1 fired")

	//timer可以随时Stop
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	//只是等待可以使用time.Sleep
	time.Sleep(2 * time.Second)
}
