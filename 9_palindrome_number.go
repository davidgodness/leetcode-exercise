package leetcode_excercise

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	const capacity = 10

	hold := make([]int, 0, capacity)

	for i, n := 0, x; i < capacity; i++ {
		hold = append(hold, n%10)
		if n/10 == 0 {
			break
		}
		n /= 10
	}

	length := len(hold)
	sum := 0
	for _, v := range hold {
		length--
		sum += v * pow10(length)
	}

	if sum == x {
		return true
	}

	return false
}

func pow10(n int) int {
	ret := 1
	for n > 0 {
		ret *= 10
		n--
	}

	return ret
}
