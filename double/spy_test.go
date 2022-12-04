package double

import "testing"

type SpyQueryer struct {
	phone           string
	searchWasCalled bool
	whatIsFirstName string
}

func (sq *SpyQueryer) Search(people []*Person, firstName string, lastName string) *Person {
	sq.searchWasCalled = true
	sq.whatIsFirstName = firstName
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     sq.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	expectPhone := "+33 12 333 444"
	expectFirstName := "Kaew"
	phoneBook := &PhoneBook{}
	spy := &SpyQueryer{phone: expectPhone}

	phone, _ := phoneBook.Find(spy, expectFirstName, "Ewka")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if expectFirstName != spy.whatIsFirstName {
		t.Errorf("expected %s, but got %s", expectFirstName, spy.whatIsFirstName)
	}

	if expectPhone != phone {
		t.Errorf("expected %s, but got %s", expectPhone, phone)
	}
}
