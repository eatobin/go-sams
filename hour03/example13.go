package main

import (
	"fmt"
)

func showMemoryAddress2(x *int) {
	fmt.Println(x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress2(&i)
}
