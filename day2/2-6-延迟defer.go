package main

import (
	"fmt"
	"os"
)

func main() {
	readFile("2-自增.go")
}

func readFile(filename string) {
	f1, err := os.Open(filename)
	defer func() {
		fmt.Println("准备关闭文件")
		_ = f1.Close()
	}()

	if err != nil {
		fmt.Println("打开文件失败err", err)
		return
	}
	buf := make([]byte, 1024)
	n, _ := f1.Read(buf)
	fmt.Println("读取文件的实际长度:", n)
	fmt.Println("读取文件的内容:", string(buf))

}
