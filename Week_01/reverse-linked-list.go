package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 1}
	l4 := &ListNode{Val: 1}
	l5 := &ListNode{Val: 1}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	reverseList(l1)
	reverseList2(l1)
	reverseList3(l1)
}

func reverseList3(head *ListNode) *ListNode {
	//自己默写

	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return pre
}

func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head.Next, head, pre = pre, head.Next, head
	}
	return pre
}
