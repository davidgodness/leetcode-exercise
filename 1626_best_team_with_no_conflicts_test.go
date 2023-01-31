package leetcode_excercise

import (
	"testing"
)

func TestBestTeamWithNoConflicts(t *testing.T) {
	scores := [][]int{
		{1, 2, 3, 5},
		{1, 3, 5, 10, 15},
		{4, 5, 6, 5},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{596, 277, 897, 622, 500, 299, 34, 536, 797, 32, 264, 948, 645, 537, 83, 589, 770},
	}
	ages := [][]int{
		{8, 9, 10, 1},
		{1, 2, 3, 4, 5},
		{2, 1, 2, 1},
		{811, 364, 124, 873, 790, 656, 581, 446, 885, 134},
		{18, 52, 60, 79, 72, 28, 81, 33, 96, 15, 18, 5, 17, 96, 57, 72, 72},
	}

	expected := []int{
		6,
		34,
		16,
		10,
		3287,
	}

	for i := 0; i < len(scores); i++ {
		realValue := bestTeamScore(scores[i], ages[i])
		if realValue != expected[i] {
			t.Errorf("expected %d, but %d\n", expected[i], realValue)
		}
	}
}
