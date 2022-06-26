package main

import "fmt"

func main() {
	names := []string{"北京", "上海", "广州", "深圳", "大庆"}

	//想基于names创建一个新的数组
	//切片可以基于一个数组，灵活的创建新的数组
	names2 := names[0:3] //左闭右开
	fmt.Println(names2)
	names2[0] = "鹤岗"
	fmt.Println(names)
	fmt.Println(names2)

	sub1 := "hellowold"[6:8]
	fmt.Println(sub1)

	names1 := make([]string, len(names))
	copy(names1, names)
	names1[1] = "香港"
	fmt.Println(names1)
	fmt.Println(names)
}
