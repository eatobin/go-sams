package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "Hello World & Eric"
	err := ioutil.WriteFile("hour21/example03.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
