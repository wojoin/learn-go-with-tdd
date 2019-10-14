package main

import "fmt"

const helloPrefix = "Hello, "

// Hello print greet
func Hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
