package main

import "fmt"

func main() {
	//声明一个空数组
	var a [5]int
	fmt.Println("emp:", a)

	//设置指定下标的值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//获取length
	fmt.Println("len:", len(a))

	//声明并赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//二维数组 row2 col3
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j :=0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
