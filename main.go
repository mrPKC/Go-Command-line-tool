package main

import (
	model "main.go/DB"
	"main.go/handler"
	personService "main.go/service/person"
	personEndpoint "main.go/transport/person"

	relationService "main.go/service/relation"
	relationEndpoint "main.go/transport/relation"
)

func main() {
	personFile, relationFile, relationshipsFile := model.DBInstances()

	personService := personService.NewPersonService(personFile)
	personEndpoint := personEndpoint.NewPersonEndpoint(personService)

	relationService := relationService.NewRelationService(relationFile, relationshipsFile, personService)
	relationEndpoint := relationEndpoint.NewRelationEndpoint(relationService)
	// Initiatig Handler
	handler := handler.NewCommandHandler(personEndpoint, relationEndpoint)
	handler.Handler()
}
