package main

import "fmt"

func test1(a int, b int, c string) (int, string) {
	return a + b, c
}

func test2(a, b int, c string) (res int, str string, x bool) {
	var i, j int
	i = 1
	j = 2
	res = i + j
	str = "test2"
	h := true
	return res, str, h
}

func main() {
	//r1, r2 := test1(1, 2,"nihao")
	res, str, x := test2(0, 1, "a")
	fmt.Println(res, str, x)
	//fmt.Println(r1, r2)
}
