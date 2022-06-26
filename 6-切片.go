package main

import "fmt"

func main() {
	names := []string{"北京", "上海", "广州", "深圳"}
	for i, v := range names {
		fmt.Println("i : ", i, "v : ", v)
	}

	names1 := append(names, "鹤岗")
	fmt.Println("names", names)
	fmt.Println("names1", names1)

	names = append(names, "大庆")
	fmt.Println("names", names)
	fmt.Println("names1", names1)
}
