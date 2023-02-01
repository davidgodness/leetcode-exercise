package leetcode_excercise

func gcdOfStrings(str1 string, str2 string) string {
	shortStr := shortString(str1, str2)
	divisor := ""
	for i := 1; i <= len(shortStr); i++ {
		prefix := shortStr[0:i]

		if isDivisor(prefix, str1) && isDivisor(prefix, str2) {
			divisor = prefix
		}
	}

	return divisor
}

func shortString(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return str1
	} else {
		return str2
	}
}

func isDivisor(prefix string, base string) bool {
	prefixLen := len(prefix)

	if len(base)%prefixLen != 0 {
		return false
	}

	for i := 0; i+prefixLen <= len(base); i += prefixLen {
		if base[i:i+prefixLen] != prefix {
			return false
		}
	}

	return true
}
