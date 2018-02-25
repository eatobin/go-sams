package main

import "fmt"

func addition2(x int, y int) int {
	return x + y
}

func main() {
	var s string = "three"
	fmt.Println(addition2(1, s))
}
