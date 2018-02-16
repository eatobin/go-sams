package main

import "fmt"

func main() {
	var cheeses = make([] string, 3)
	cheeses[0] = "Blue"
	cheeses[1] = "Yellow"
	cheeses[2] = "White"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	var cheeses2 = append(cheeses[:0], cheeses[:1]...)
	fmt.Println(len(cheeses2))
	fmt.Println(cheeses2)
	var cheeses3 = append(cheeses[:1], cheeses[1+1:]...)
	fmt.Println(len(cheeses3))
	fmt.Println(cheeses3)
	var cheeses4 = append(cheeses[:0], cheeses[2:]...)
	fmt.Println(len(cheeses4))
	fmt.Println(cheeses4)
}
