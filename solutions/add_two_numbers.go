/*
https://leetcode.com/problems/add-two-numbers/description/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(carry *int, v1, v2 int) int {
	sum := v1 + v2 + *carry
	if sum >= 10 {
		*carry = 1
		return sum - 10
	} else {
		*carry = 0
		return sum
	}

}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	mainNode := &ListNode{
		Val: add(&carry, l1.Val, l2.Val),
	}
	l1, l2 = l1.Next, l2.Next
	current := mainNode
	var next *ListNode
	for {
		if l1 == nil {
			if l2 == nil {
				if carry != 0 {
					current.Next = &ListNode{
						Val: 1,
					}
				}
				break
			}
			next = &ListNode{
				Val: add(&carry, l2.Val, 0),
			}
			l2 = l2.Next
		} else if l2 == nil {
			next = &ListNode{
				Val: add(&carry, l1.Val, 0),
			}
			l1 = l1.Next
		} else {
			next = &ListNode{
				Val: add(&carry, l1.Val, l2.Val),
			}
			l1, l2 = l1.Next, l2.Next
		}
		current.Next = next
		current = next
	}
	return mainNode
}
