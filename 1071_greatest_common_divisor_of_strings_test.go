package leetcode_excercise

import "testing"

func TestGcdOfStrings(t *testing.T) {
	params := []struct {
		str1     string
		str2     string
		expected string
	}{
		{str1: "ABCABC", str2: "ABC", expected: "ABC"},
		{str1: "ABABAB", str2: "ABAB", expected: "AB"},
		{str1: "LEET", str2: "CODE", expected: ""},
		{str1: "AAA", str2: "AA", expected: "A"},
	}

	for _, param := range params {
		if gcdOfStrings(param.str1, param.str2) != param.expected {
			t.Fail()
		}
	}
}
