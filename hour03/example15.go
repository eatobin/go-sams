package main

import (
	"fmt"
)

const greeting2 string = "Hello, world"

func main() {
	greeting2 = "Goodbye, cruel world"
	fmt.Println(greeting2)
}
