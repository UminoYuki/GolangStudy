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


//基础数据类型
//1、整型
var i = 123
var i8 int8 = 456
var i16 int16 = 789
var i32 int32 = 123
var i64 int64 = 123
//2、浮点型
var f = 123.0
var f32 float32 = 123.00
var f64 float64 = 345.66
//3、布尔类型
var b = false
var a bool = false
//4、字符串



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

	/*关键字
	定义变量: var
	定义常量: const
	**/
	//1、先定义，再赋值
	var name string
	name = "yuki"
	fmt.Println("name is ", name)
	//2、定义时赋值
	var age int = 25
	fmt.Println("age is ", age)

	fmt.Println("hello var!")
	fmt.Println(v4[0])
	fmt.Println(p1.age, p1.name)
	fmt.Println(p2.age, p2.name)
}