package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
	fmt.Println(valss())
}

func vals() (int, int) {
	return 3, 7
}

func valss() (int, int, int ,int){
	return 1, 2, 3, 4
}
