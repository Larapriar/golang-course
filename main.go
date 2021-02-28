package main

import "fmt"

type bot interface {
	getGreeting() string
}

type bot2 interface {
	getBye() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printBye(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi there!"
}

func (englishBot) getBye() string {
	//very custom logic for generating an english greeting
	return "Bye"
}

func (spanishBot) getGreeting() string {
	//very custom logic for generating an spanish greeting
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printBye(b bot2) {
	fmt.Println(b.getBye())
}
