package main

import (
	"fmt"
)

func addition(x int, y int) int {
	return x + y
}

func main() {
	//var s string = "three"
	var s = 3
	fmt.Println(addition(2, s))
}
