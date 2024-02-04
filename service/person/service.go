package person

import (
	"encoding/csv"
	"fmt"
	"os"

	"main.go/utils"
)

type PersonService interface {
	AddPerson(person utils.Person)
	GetPerson(name string) *utils.Person
	GetPersons()
}

type personService struct {
	personFile *os.File
}

func NewPersonService(personFile *os.File) PersonService {
	return &personService{personFile: personFile}
}

// As there is not any parameter is defined for uniquly identify a person so I'm taking "name" as that parameter
func (s *personService) AddPerson(person utils.Person) {

	isPersonExist := s.GetPerson(person.Name)
	if isPersonExist != nil {
		fmt.Println("Person already exist with name: ", person.Name)
		return
	}

	writer := csv.NewWriter(s.personFile)
	defer writer.Flush()

	data := [][]string{
		{person.Name, string(person.Gender), person.DateOfBirth},
	}

	err := writer.WriteAll(data)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	defer s.personFile.Close()
	fmt.Println("Person Created Successfully!")
}

func (s *personService) GetPerson(name string) *utils.Person {
	reader := csv.NewReader(s.personFile)

	// Read all records from the CSV file
	persons, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	// Process each record
	for _, value := range persons {
		if value[0] == name {
			var person utils.Person
			person.Name = value[0]
			person.Gender = utils.Gender(value[1])
			person.DateOfBirth = value[2]

			return &person
		}
	}
	return nil
}

func (s *personService) GetPersons() {
	reader := csv.NewReader(s.personFile)

	// Read all records from the CSV file
	persons, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Process each record
	for _, value := range persons {
		var person utils.Person
		person.Name = value[0]
		person.Gender = utils.Gender(value[1])
		person.DateOfBirth = value[2]

		utils.PrintJson(person)
	}
}
