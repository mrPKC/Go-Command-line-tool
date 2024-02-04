package utils

type Person struct {
	Name         string `json:"name"`
	Gender       Gender `json:"gender"`
	DateOfBirth  string `json:"date_of_birth"`
	DateOfDemise string `json:"date_of_demise,omitempty"`
}

type RelationShip struct {
	FirstPerson  string `json:"first_person"`
	SecondPerson string `json:"second_person"`
	Relation     string `json:"relation"`
}
