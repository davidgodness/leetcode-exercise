package leetcode_excercise

import (
	"testing"
)

func TestLongestSubstringLength(t *testing.T) {
	expected := make(map[string]int)

	expected["abcabcbb"] = 3
	expected["bbbbb"] = 1
	expected["pwwkew"] = 3
	expected["dvdf"] = 3

	for k, v := range expected {
		if v != lengthOfLongestSubstring(k) {
			t.Errorf("expected %d, but %d\n", v, lengthOfLongestSubstring(k))
		}
	}
}
