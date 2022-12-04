package sum

func sum(nums ...int) int {
	sum := 0

	for _, x := range nums {
		sum += x
	}

	return sum
}
