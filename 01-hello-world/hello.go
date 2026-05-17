package main

import (
	"fmt"

	integers "example.com/hello/02-integers"
)

const (
	spanish = "Spanish"
	french  = "French"
	german  = "German"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	germanHelloPrefix  = "Hallo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// greetingPrefix returns the prefix for saying 'Hello' in the given language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(integers.Add(1, 2))
	fmt.Println(Hello("world", "English"))
}
