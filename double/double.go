package double

import (
	"errors"
)

var (
	ErrMissingArgs   = errors.New("FirstName and LastName are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type Queryer interface {
	Search(people []*Person, firstName string, lastName string) *Person
}

type PhoneBook struct {
	People []*Person
}

func (p *PhoneBook) Find(query Queryer, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := query.Search(p.People, firstName, lastName)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
