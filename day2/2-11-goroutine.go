package main

import (
	"fmt"
	"time"
)

//这个将用于子go程使用
func display(num int) {
	count := 1
	for {
		fmt.Println("这是子go程 : ", num, "当前count值", count)
		count++
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	for i := 0; i < 3; i++{
		go display(i)
	}
	//主go程
	count := 1
	for {
		fmt.Println("这是主go程 : ", count)
		count++
		time.Sleep(1 * time.Second)
	}

	//启动子go程
}
