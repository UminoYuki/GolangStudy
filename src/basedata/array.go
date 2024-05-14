package main

import "fmt"

func main(){
	//定长数组的定义方式
	nums := [10]int{1,2,3,4}
	
	//传统遍历方式
	for i:=0; i< len(nums); i++{
		fmt.Println(nums[i])
	}

	//通过for range方式遍历
	//key是索引号，value是对应每个数组空间的值
	for key, value := range nums{
		value += 1
		fmt.Println(nums[key],key,value)
	}
	
	for value := range nums{
		fmt.Println(value,nums[value])
	}

	//忽略索引号
	for _, value := range nums{
		fmt.Println(value)
	}

	//不定长数组
	//定义一个切片
	names := []string{"北京","上海","天津","南京"}
	for key, value := range names{
		fmt.Println(key,names[key],value)
	}
	fmt.Println(len(names))

	//追加一个数据到新的变量
	names2 := append(names,"成都","新疆")
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
	//新数组相比原数组容量(cap)翻倍
	fmt.Println(names2)
	fmt.Println(len(names2))
	fmt.Println(cap(names2))

	//追加数据到原数组
	names = append(names,"Mexico City")
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))

	nums3 := []int{}
	for i := 0; i<50; i++{
		nums3 = append(nums3, i)
		fmt.Println(len(nums3), cap(nums3))
	}
}