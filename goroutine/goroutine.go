package main

import (
	"fmt"
	"sync"
)

//协程

func f(from string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(7)

	go f("direct", &wg)

	go f("goroutine", &wg)

	go func(msg string, wg *sync.WaitGroup) {
		fmt.Println(msg)
		wg.Done()
	}("going", &wg)

	fmt.Println("done")

	wg.Wait()
}
