package main

import "fmt"

func main() {
	name := "zhaojinfan"
	l1 := len(name)
	fmt.Println(l1)

	for i := 0; i < l1; i++ {
		fmt.Printf("i : %d, v : %c\n", i, name[i])
	}

	x, y := "hello", "world"
	fmt.Println("x + y = ", x+y)
}
