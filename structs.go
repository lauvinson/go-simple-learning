package main

import "fmt"

/*
 Go 中的struct与C中的struct非常相似，并且Go没有class
使用 type struct{} 定义结构，名称遵循可见性规则
支持指向自身的指针类型成员
支持匿名结构，可用作成员或定义成员变量
匿名结构也可以用于map的值
可以使用字面值对结构进行初始化
允许直接通过指针来读写结构成员
相同类型的成员可进行直接拷贝赋值
支持 == 与 !=比较运算符，但不支持 > 或 <
支持匿名字段，本质上是定义了以某个类型名为名称的字段
嵌入结构作为匿名字段看起来像继承，但不是继承
可以使用匿名字段指针
*/

type person struct {
	name string
	age int
}

//包含匿名结构作为内部成员变量
type personContact struct {
	name string
	contact struct{
		phone string
	}
}

//匿名字段结构
type noField struct {
	string
	int
}

//嵌入结构，在out中嵌入了in
type in struct {
	a string
}
type out struct {
	in
	name string
}

func newPerson(name string) *person{
	p := person{name: name}
	p.age = 42
	return &p
}

func (p *person) setAge(age int) {
	p.age = age
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.setAge(51)
	fmt.Println(sp.age)

	pc := personContact{name: "vinson", contact: struct{ phone string }{phone: "10086"}}
	fmt.Println(pc)

	nf := noField{"a", 1}
	fmt.Println(nf)

	out := out{in: in{a: "1"}, name: "vinson"}
	fmt.Println(out)

	//匿名结构
	a := &struct {
		string
		int
	}{
		"vinson",
		23,
	}
	fmt.Println(a)
}