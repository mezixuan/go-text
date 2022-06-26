package main

import "fmt"

//定义一个接口，注意类型是interface
type IAttack interface {
	//接口的函数可以有多个，但是只能有函数原型不可以实现
	Attack()
}

type HumanLowLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是 ：", a.name, "等级为 : ", a.level)
}

func DoAttack(a IAttack) {
	a.Attack()
}
func main() {
	var player IAttack
	lowlevel := HumanLowLevel{
		name:  "赵津樊",
		level: 1,
	}
	player = &lowlevel
	lowlevel.Attack()
	player.Attack()
	DoAttack(player)
}
