package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person4 struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

func main() {
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := Person4{
		Name:    "",
		Age:     40,
		Hobbies: hobbies,
	}
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
}
