package main

import "fmt"

type Drink4 struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink4{
		Name: "Lemonade",
		Ice:  true,
	}
	b := &a
	b.Ice = false
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)
}
