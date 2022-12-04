package double

import "testing"

type MockQueryer struct {
	phone         string
	methodsToCall map[string]bool
}

func (mq *MockQueryer) Search(people []*Person, firstName string, lastName string) *Person {
	mq.methodsToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     mq.phone,
	}
}

func (mq *MockQueryer) Create(people []*Person, firstName string, lastName string) *Person {
	mq.methodsToCall["Create"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     mq.phone,
	}
}

func (mq *MockQueryer) ExpectToCall(methodName string) {
	if mq.methodsToCall == nil {
		mq.methodsToCall = make(map[string]bool)
	}
	mq.methodsToCall[methodName] = false
}

func (mq *MockQueryer) Verify(t *testing.T) {
	for methodName, called := range mq.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	expectPhone := "+33 12 333 444"
	phoneBook := &PhoneBook{}
	mock := &MockQueryer{phone: expectPhone}
	mock.ExpectToCall("Search")

	phone, _ := phoneBook.Find(mock, "Kaew", "Ewka")

	if expectPhone != phone {
		t.Errorf("expected %s, but got %s", expectPhone, phone)
	}

	mock.Verify(t)
}
