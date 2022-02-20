package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) travese(arr []int) []int {
	arr = append(arr, l.Val)
	if l.Next != nil {
		return l.Next.travese(arr)
	}
	return arr
}

// ToList converts linked list to an array of its values
func (l *ListNode) ToList() []int {
	var values []int
	if l != nil {
		return l.travese(values)
	}
	return []int{0}
}

// addTwoNumbers takes in 2 linked lists. It adds the values up fom each linked list
// each value in the linked list will be between 0 <= 9
// if the sum of those values is greater than 10 then another linked list will need to be created
// Note: initially i tried to a fully recursive solution but was fiddly when l1 and l2 lengths dont match
// Difficulty: medium
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummyHead := &ListNode{Val: 0}
	curr := dummyHead
	x := 0
	y := 0

	p := l1
	q := l2

	for {
		if p == nil && q == nil {
			break
		}
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		// sum the previous carry
		sum := x + y + carry
		// i.e. sum = 13 13/10 = 1.3 => 1
		carry = sum / 10
		// largest carry can be is 1 or 0 i.e 9 + 9 + 1 = 19
		// create a new node with the carry
		curr.Next = &ListNode{Val: sum % 10}

		// advance pointers
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		// the final remaining one
		curr.Next = &ListNode{Val: carry}
	}

	// base case
	return dummyHead.Next

}
