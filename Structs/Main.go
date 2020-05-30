package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"} //creating a new struct

	//alex := person{firstName: "Alex", lastName: "Anderson"} to specify

	// var alex person

	fmt.Println(alex)

	alex.lastName = "A" //using dots instead of getter

	fmt.Printf("%+v", alex)
}
