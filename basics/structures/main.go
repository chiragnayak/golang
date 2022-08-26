package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	phone string
}

type citizenshipInfo struct {
	country string
	status  string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // onlt mentioning "contactInfo" is enough
	citizenshipInfo
}

func main() {

	citizenInfo := citizenshipInfo{"Alien", "Invalid"}
	fmt.Println(citizenInfo)
	fmt.Println("%+v", citizenInfo)

	contactInfo_1 := contactInfo{email: "xxxxx@xxx.com", phone: "12323422342"}
	fmt.Println(contactInfo_1)

	person_1 := person{"Chacha", "Chaudhary", contactInfo_1, citizenInfo}
	fmt.Println("%+v", person_1)

	var person_3 person
	person_3.firstName = "Lorum"
	person_3.lastName = "Ipsum"
	person_3.print()

	jim := person{
		firstName: "Dolce",
		lastName:  "Gabanna",
		contact: contactInfo{
			email: "dgabanna@gmail.com",
			phone: "98474712343",
		},
		citizenshipInfo: citizenshipInfo{
			country: "India",
			status:  "valid",
		},
	}

	jim.print()

	/*

		&variable >>> Give me the memory address of the value this variable is pointing at
		*pointer  >>> Give me the value this memory address pointing at

		suppose below is some memory location :
		[---address-----][-----value------]

		Turn address into a value >>> *address
		Turn value into address >>> &value

	*/
	jim_pointer := &jim
	fmt.Println("&jim >>> ", &jim)               // expected to print memory address, but printes object values itself. Feature of fmt ??
	fmt.Println("jim_pointer >>> ", jim_pointer) // expected to print memory address, but printes object values itself. Feature of fmt ??
	jim_pointer.updateName("jimmy bro")

	jim.print()
}

// *person >>> variable that can receive a memory address. Basically, pointer to a type (here it is "person")
func (pointerToPersion *person) updateName(newFirstName string) {
	//(*pointerToPerson)  >>> this is to manipulate the value to address where pointer is referencing
	//removing () gives error
	(*pointerToPersion).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("%+v", p)
}
