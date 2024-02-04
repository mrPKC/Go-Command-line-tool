package handler

import (
	"os"

	personEndpoint "main.go/transport/person"
	relationEndpoint "main.go/transport/relation"
)

type CommandHandler interface {
	Handler()
}

type commandHandler struct {
	personEndpoint   personEndpoint.PersonEndpoint
	relationEndpoint relationEndpoint.RelationEndpoint
}

func NewCommandHandler(personEndpoint personEndpoint.PersonEndpoint, relationEndpoint relationEndpoint.RelationEndpoint) CommandHandler {
	return &commandHandler{personEndpoint: personEndpoint, relationEndpoint: relationEndpoint}
}

func (h *commandHandler) Handler() {

	if len(os.Args) >= 3 {
		subCommand := os.Args[1]
		switch subCommand {
		case "add":
			addHandler(h, os.Args[2:])
		case "get":
			getHandler(h, os.Args[2:])
		case "connect":
			connectHandler(h, os.Args[2:])
		case "count":
			countHandler(h, os.Args[2:])
		default:
			h.relationEndpoint.GetAllRelationShips(os.Args[1:])
		}
	}
}

func connectHandler(h *commandHandler, args []string) {
	h.relationEndpoint.CreateRelationShip(args)
}
func countHandler(h *commandHandler, args []string) {
	h.relationEndpoint.CountRelationShip(args)
}

func getHandler(h *commandHandler, args []string) {

	parentCommand := args[0]
	subCommand := args[1]
	switch parentCommand {
	case "person":
		switch subCommand {
		case "list":
			h.personEndpoint.GetPersons()
		default:
			h.personEndpoint.GetPerson(args[1:])
		}
	case "relationship":
		switch subCommand {
		case "list":
			h.relationEndpoint.GetRelations()
		default:
			h.relationEndpoint.GetRelation(args[1:])
		}
	}

}

func addHandler(h *commandHandler, args []string) {

	subCommand := args[0]
	switch subCommand {
	case "person":
		h.personEndpoint.AddPerson(args[1:])
	case "relationship":
		h.relationEndpoint.AddRelation(args[1:])
	}
}
