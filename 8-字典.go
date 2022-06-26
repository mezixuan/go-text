package main

import "fmt"

func main() {
	var idNames map[int]string

	//分配空间
	idNames = make(map[int]string, 10) //建议使用这种方式

	idNames[0] = "zhaojinfan"
	idNames[1] = "xiapeiyuan"

	//for key, value := range idNames{
	//	fmt.Println("key : ", key, "value : ", value)
	//}
	//value, ok := idNames[2]
	//if ok{
	//	fmt.Println("id == 2是存在的, value为 ：",value)
	//}
	delete(idNames, 0)
	for key, value := range idNames {
		fmt.Println("key : ", key, "value : ", value)
	}
	//value1, ok2 := idNames[1]
	//if ok2{
	//	fmt.Println("id == 1是存在的，value1为",value1)
	//}
}
