package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	close(queue)
	for s := range queue {
		fmt.Println(s)
	}
}
