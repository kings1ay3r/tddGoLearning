package main

import (
	"d_i/dependency_injection"
	"log"
	"net/http"
)

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

	log.Fatal(http.ListenAndServe(":5050", http.HandlerFunc(dependency_injection.MyGreeterHandler)))
}
