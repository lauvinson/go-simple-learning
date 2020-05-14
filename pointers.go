package main

import "fmt"

//&取址，*取值

func zeroval(i int){
	i = 0
}

func zeroptr(i *int){
	*i = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)

}
