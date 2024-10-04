package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}

	poxancum := 0

	var arjeq int

	for l1.Next != nil || l2.Next != nil {
		arjeq = (l1.Val + l2.Val + poxancum)
		poxancum = arjeq % 10
		l3.Val = arjeq / 10
		l3 = l3.Next
	}

	return l3
}

func main() {
	l1 := &ListNode{}
	l1.Val = 2
	l1.Next.Val = 4
	l1.Next.Next.Val = 3

	l2 := &ListNode{}

	l2.Val = 5
	l2.Next.Val = 6
	l2.Next.Next.Val = 4

	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3.Val)

}
