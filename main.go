package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	fmt.Printf(alex.firstName)
	miranda := person{firstName: "Miranda", lastName: "Anderson"}
	fmt.Printf(miranda.firstName)

}
