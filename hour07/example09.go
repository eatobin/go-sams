package main

import (
	"fmt"
	"reflect"
)

type Drink2 struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink2{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink2{
		Name: "Lemonade",
		Ice:  true,
	}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
