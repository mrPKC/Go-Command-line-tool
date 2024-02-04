package relation

import (
	"fmt"

	relationService "main.go/service/relation"
	"main.go/utils"
)

type RelationEndpoint interface {
	AddRelation(args []string)
	GetRelation(args []string)
	GetRelations()
	CreateRelationShip(args []string)
	CountRelationShip(args []string)
	GetAllRelationShips(args []string)
}

type relationEndpoint struct {
	relationService relationService.RelationService
}

func NewRelationEndpoint(relationService relationService.RelationService) RelationEndpoint {
	return &relationEndpoint{relationService: relationService}
}

func (e *relationEndpoint) AddRelation(args []string) {
	if len(args) != 1 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `add relationship <name>`")
		return
	}

	name := args[0]

	e.relationService.AddRelation(name)
}

func (e *relationEndpoint) GetRelation(args []string) {
	if len(args) != 1 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `get relationship <name>`")
		return
	}

	name := args[0]

	relation := e.relationService.GetRelation(name)
	if relation == nil {
		fmt.Println("Relation not found!")
		return
	}
	utils.PrintJson(relation)
}

func (e *relationEndpoint) GetRelations() {

	e.relationService.GetRelations()
}

func (e *relationEndpoint) CreateRelationShip(args []string) {
	if len(args) != 5 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `connect <name 1> as <relationship> of <name 2>`")
		return
	}
	var relationShip utils.RelationShip
	relationShip.FirstPerson = args[0]
	relationShip.Relation = args[2]
	relationShip.SecondPerson = args[4]

	e.relationService.AddRelationShip(relationShip)
}

func (e *relationEndpoint) CountRelationShip(args []string) {
	if len(args) != 3 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `connect <name 1> as <relationship> of <name 2>`")
		return
	}

	e.relationService.CountRelationShip(args[2], args[0])
}
func (e *relationEndpoint) GetAllRelationShips(args []string) {
	if len(args) != 3 {
		fmt.Println("Command syntax not matched! \n Please enter in the desired format `connect <name 1> as <relationship> of <name 2>`")
		return
	}

	e.relationService.GetAllRelationShips(args[2], args[0])
}
