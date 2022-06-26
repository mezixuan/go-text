package main

import "fmt"

const (
	MONDAR           = iota //0
	TUESDAY          = iota //1
	WEDENSDAY        = iota
	THUSEDAY         = iota
	FRIDAY           = iota
	SATURDAY, SUNDAY = iota, iota
)

const (
	FEBRARY = iota + 1
	OCTOBER = iota + 1
)

func main() {
	fmt.Println(MONDAR)
	fmt.Println(THUSEDAY)
	fmt.Println(SATURDAY, SUNDAY)
	fmt.Println(FEBRARY)
	fmt.Println(OCTOBER)
	//var(
	//	nums int
	//	names string
	//	flag bool
	//)

}
