package main

import "fmt"

//结构体需要先定义再声明
type v6 struct{
	age int
	name string
}

func main() {
	//声明一个结构体，并定义
	var p1 v6
	p1.age = 18
	p1.name = "yuki"

	//声明结构体的过程中就定义
	p2 := v6{22, "umino"}

	fmt.Println(p1.age, p1.name)
	fmt.Println(p2.age, p2.name)
}
