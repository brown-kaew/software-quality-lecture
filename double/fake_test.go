package double

import "testing"

type FakeQueryer struct{}

func (sq *FakeQueryer) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}
	return people[0]
}

func TestFindCallsSearchAndReturnEmptyStrinngForNoPerson(t *testing.T) {
	expectPhone := ""
	phoneBook := &PhoneBook{}

	phone, _ := phoneBook.Find(&FakeQueryer{}, "Kaew", "Ewka")

	if expectPhone != phone {
		t.Errorf("expected %s, but got %s", expectPhone, phone)
	}
}
