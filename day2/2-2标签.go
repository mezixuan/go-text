package main

import "fmt"

func main() {

LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL1
				continue LABEL1
			}
			fmt.Println("i", i, "j", j)
		}
	}

	fmt.Println("over！")
}
