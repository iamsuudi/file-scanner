package main

import "fmt"

var config map[string]string

func init() {
	config = map[string]string{
		"env": "development",
	}
	fmt.Println("config loaded") // runs before main
}

func main() {
	fmt.Println(config["env"]) // development
}
