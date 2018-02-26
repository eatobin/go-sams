package main

import "fmt"

type Movie2 struct {
	Name   string
	Rating float32
}

func main() {
	var m Movie2
	fmt.Printf("%+v\n", m)
}
