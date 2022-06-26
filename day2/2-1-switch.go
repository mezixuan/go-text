package main

import (
	"fmt"
	"os"
)

func main() {
	cmds := os.Args
	for key, cmd := range cmds {
		fmt.Println("key : ", key, ", cmd", cmd, "cmds len", len(cmds))
	}

	if len(cmds) < 2 {
		fmt.Println("请正确输入参数!")
		return
	}

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default called!")
	}

}
