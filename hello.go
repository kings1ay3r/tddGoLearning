package main

import "fmt"

func Hello(name string, language string) string {
	var helloPrefix string
	if name == "" {
		name = "World"
	}
	switch language {
	case "es":
		helloPrefix = "Hola, "
	case "fr":
		helloPrefix = "Bonjour, "
	default:
		helloPrefix = "Hello, "
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Kevin", ""))
}
