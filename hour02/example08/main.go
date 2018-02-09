package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b = true
	fmt.Println(reflect.TypeOf(b))
	var s = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
