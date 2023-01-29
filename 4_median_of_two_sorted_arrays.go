package leetcode_excercise

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := mergeNums(nums1, nums2)

	if len(merged)%2 != 0 {
		return float64(merged[len(merged)/2])
	} else {
		return float64(merged[len(merged)/2]+merged[len(merged)/2-1]) / 2
	}
}

func mergeNums(nums1 []int, nums2 []int) []int {
	nums1len := len(nums1)
	nums2len := len(nums2)
	merged := make([]int, 0, nums1len+nums2len)

	i, j := 0, 0
	for i < nums1len && j < nums2len {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	if i == nums1len {
		for j < nums2len {
			merged = append(merged, nums2[j])
			j++
		}
	} else if j == nums2len {
		for i < nums1len {
			merged = append(merged, nums1[i])
			i++
		}
	}

	return merged
}
