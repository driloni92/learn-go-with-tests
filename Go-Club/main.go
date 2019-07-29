package main

const name = "Bob"

func greetWorld() string {
	return "Hello, world"
}

func greet(name string) string {
	if name == "" {
		return "Hello stranger"
	}
		return "Hello" + name	
}