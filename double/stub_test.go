package double

import "testing"

type StubQueryer struct {
	Phone string
}

func (sq *StubQueryer) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     sq.Phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	expectPhone := "+33 12 333 444"
	phoneBook := &PhoneBook{}

	phone, _ := phoneBook.Find(&StubQueryer{Phone: expectPhone}, "Kaew", "Ewka")

	if expectPhone != phone {
		t.Errorf("expected %s, but got %s", expectPhone, phone)
	}
}
