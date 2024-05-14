package main

import "fmt"

//map 字典
/*
哈希表，key => value，存储的key是经过哈希运算的
使用map的时候一定要分配空间
*/
func main(){
	//使用map分配空间
	idNames := make(map[int]string, 10)
	idNames[0] = "yuki"
	idNames[1] = "umino"
	idNames[5] = "amiya"

	fmt.Println(idNames)

	//key根据哈希表遍历
	for key, value := range idNames{
		fmt.Println(key,value)
	}

	/*
	1、在map中不存在访问越界的问题，它认为所有的key都是有效的，所以访问一个不存在的key不会崩溃，返回这个类型的默认值（零值）
	2、零值： bool => false, 数字 => 0, 字符串 => 空
	*/
	name3 := idNames[3]
	fmt.Println(name3,"test")

	//判断某key值对应的空间是否存在数据
	value, flag := idNames[1]
	if flag{
		fmt.Println("存在",value)
	}else{
		fmt.Println("不存在",value)
	}

	// 删除无效的key，不会报错
	delete(idNames,100)
	fmt.Println(idNames)

	// 多个并发处理，需要对map进行上锁
}