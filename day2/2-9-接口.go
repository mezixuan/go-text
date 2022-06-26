package main

import "fmt"

func main() {
	var i interface{}
	names := []string{"zhaojinfan", "xiapeiyuan"}
	i = names
	fmt.Println(i)
	kvalue, ok := i.([]string)
	if !ok {
		fmt.Println("不是切片类型")
	} else {
		fmt.Println("是切片类型：", kvalue)
	}
	arr := make([]interface{}, 3)
	arr[0] = 20
	arr[1] = "zjf"
	arr[2] = "男"
	for key, value := range arr {
		//fmt.Println("key", key, "value", value)
		switch v := value.(type) {
		case int:
			fmt.Println("int类型，内容为:", key, v)
		case string:
			fmt.Println("string类型，内容为", key, v)
		default:
			fmt.Println("不是合理的输入类型")
		}

	}
}
