package main

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	Eliason := person{
		firstName: "Emmmanuel",
		lastName:  "Eliason-Armstrong",
		contactInfo: contactInfo{
			email:   "armstrongemmanuel@yahoo.fr",
			zipcode: 123,
		},
	}
}
fmt.Printf("%+v",Eliason)
