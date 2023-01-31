package leetcode_excercise

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := new(ListNode)
	curList := mergedList

	for {
		if list1 == nil {
			curList.Next = list2
			break
		}
		if list2 == nil {
			curList.Next = list1
			break
		}
		if list1.Val < list2.Val {
			curList.Next = list1
			list1 = list1.Next
		} else {
			curList.Next = list2
			list2 = list2.Next
		}
		curList = curList.Next
	}

	return mergedList.Next
}
