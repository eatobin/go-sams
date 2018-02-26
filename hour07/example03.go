package main

import "fmt"

type Movie3 struct {
	Name   string
	Rating float32
}

func main() {
	var m Movie3
	fmt.Printf("%+v\n", m)
	m.Name = "Metropolis"
	m.Rating = 0.99
	fmt.Printf("%+v\n", m)
}
