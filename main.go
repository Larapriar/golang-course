package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

//second way of embebing structs
type secondPerson struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	//first way
	alexInfo := contactInfo{"alex.anderson@email.com", 33430}
	alex := person{"Alex", "Anderson", alexInfo}
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

	//embebing structs in structs
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@email.com",
			zipCode: 33430,
		},
	}
	fmt.Printf("%+v", jim)

	Miumiu := secondPerson{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim.party@email.com",
			zipCode: 33430,
		},
	}
	fmt.Printf("%+v", Miumiu)

}
