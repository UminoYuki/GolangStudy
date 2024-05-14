package main

import "fmt"

func main(){
	//声明并定义
	name :="simple"
	//如果要换行使用`` (反引号)括起来
	uname := `
	./dawdw
	dawd
	- s x`
	fmt.Println("name",name)
	fmt.Println("uname",uname)
	//golong中使用len()函数来实现查看字符串长度
	var nameLength int
	var unameLength int
	nameLength = len(name)
	unameLength = len(uname)
	fmt.Println(nameLength)
	fmt.Println(unameLength)

	//for循环打印字符串
	for i:=0; i<len(name);i++{
		fmt.Println(i,name[i])
	}

	//字符串拼接
	a, b := "hello", "world"
	fmt.Println(a+b)

	// 使用const修饰为常量，不能修改，不用声，即不能使用:=
	const city = "Mexico"
	fmt.Println(city)
}