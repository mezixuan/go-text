package main

import "fmt"

type Myint int

type Student struct {
	id     int
	name   string
	gender string
}

func main() {
	//var x Myint
	//x = 10
	lili := Student{
		id:     12,
		name:   "zhaojinfan",
		gender: "女",
	}

	zhaojinfan := Student{
		id:     22,
		gender: "男",
	}
	fmt.Println(zhaojinfan)
	fmt.Println(lili.name, lili.id, lili.gender)
}
