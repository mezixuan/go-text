package sub

import "fmt"

func init() {
	fmt.Println("this is first package sub init")
}
func Sub(x, y int) int {
	return x - y
}
