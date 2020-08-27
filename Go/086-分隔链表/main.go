/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/

package main

import "fmt"

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l1, l2 := &ListNode{}, &ListNode{}
	little, large := l1, l2
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < x {
			little.Next = curr
			little = curr
		} else {
			large.Next = curr
			large = curr
		}
	}

	little.Next = l2.Next
	large.Next = nil
	return l1.Next
}

func main() {
	list := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	list = partition(list, 3)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")
}
