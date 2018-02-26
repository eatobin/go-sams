package main

import (
	"fmt"
)

type Triangle4 struct {
	base   float64
	height float64
}

func (t Triangle4) changeBase(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle4{base: 3, height: 1}
	t.changeBase(4)
	fmt.Println(t.base)
}
