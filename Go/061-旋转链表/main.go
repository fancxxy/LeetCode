/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL 向右旋转 4 步: 2->0->1->NULL
*/
package main

import "fmt"

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var len int
	for curr := head; curr != nil; curr = curr.Next {
		len++
	}
	k = k % len

	first, second := head, head
	for first != nil && k > 0 {
		first = first.Next
		k--
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = head
	head = second.Next
	second.Next = nil

	return head
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list = rotateRight(list, 2)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")

	list = &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	list = rotateRight(list, 4)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")
}
