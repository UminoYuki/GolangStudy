//导入包，类似其他语言中的库（libraries）或者模块（modules）
package main

//导入fmt包
import (
	"fmt"
)

//main是一个特殊函数类型，所以程序的入口
func main() {
	fmt.Println("hello world, github!")
}

//包的管理严格控制，缺少了包或者导入了多余的包都无法编译通过
//import声明必须在package声明之后，随后才是变量，函数，常理等
//Go语言不需要添加； 编译器会自动把特定符号后换行符转化为分号