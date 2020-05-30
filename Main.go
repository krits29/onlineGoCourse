package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo  //can also leave it at just contactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex := Person{"Alex", "Anderson"} //creating a new struct

	//alex := person{firstName: "Alex", lastName: "Anderson"} to specify

	// var alex person

	//fmt.Println(alex)

	//alex.lastName = "A" //using dots instead of getter

	//fmt.Printf("%+v", alex)

	jim := Person {
		firstName: "Bob",
		lastName:  "Jones",
		contact: ContactInfo {
			email:   "jim@email.com",
			zipCode: 43563,
		},
	}

	fmt.Printf("%+v", jim)
	fmt.Println()
	jim.print()
	fmt.Println()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
