package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	Eliason := person{
		firstName: "Emmmanuel",
		lastName:  "Eliason-Armstrong",
	}
	Eliason.updateName("Eliarms")
	Eliason.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
