package ticket

import "testing"

func TestPrice(t *testing.T) {

	t.Run("should return 0 when age is 0", func(t *testing.T) {
		want := 0.0
		var age uint = 0

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})

	t.Run("should return 0 when age is 3", func(t *testing.T) {
		want := 0.0
		var age uint = 3

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})

	t.Run("should return 15 when age is 4", func(t *testing.T) {
		want := 15.0
		var age uint = 4

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})

	t.Run("should return 15 when age is 15", func(t *testing.T) {
		want := 15.0
		var age uint = 15

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})

	t.Run("should return 30 when age is 16", func(t *testing.T) {
		want := 30.0
		var age uint = 16

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})

	t.Run("should return 30 when age is 50", func(t *testing.T) {
		want := 30.0
		var age uint = 50

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})

	t.Run("should return 5 when age is 51", func(t *testing.T) {
		want := 5.0
		var age uint = 51

		got := price(age)

		if want != got {
			t.Errorf("Want %f but got %f", want, got)
		}
	})
}
