package main

import "fmt"

/*
 *接口是golang中实现多态性的唯一好途径
 *在java中，通过implements显式继承某接口
 *go里，只要struct定义了某接口中得方法，那就是继承了接口
 */

type geometry interface {
	area() int
	perim() int
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 1, height: 1}
	measure(r)
}