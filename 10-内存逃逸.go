package main

import "fmt"

func main() {
	res := testptr1()
	fmt.Println(*res)
	*res = "大庆"
	fmt.Println(*res)
	ans := testptr1()
	fmt.Println(*ans)
}

func testptr1() *string {
	city := "上海"
	ptr := &city
	return ptr
}
