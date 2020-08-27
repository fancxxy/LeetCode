/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
*/

package main

import "fmt"

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1 -> 2 -> 3 -> 2 -> 1 -> nil
   f/s   s   f/s        f

	1 -> 2 -> 2 -> 1 -> nil
   f/s   s   f/s         f
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	first, second := head, head
	var prev, next *ListNode
	for first != nil && first.Next != nil {
		first = first.Next.Next

		next = second.Next
		second.Next = prev
		prev = second
		second = next
	}

	if first != nil {
		second = second.Next
	}

	for prev != nil && second != nil {
		if prev.Val != second.Val {
			return false
		}
		prev = prev.Next
		second = second.Next
	}

	if prev != second {
		return false
	}

	return true
}

func main() {
	list := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(isPalindrome(list))

	list = &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(isPalindrome(list))
}
