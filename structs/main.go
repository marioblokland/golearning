package main

import "fmt"

// A struct would be something like a class in other languages
type Person struct {
	// These attributes are called fields in Go
	firstName string
	lastName  string
	age       int
}

type badPerson struct {
	Person
	firstName string
}

func (p Person) fullName() string {
	return p.firstName + "" + p.lastName
}

func main() {

	// Creating 2 person structs
	person := Person{"James", "Bond", 50}
	badPerson := badPerson{
		Person: Person{"Bad Clark", "Kent", 30},
	}

	fmt.Println(person.firstName, person.lastName, person.age)

	// If a struct has a field with the same name as an embedded struct, then that outer struct's field will be
	// the one which can be accessed with "person.first". To get to the inner struct's field, we need to use
	// "badPerson.Person.firstName", whereas there is no "lastName" field in the "BadPerson" struct, so the inner
	// struct's "lastName" field is promoted and accessible via "badPerson.lastName" instead of "badPerson.Person.lastName".
	fmt.Println(badPerson.Person.firstName, badPerson.lastName, badPerson.age)

}
