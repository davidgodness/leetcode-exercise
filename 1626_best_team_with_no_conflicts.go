package leetcode_excercise

import (
	"sort"
)

type Pair struct {
	score int
	age   int
}

func initPairs(scores []int, ages []int) []Pair {
	pairs := make([]Pair, 0, len(scores))

	for i := 0; i < len(scores); i++ {
		pairs = append(pairs, Pair{
			score: scores[i],
			age:   ages[i],
		})
	}

	return pairs
}

func bestTeamScore(scores []int, ages []int) int {
	pairs := initPairs(scores, ages)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].age < pairs[j].age {
			return true
		} else if pairs[i].age == pairs[j].age {
			return pairs[i].score < pairs[j].score
		} else {
			return false
		}
	})

	bestScore := 0
	dp := make([]int, len(pairs))
	for i := 0; i < len(pairs); i++ {
		dp[i] = pairs[i].score

		for j := 0; j < i; j++ {
			if pairs[j].score <= pairs[i].score {
				dp[i] = max(dp[i], dp[j]+pairs[i].score)
			}
		}

		bestScore = max(bestScore, dp[i])
	}

	return bestScore
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}
