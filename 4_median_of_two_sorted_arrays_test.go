package leetcode_excercise

import (
	"fmt"
	"testing"
)

func TestFindMedian(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}

	median := findMedianSortedArrays(input1, input2)

	fmt.Printf("%.2f\n", median)

	if fmt.Sprintf("%.2f", median) != "2.50" {
		t.Fail()
	}
}
