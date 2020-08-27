/*
编写一个程序，找到两个单链表相交的起始节点。

4 -> 1 -> 8 -> 4 -> 5
		  ^
		  |

5 -> 0 -> 1


0 -> 9 -> 1 -> 2 -> 4
               ^
               |
			   3

2 -> 6 -> 4
1 -> 5

注意：

如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
*/

package main

import "fmt"

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}

		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}

	return nodeA
}

func main() {
	list1 := &ListNode{4, &ListNode{1, &ListNode{8, &ListNode{4, &ListNode{5, nil}}}}}
	list2 := &ListNode{5, &ListNode{0, &ListNode{1, list1.Next.Next}}}
	list := getIntersectionNode(list1, list2)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")

	list1 = &ListNode{3, &ListNode{2, &ListNode{4, nil}}}
	list2 = &ListNode{0, &ListNode{9, &ListNode{1, list1.Next}}}
	list = getIntersectionNode(list1, list2)
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Println("NULL")
}
