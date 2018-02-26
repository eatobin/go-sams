package main

import "fmt"

type Movie4 struct {
	Name   string
	Rating float32
}

func main() {
	m := new(Movie4)
	m.Name = "Metropolis"
	m.Rating = 0.99
	fmt.Printf("%+v\n", m)
}
