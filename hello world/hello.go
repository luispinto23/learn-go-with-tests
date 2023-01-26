package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Luís"))
}
func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name
}
