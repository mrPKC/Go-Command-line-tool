package person

import (
	"fmt"
	"time"

	personService "main.go/service/person"
	"main.go/utils"
)

type PersonEndpoint interface {
	AddPerson(args []string)
	GetPerson(args []string)
	GetPersons()
}

type personEndpoint struct {
	personService personService.PersonService
}

func NewPersonEndpoint(personService personService.PersonService) PersonEndpoint {
	return &personEndpoint{personService: personService}
}

func (e *personEndpoint) AddPerson(args []string) {
	if len(args) != 3 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `add person <name> <gender> <date of birth>`")
		return
	}

	name := args[0]
	gender := args[1]

	dob := args[2]
	if name == "" || gender == "" {
		fmt.Println("Please provide gender and name", name, gender)
		return
	}
	_, err := time.Parse("2006-01-02", dob)
	if err != nil {
		fmt.Println("Please Enter the --DOB in this 2006-01-02 format.")
		return
	}

	var person utils.Person
	person.Name = name
	person.Gender = utils.Gender(gender)
	person.DateOfBirth = dob

	e.personService.AddPerson(person)
}

func (e *personEndpoint) GetPerson(args []string) {
	if len(args) != 1 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `get person <name>`")
		return
	}

	name := args[0]

	person := e.personService.GetPerson(name)
	if person == nil {
		fmt.Println("Person not found!")
		return
	}
	utils.PrintJson(person)
}

func (e *personEndpoint) GetPersons() {

	e.personService.GetPersons()
}
