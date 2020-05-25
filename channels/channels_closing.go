package main

import "fmt"

//关闭channel来标记没有更多等待
//value, more := <- channel

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//接收value和more
			j, more := <- jobs
			if more {
				fmt.Println("received:", j)
			}else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	//close the jobs channel
	close(jobs)
	fmt.Println("send all jobs")

	fmt.Println(<-done)
}