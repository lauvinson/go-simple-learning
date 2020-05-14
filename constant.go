package main

import (
	"fmt"
	"math"
)

const s string = "constant"
const ss = true

func main() {
	fmt.Println(s)

	fmt.Println(ss)

	//任何位置声明常量
	const n = 500000000

	//任何精度处理
	const d = 3e20 / n
	fmt.Println(d)

	//转换类型
	fmt.Println(int64(d))

	//计算操作
	fmt.Println(math.Sin(n))
}
