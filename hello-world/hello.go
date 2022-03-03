package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	helloPrefix := englishHelloPrefix

	switch language {
	case french:
		helloPrefix = frenchHelloPrefix
	case spanish:
		helloPrefix = spanishHelloPrefix
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Luís", "Spanish"))
}
