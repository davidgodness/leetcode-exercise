package leetcode_excercise

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !isPalindrome(121) {
		t.Fail()
	}
	if isPalindrome(-121) {
		t.Fail()
	}
	if isPalindrome(10) {
		t.Fail()
	}
}
