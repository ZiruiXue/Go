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

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 90007,
		},
	}

	jim.updateName("Jimmy")
	jim.print()


	mySlice := []string{"Hi","There","How","are","you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

// pass by value
// use pointer

// &variable - Give me the memory address of the value this variable is pointing at
// *pointer - Gime me the value this memory address is pointing at

// Turn address into value with *address
// Turn value into address with &value

// *person - This is a type description - it means we are working with a pointer to a person.
// *pointerToPerson - This is an operator - it means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (s []string)updateSlice {
	s[0] = "Byt"
}
