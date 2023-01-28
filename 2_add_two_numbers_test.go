package leetcode_excercise

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  9,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	ret := addTwoNumbers(l1, l2)

	fmt.Println(ret)
}
