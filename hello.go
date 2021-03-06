package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func main() {
	fmt.Println(Hello("Neil", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
