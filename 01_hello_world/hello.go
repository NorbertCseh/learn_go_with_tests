package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	spanish   = "Spanish"
	french    = "French"
	hungarian = "Hungarian"

	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefix    = "Bonjour, "
	hungarianHelloPrefix = "Szia, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hungarian:
		prefix = hungarianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
