package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res = plusPlus(1, 2, 3)
	fmt.Println(res)
}

func plusPlus(i int, i2 int, i3 int) int {
	return i + i2 + i3
}

func plus(i int, i2 int) int {
	return i + i2
}
