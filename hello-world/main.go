package main

import "fmt"

const helloPrefix = "Hello, "

// Hello print greet
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
