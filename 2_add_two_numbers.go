package leetcode_excercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{
		Val:  0,
		Next: nil,
	}

	pre := ret
	cur := pre.Next
	l1Digit := l1
	l2Digit := l2

	for l1Digit != nil || l2Digit != nil {
		over := false

		if cur == nil {
			cur = &ListNode{
				Val:  0,
				Next: nil,
			}
			pre.Next = cur
		}

		if l1Digit == nil {
			cur.Val += l2Digit.Val
			l2Digit = l2Digit.Next
		} else if l2Digit == nil {
			cur.Val += l1Digit.Val
			l1Digit = l1Digit.Next
		} else {
			cur.Val += l1Digit.Val + l2Digit.Val
			l1Digit = l1Digit.Next
			l2Digit = l2Digit.Next
		}

		if cur.Val > 9 {
			cur.Val -= 10
			over = true
		}

		if over {
			cur.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}

		cur = cur.Next
		pre = pre.Next
	}

	return ret.Next
}
