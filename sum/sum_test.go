package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return 3 when 1 and 2", func(t *testing.T) {
		want := 3

		got := sum(1, 2)

		if want != got {
			t.Errorf("sum(1, 2) = %d; wannt 3", got)
		}
	})

	t.Run("should return -2 when -1 and -1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if want != got {
			t.Errorf("sum(-1, -1) = %d; wannt -2", got)
		}
	})
}
