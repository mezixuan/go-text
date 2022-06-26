package main

import (
	"11-import/add"
	"11-import/sub"
	"fmt"
) //sub是文件夹名更重要的是包名

func main() {
	res := sub.Sub(20, 10)
	ans := add.Add(5, 10)

	fmt.Println(res)
	fmt.Println(ans)
}
