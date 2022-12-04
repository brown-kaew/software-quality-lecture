package double

import "testing"

type DummyQueryer struct{}

func (dq *DummyQueryer) Search(people []*Person, firstName string, lastName string) *Person {
	return nil
}

func TestFindItShouldReturnErrorWhenFirstNameOrLastNameEmpty(t *testing.T) {
	phoneBook := &PhoneBook{}
	want := ErrMissingArgs

	_, err := phoneBook.Find(&DummyQueryer{}, "", "")

	if want != err {
		t.Errorf("Want %s, but got %s", want, err)
	}
}
