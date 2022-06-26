package main

import (
	"day2/2-8-class_public_private/src"
	"fmt"
)

func main() {
	s1 := src.Student1{
		Hum: src.Human{
			Name:   "赵津樊",
			Age:    20,
			Gender: "男",
		},
		School: "鹤岗一中",
	}
	h1 := src.Human{
		Name:   "xxx",
		Age:    10,
		Gender: "无",
	}
	t1 := src.Teacher{}
	t1.Name = "夏裴媛"
	t1.Gender = "女"
	t1.Subject = "学前教育"
	t1.Eat()
	fmt.Println(s1)
	fmt.Println(t1)
	fmt.Println(h1)
}
