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

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type rect struct {
	width, height int
}

//rect使用了指针实现
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius int
}

//circle使用了值实现
func (c circle) area() int {
	return c.radius * 3
}

func (c circle) perim() int {
	return c.radius * 2
}

func main() {
	r := rect{width: 1, height: 1}
	//传递指针
	measure(&r)

	c := circle{radius: 2}
	//传递值
	measure(c)
}
