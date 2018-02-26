package main

import "fmt"

type Movie5 struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie5{
		Name:   "Metropolis",
		Rating: 0.99,
	}
	fmt.Printf("%+v\n", m)
}
