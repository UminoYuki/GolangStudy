package main

import "fmt"

//make, copy, 数组浅拷贝和深拷贝

func main(){
	//定义一个定长数组
	names := [7]string{"北京","上海","广州","深圳","成都","重庆","天津"}
	// 基于names变量创建一个新的数组，数组仅包含"北京"、"上海"、"广州"
	//传统做法即声明并创建一个新数组并逐个赋值
	
	//通过切片可以基于一个数组，灵活创建数组
	// name2 := names[:5] 前五个
	// name2 := names[:5] 从5开始到最后
	name2 := names[0:3]	// 左闭右开
	fmt.Println("name2:",name2)

	//注意：此时name2和names指向的内存相同，例如修改name2的数据后，names的数据也会修改
	name2[2] = "合肥"
	fmt.Println(names)
	fmt.Println(name2)

	// 基于字符串进进行切片
	sub1 := "hello world"[6:8]
	fmt.Println(sub1)

	// 可以在创建空切片的时候，明确指定切片的容量，这样可以提高运行效率
	// 长度10，容量20
	str2 := make([]string,10,20)
	fmt.Println(len(str2),cap(str2))

	// 如果相让切片完全独立于原始数组，可以使用copy()函数完成
	namesCopy := make([]string,len(names))
	copy(namesCopy,names[:])
	namesCopy[0] = "Yuki"
	fmt.Println(namesCopy)
	fmt.Println(names)

}