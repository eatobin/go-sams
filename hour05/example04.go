package main

import "fmt"

func main() {
	i := 2
	if i == 3 {
		fmt.Println("i is 3")
	} else if i == 9 {
		fmt.Println("i is 9")
	} else {
		fmt.Print("i is 2!")
	}
}
