package main

import (
	"fmt"
	"time"
)

func slowFunc2() {
	time.Sleep(time.Second * 2)
	fmt.Println("slowFunc2() finished")
}

func main() {
	go slowFunc2()
	fmt.Println("I am now shown straightaway!")
}
