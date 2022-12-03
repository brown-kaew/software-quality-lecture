package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return 3 when 1 and 2", func(t *testing.T) {
		want := 3

		got := sum(1, 2)

		if want != got {
			t.Errorf("sum(1, 2) = %d; want %d", got, want)
		}
	})

	t.Run("should return -2 when -1 and -1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if want != got {
			t.Errorf("sum(-1, -1) = %d; want %d", got, want)
		}
	})

	t.Run("should return 6 when  1, 2, 3", func(t *testing.T) {
		want := 6

		got := sum(1, 2, 3)

		if want != got {
			t.Errorf("sum(1, 2, 3) = %d; want %d", got, want)
		}
	})

	t.Run("should return 15 when  1, 2, 3, 4, 5", func(t *testing.T) {
		want := 15
		nums := []int{1, 2, 3, 4, 5}

		got := sum(nums...)

		if want != got {
			t.Errorf("sum(1, 2, 3, 4, 5) = %d; want %d", got, want)
		}
	})
}
