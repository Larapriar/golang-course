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

	// First way of doing pointers
	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateNameLongVersion("Jimmy")
	jim.print()

	// Second way of doing pointers
	jim.updateNameShortVersion("George")
	jim.print()

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

func (p person) print() {
	fmt.Printf("%+v", p)
}

//receiver as pointer and not value
func (p *person) updateNameLongVersion(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) updateNameShortVersion(newFirstName string) {
	(*p).firstName = newFirstName
}
