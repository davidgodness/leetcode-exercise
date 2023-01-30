package leetcode_excercise

func tribonacci(n int) int {
	hold := [3]int{0, 1, 1}

	if n < 3 {
		return hold[n]
	}

	for i := 3; i < n; i++ {
		sum := hold[0] + hold[1] + hold[2]
		hold[0], hold[1] = hold[1], hold[2]
		hold[2] = sum
	}

	return hold[0] + hold[1] + hold[2]
}
