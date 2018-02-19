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

	x := []int{2,3,5,7,11}
	fmt.Println(x)
	fmt.Println(x[:1])
	fmt.Println(x[0:1])
	fmt.Println(x[4:])

	y := x[1:3]
	fmt.Println(y)
	fmt.Println(y[0])
	fmt.Println(y[1])
	fmt.Println(y[0:4])
}
