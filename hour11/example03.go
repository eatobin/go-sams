package main

import (
	"fmt"
	"time"
)

func slowFunc3() {
	time.Sleep(time.Second * 2)
	fmt.Println("slowFunc2() finished")
}

func main() {
	go slowFunc3()
	fmt.Println("I am now shown straightaway!")
	time.Sleep(time.Second * 3)
}
