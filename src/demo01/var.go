package main

import "fmt"

//声明一个整型变量，默认值为0
var v1 int
//声明一个浮点类型变量，默认值为0
/**
float32 即为单精度浮点数格式，内存占用4个字节（32 bits）
float64 即为双精度浮点数格式，内存占用8个字节（64 bits）
**/
var v2 float32
var v3 float64 
//声明一个数组变量，默认为nil
var v4 [10]int
//声明一个切片变量，默认是nil
var v5 []float32
//结构体需要先定义再声明
type v6 struct{
	age int
	name string
}

//声明一个指针变量，默认值nil
//var v7 *int
//声明一个字典变量，默认值nil
//var v8 map[string]string
//声明一个方法变量
//var v9 func(x int)int
//声明一个接口变量
//var v9 interface{}


func main() {
	v4[1] = 1
	v4[0] = 100
	v2 = 3
	v3 = 10
	//声明一个结构体，并定义
	var p1 v6
	p1.age = 18
	p1.name = "yuki"

	//声明结构体的过程中就定义
	p2 := v6{22, "umino"}

	fmt.Println("hello var!")
	fmt.Println(v4[0])
	fmt.Println(p1.age, p1.name)
	fmt.Println(p2.age, p2.name)
}