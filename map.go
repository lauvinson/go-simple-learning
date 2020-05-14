package main

import "fmt"

func main() {

	//create an empty map, syntax make(map[key type]value type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//length
	fmt.Println("len:", len(m))

	//builtin `delete` removes key/value pairs from the map
	delete(m, "k2")
	fmt.Println("map:", m)

	//第一个为值，使用_即忽略，第二个值为是否存在
	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	//声明并初始化map
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)
}
