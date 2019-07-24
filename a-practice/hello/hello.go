package main

import "fmt"

const enhelloPrefix = "Hello, "
const esHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

const frenchLanguageCode = "fr"
const spanishLanguageCode = "es"

func SayHello(name, language string) string {

	if language == "fr" {
		return frHelloPrefix + name

	} else if language == "es" {

		return esHelloPrefix + name
	}
	return enhelloPrefix + name

}

//this is the entry point to our program
func main() {
	fmt.Println(SayHello("World", "en"))
}
