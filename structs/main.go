package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//var alex person
	//fmt.Println(alex) //firstName && lastName == ""
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()

}

func (p *person) updateName(newFirstName string) { //*person -> use pointer (&address) of a person type
	(*p).firstName = newFirstName //(*p) -> value of the pointer, that is, p would simply reference the &address in ram
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
