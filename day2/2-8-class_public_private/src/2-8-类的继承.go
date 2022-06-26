package src

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (this *Human) Eat() {
	fmt.Println("this is : ", this.Name)
}

//定义一个学生类,去嵌套Human但是不是继承
type Student1 struct {
	Hum    Human //包含Humen类型的变量
	School string
	Score  float64
}

type Teacher struct {
	Human   //直接写Human类型，没有字段名字
	Subject string
}
