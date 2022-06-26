package main

import "fmt"

func main() {

	name := "xpy"
	ptr := &name
	fmt.Println("name", *ptr)
	fmt.Println("name ptr", ptr)

	//使用new关键字定义
	name2ptr := new(string)
	*name2ptr = "zhaojinfan"
	fmt.Println("name2ptr", name2ptr)
	fmt.Println("name2ptr value", *name2ptr)

	res := testptr()
	fmt.Println("res ptr: ", *res)

	if res == nil {
		fmt.Println("res 是空")
	} else if res != nil {
		fmt.Println("res 不是空")
	}
}

func testptr() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
