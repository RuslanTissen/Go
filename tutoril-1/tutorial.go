package main

import "fmt"

type contact struct {
	firstName   string
	lastName    string
	phoneNumber string
	email       string
	address     string
	dateOfBirth string
}

func (c contact) printInfo() {
	fmt.Printf("Name: %s\nSirname: %s\nPhone: %s\nE-mail: %s\nAddress: %s\nBirthday: %s\n",
		c.firstName, c.lastName, c.phoneNumber, c.email, c.address, c.dateOfBirth)
}

func (c *contact) setName(name string) {
	c.firstName = name
}

func main() {

	c1 := contact{
		firstName:   "Peter",
		lastName:    "Ganz",
		phoneNumber: "656546584",
		email:       "kjbvdk@gmail.com",
		address:     "Fuhrenkamp 32, 29313 Hambühren",
		dateOfBirth: "23.11.1973",
	}

	c1.setName("Carl")

	c1.printInfo()

	// contactList := map[string]string{
	// 	"Ben":   "65479867970",
	// 	"Smith": "146586797859875",
	// 	"Peter": "6545646848",
	// }

	// fmt.Println(contactList["Ben"])

	// for name, phoneNumber := range contactList {
	// 	fmt.Println(name, phoneNumber)
	// }

	// delete(contactList, "Smith")

	// for name, phoneNumber := range contactList {
	// 	fmt.Println(name, phoneNumber)
	// }

	// contactList := []string{"Ben", "Smith", "Peter"}

	// fmt.Println("List of contacts:")
	// for index, value := range contactList {
	// 	fmt.Printf("%d. %s\n", index+1, value)
	// }

	// contactList = append(contactList, "Frank", "Henrik", "Jon")

	// for i := 0; i < len(contactList); i++ {
	// 	fmt.Println(i, contactList[i])
	// }
}
