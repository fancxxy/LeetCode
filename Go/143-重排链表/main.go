/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/

package main

import "fmt"

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1 -> 2 -> 3 -> 4 -> 5
	1 -> 2 -> nil  3 <- 4 <- 5

	1 -> 2 -> 3 -> 4
	1 -> 2 -> nil 3 <- 4
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	first, second := head, head
	var prev *ListNode
	for first != nil && first.Next != nil {
		first = first.Next.Next
		prev = second
		second = second.Next
	}
	prev.Next = nil
	prev = nil

	for second != nil {
		next := second.Next
		second.Next = prev
		prev = second
		second = next
	}

	curr, left, right := head, head.Next, prev
	for left != nil {
		curr.Next = right
		right = right.Next
		curr = curr.Next

		curr.Next = left
		left = left.Next
		curr = curr.Next
	}
	curr.Next = right
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(list)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")

	list = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(list)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")
}
