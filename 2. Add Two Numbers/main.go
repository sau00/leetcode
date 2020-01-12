package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	p := l1
	q := l2
	result := &ListNode{}
	current := result
	for p != nil || q != nil {
		x := 0
		y := 0

		if p != nil {
			x = p.Val
			p = p.Next
		}

		if q != nil {
			y = q.Val
			q = q.Next
		}

		val := x + y + carry
		carry = val / 10

		current.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}
		current = current.Next
	}

	if carry == 1 {
		current.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return result.Next
}

func main() {
	a1 := ListNode{2, nil}
	a2 := ListNode{4, &a1}
	a3 := ListNode{3, &a2}

	a3.print()

	b1 := ListNode{5, nil}
	b2 := ListNode{6, &b1}
	b3 := ListNode{4, &b2}

	c := addTwoNumbers(&a3, &b3)
	c.print()
}

func (l *ListNode) print() {
	for l.Next != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Print(l.Val)

	fmt.Println()
}
