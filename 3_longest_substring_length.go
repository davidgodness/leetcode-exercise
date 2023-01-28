package leetcode_excercise

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	longestSubStr := ""

	for i := 0; i < length; i++ {
		j := i + 1
	substr:
		for ; j < length; j++ {
			for m := i; m < j; m++ {
				if s[m] == s[j] {
					break substr
				}
			}
		}

		if len(s[i:j]) > len(longestSubStr) {
			longestSubStr = s[i:j]
		}
	}

	return len(longestSubStr)
}
