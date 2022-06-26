//每一个go程序都必须有一个包名
//每个go程序，都是go结尾的
//go程序没有.h，没有.o，只有.go
//一个package，包名，相当于命名空间
package main

//这是导入一个标准包fmt，format一般用于格式化输出输入
import "fmt"

//主函数所有函数必须使用func开头
//一个函数的返回值，不会放在func前面，而是放在参数后面
//函数的左大括号必须与函数名同行不能写到下一行
func main() {

	//go语言语句不需要使用分号结尾
	fmt.Println("hello world")
}
