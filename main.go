package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//first way
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex.firstName)

	//second way
	miranda := person{firstName: "Miranda", lastName: "Anderson"}
	fmt.Println(miranda.firstName)
	fmt.Println(alex)
	//prints--> {firstName:Alex lastName:Anderson}
	fmt.Printf("%+v", alex)

	//third way
	var unai person
	unai.firstName = "Unai"
	unai.lastName = "Aanders"
	fmt.Println(unai)

}
