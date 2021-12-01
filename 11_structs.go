package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// alternative version
	// firstName, lastName, city, gender string
	// age								 int
}

// Greeting method (value receiver)
// input: a person
// name of method: greet
// return value: string
func (p Person) greet() string {
	return ">> " + p.firstName + " " + p.lastName + "(" + strconv.Itoa(p.age) + ")"
}

// hasBirthday method (pointer receiver)
// did not put return value since changing value with pointer
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
// input: husband's surname
func (p *Person) getMarried(newLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = newLastName
	}
}

func main() {

	person1 := Person{firstName: "Cem", lastName: "Gulec", city: "Ist", gender: "M", age: 23}

	// Alternative definition
	person2 := Person{"Cem", "Gulec", "Ist", "M", 23}

	fmt.Println(person1, person2, person1.firstName)

	// Change feature
	fmt.Println(person1.age)
	person1.age++
	fmt.Println(person1.age)

	// Finally using created methods
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())

	// getmarried method

}
