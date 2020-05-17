package model

func (person *Person) CheckID(personID *string) bool {
	return person.ID == *personID
}