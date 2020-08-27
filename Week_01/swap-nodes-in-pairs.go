package main

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	 //递归，每次交换一对

	//保存当前节点和前置节点
	curr, pre := head, head.Next
	curr.Next = swapPairs(pre.Next)
	pre.Next = curr

	return pre
}
