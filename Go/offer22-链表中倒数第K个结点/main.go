/*
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。



示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
*/

package main

import "fmt"

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1 -> 2 -> 3 -> 4 -> 5 -> nil

	s         f    s         f
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first, second := head, head
	for first != nil && k > 0 {
		first = first.Next
		k--
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	return second
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	list = getKthFromEnd(list, 2)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")
}
