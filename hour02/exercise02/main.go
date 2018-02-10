package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "42"
	if i, err := strconv.Atoi(s); err == nil {
		fmt.Printf("%T, %v\n", i, i)
		i += 1
		fmt.Printf("%T, %v\n", i, i)
		ns := strconv.Itoa(i)
		fmt.Printf("%T, %v\n", ns, ns)
	}
	if fl, err := strconv.ParseFloat(s, 64); err == nil {
		fmt.Printf("%T, %v\n", fl, fl)
	}
	fmt.Printf("%T, %v\n", 61, 61)
	fmt.Println(61)
}
