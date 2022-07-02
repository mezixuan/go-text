package main

import (
	"fmt"
)

func main(){
		go func (){
			func(){
				fmt.Println("这是子go程里面的函数")
				return
			}()
			fmt.Println("子go程执行结束")

		}()
		fmt.Println("这是主go程")
		fmt.Println("OVER!")
	}
