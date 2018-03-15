package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("hour21/example06/deleteme.txt")
	if err != nil {
		log.Fatal(err)
	}
}
