package tests

import (
	"leetcode/solutions"
	"testing"
)

//go test tests/add_two_numbers_test.go -v

func TestAddTwoNumbers(t *testing.T) {
	l1 := &solutions.ListNode{
		Val: 2,
		Next: &solutions.ListNode{
			Val: 4,
			Next: &solutions.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &solutions.ListNode{
		Val: 5,
		Next: &solutions.ListNode{
			Val: 6,
			Next: &solutions.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	x := solutions.AddTwoNumbers(l1, l2)
	if x.Val != 7 || x.Next.Val != 0 || x.Next.Next.Val != 8 {
		t.Error("Incorrect Value, expected 708")
	}
	l1 = &solutions.ListNode{
		Val: 9,
		Next: &solutions.ListNode{
			Val: 9,
			Next: &solutions.ListNode{
				Val: 9,
				Next: &solutions.ListNode{
					Val: 9,
					Next: &solutions.ListNode{
						Val: 9,
						Next: &solutions.ListNode{
							Val: 9,
							Next: &solutions.ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	l2 = &solutions.ListNode{
		Val: 9,
		Next: &solutions.ListNode{
			Val: 9,
			Next: &solutions.ListNode{
				Val: 9,
				Next: &solutions.ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	x = solutions.AddTwoNumbers(l1, l2)
	expected := [...]int{8, 9, 9, 9, 0, 0, 0, 1}
	i := 0
	for x != nil {
		if x.Val != expected[i] {
			t.Error("Incorrect Output")
		}
		x = x.Next
		i++
	}
}
