package ticket

import "testing"

func TestPrice(t *testing.T) {

	tests := []struct {
		name string
		age  uint
		want float64
	}{
		{name: "should return 0 when age is 0", age: 0, want: 0.0},
		{"should return 0 when age is 3", 3, 0.0},
		{"should return 15 when age is 4", 4, 15.0},
		{"should return 15 when age is 15", 15, 15.0},
		{"should return 30 when age is 16", 16, 30.0},
		{"should return 30 when age is 50", 50, 30.0},
		{"should return 5 when age is 51", 51, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			var age uint = tt.age

			got := price(age)

			if want != got {
				t.Errorf("Want %f but got %f", want, got)
			}
		})
	}
}
