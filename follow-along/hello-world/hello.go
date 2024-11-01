package main

import "fmt"

const englishHelloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("World"))
}
