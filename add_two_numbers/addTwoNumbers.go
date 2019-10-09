package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (ListNode *ListNode) String() string {
	s := ""
	cursor := ListNode
	for cursor != nil {
		s = strconv.Itoa(cursor.Val) + s
		cursor = cursor.Next
	}
	return s
}

func createListNode(s string) *ListNode {
	if len(s) == 0 {
		return nil
	}

	val, _ := strconv.Atoi(s[len(s) - 1:])
	root := &ListNode{
		Val: val,
		Next: createListNode(s[: len(s) - 1]),
	}
	return root
}

func main() {
	// 2 -> 4 -> 3
	left := &ListNode {
		Val: 2,
		Next: &ListNode {
			Val: 4,
			Next: &ListNode {
				Val: 3,
				Next: nil,
			},
		},
	}
	// 5 -> 6 -> 4
	right := &ListNode{ 5, &ListNode{ 6 , &ListNode{ 4, nil }}}
	// 7 -> 0 -> 8
	// result := &ListNode{ 7, &ListNode{ 0, &ListNode{ 8, nil }}}

	fmt.Printf("数学法: %s + %s = %s\n", left, right, addTwoNumbers(left, right))

	left, right = createListNode("1234"), createListNode("4321")
	fmt.Printf("数学法: %s + %s = %s\n", left, right, addTwoNumbers(left, right))
	left, right = createListNode("5"), createListNode("5")
	fmt.Printf("数学法: %s + %s = %s\n", left, right, addTwoNumbers(left, right))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    root := &ListNode{0, nil}
    cursor := root
    carry := 0
    for l1 != nil || l2 != nil || carry != 0 {
        var val1, val2 int
        if l1 != nil {
            val1 = l1.Val
        }
        if l2 != nil {
            val2 = l2.Val
        }
        sumVal := val1 + val2 + carry
        carry = sumVal / 10
        
        sumNode := &ListNode{sumVal % 10, nil}
        cursor.Next = sumNode
        cursor = sumNode
        
        if l1 != nil {
			l1 = l1.Next
		}
        if l2 != nil {
			l2 = l2.Next
		}
	}
	
	return root.Next
}

