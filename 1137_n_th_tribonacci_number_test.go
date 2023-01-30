package leetcode_excercise

import (
	"fmt"
	"testing"
)

func TestTribonacci(t *testing.T) {
	expected := tribonacci(25)

	if expected != 1389537 {
		fmt.Println(expected)
		t.Fail()
	}
}
