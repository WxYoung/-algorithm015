package main

func main() {

}

func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}

	//快慢指针法
	slow, fast := head, head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		//如果相遇则有环
		if slow == fast {
			return true
		}

		//慢指针走1步
		slow = slow.Next

		//慢指针走2步
		fast = fast.Next.Next
	}

	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	//hash遍历法
	cache := make(map[*ListNode]struct{})

	//遍历链表
	for head.Next != nil {
		//如果已经存在，则为循环
		if _, ok := cache[head]; ok {
			return true
		}

		//保存到cache中
		cache[head] = struct{}{}
		head = head.Next
	}

	return false
}
