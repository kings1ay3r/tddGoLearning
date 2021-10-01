package main

import "fmt"

func Hello(name string, language string) string {
	var helloPrefix string
	if name == "" {
		name = "World"
	}
	if language == "es" {
		helloPrefix = "Hola, "
	} else if language == "fr" {
		helloPrefix = "Bonjour, "
	} else {
		helloPrefix = "Hello, "
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Kevin", ""))
}
