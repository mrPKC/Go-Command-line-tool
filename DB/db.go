package model

import (
	"fmt"
	"os"
)

// Using CSV files as db strage for this app
func DBInstances() (*os.File, *os.File, *os.File) {

	personFile, err := os.OpenFile("DB/person.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, nil
	}

	relationFile, err := os.OpenFile("DB/relations.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, nil
	}

	relationshipsFile, err := os.OpenFile("DB/relationship.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, nil
	}

	return personFile, relationFile, relationshipsFile
}
