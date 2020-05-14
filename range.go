package main

import "fmt"

func main() {
	//range 遍历数组，支持同时得到index和value
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range 遍历map，支持同时得到key value
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k,v := range kvs {
		//字符串占位
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range 遍历map 只取key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range 遍历字符串 index ascii(int)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}