package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var file []byte
	var err error
	file, err = ioutil.ReadFile("hour10/foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", file)
}
