package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

//在类外绑定方法
func (this *Person) eat() {
	fmt.Println("person is eating")
	this.name = "xiapeiyuan"
}
func (this Person) eat2() {
	fmt.Println("this is eat2")
	this.name = "zhaojinfan"
}

func main() {
	lili := Person{
		name:   "Lily",
		age:    20,
		gender: "女",
		score:  90,
	}
	lili1 := lili
	fmt.Println("lily : ", lili)
	lili.eat()
	fmt.Println(lili)
	fmt.Println(lili1)
	lili1.eat2()
	fmt.Println(lili1)
}
