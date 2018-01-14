package main

import (
	"fmt"
)

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string {
	return "Hello There!"
}

func (spanishBot) getGreetings() string {
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}
