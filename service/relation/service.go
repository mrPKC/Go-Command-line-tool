package relation

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	personService "main.go/service/person"
	"main.go/utils"
)

type RelationService interface {
	AddRelationShip(relationship utils.RelationShip)
	AddRelation(name string)
	GetRelation(relation string) *string
	GetRelations()
	CountRelationShip(personName string, relation string)
	GetAllRelationShips(personName string, relation string)
}

type relationService struct {
	relationFile     *os.File
	relationShipFile *os.File
	personService    personService.PersonService
}

func NewRelationService(relationFile *os.File, relationShipFile *os.File, personService personService.PersonService) RelationService {
	return &relationService{relationFile: relationFile, relationShipFile: relationShipFile, personService: personService}
}

func (s *relationService) AddRelation(name string) {

	isRelationExist := s.GetRelation(name)
	if isRelationExist != nil {
		fmt.Println("Relation already exist with name: ", name)
		return
	}

	writer := csv.NewWriter(s.relationFile)
	defer writer.Flush()

	data := [][]string{
		{name},
	}

	err := writer.WriteAll(data)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	defer s.relationFile.Close()
	fmt.Println("Relation Created Successfully!")
}

func (s *relationService) GetRelation(relation string) *string {
	reader := csv.NewReader(s.relationFile)

	// Read all records from the CSV file
	relations, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	// Process each record
	for _, value := range relations {
		if value[0] == relation {
			return &value[0]
		}
	}

	return nil
}

func (s *relationService) GetRelations() {
	fmt.Println("service Relations")
	reader := csv.NewReader(s.relationFile)

	// Read all records from the CSV file
	relations, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Process each record
	for _, value := range relations {
		utils.PrintJson(value)
	}
}

// As given in test cases CLI commands should be able to create a relation so we can't have functionality of registering relation as son
// ans fetching values as sons and also it is not possible to create vice-versa relation like son of someone sutomatically make other person his father
func (s *relationService) AddRelationShip(relationship utils.RelationShip) {
	isRelationExist := s.GetRelation(relationship.Relation)
	if isRelationExist == nil {
		fmt.Println("Relation does not exist with name: ", relationship.Relation)
		return
	}
	isRelationShipExist := s.GetRelationShip(relationship.FirstPerson, relationship.SecondPerson)
	if isRelationShipExist != nil {
		fmt.Println("Relationship Already exist!")
		return
	}
	isFirstPersonExist := s.personService.GetPerson(relationship.FirstPerson)
	if isFirstPersonExist == nil {
		fmt.Println("Person not found! Name : ", relationship.FirstPerson)
		return
	}
	isSecondPersonExist := s.personService.GetPerson(relationship.SecondPerson)
	if isSecondPersonExist != nil {
		fmt.Println("Person not found! Name : ", relationship.SecondPerson)
		return
	}

	writer := csv.NewWriter(s.relationShipFile)
	defer writer.Flush()

	data := [][]string{
		{relationship.FirstPerson, relationship.Relation, relationship.SecondPerson},
	}

	err := writer.WriteAll(data)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	defer s.relationFile.Close()
	fmt.Println("Relation Created Successfully!")
}

func (s *relationService) GetRelationShip(firstPerson string, secondPerson string) *utils.RelationShip {
	reader := csv.NewReader(s.relationShipFile)

	// Read all records from the CSV file
	relationships, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	// Process each record
	for _, value := range relationships {
		if value[0] == firstPerson && value[2] == secondPerson {
			var relationship utils.RelationShip
			relationship.FirstPerson = firstPerson
			relationship.SecondPerson = secondPerson
			relationship.Relation = value[1]
			return &relationship
		}
	}

	return nil
}

func (s *relationService) CountRelationShip(personName string, relation string) {
	reader := csv.NewReader(s.relationShipFile)

	// Read all records from the CSV file
	relationships, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Process each record
	count := 0
	for _, value := range relationships {
		if value[1] == relation && value[2] == personName {
			count += 1
		}
	}
	fmt.Println(`Number of ` + relation + ` of ` + personName + ` = ` + strconv.Itoa(count) + ``)
}

func (s *relationService) GetAllRelationShips(personName string, relation string) {
	reader := csv.NewReader(s.relationShipFile)

	// Read all records from the CSV file
	relationships, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Process each record
	for _, value := range relationships {
		if value[1] == relation && value[2] == personName {
			utils.PrintJson(value[0])
		}
	}
}
