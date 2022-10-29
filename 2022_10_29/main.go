package main

import (
	"fmt"
)

/*
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。


示例 1：


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(nums []int) *ListNode {
	head := new(ListNode)
	head.Val = nums[0]
	var tail *ListNode
	tail = head
	for i := 1; i < len(nums); i++ {
		node := new(ListNode)
		node.Val = nums[i]
		tail.Next = node
		tail = node
	}
	return head
}
func shownode(node *ListNode) []int {
	res := []int{}
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	add := 0
	for l1 != nil || l2 != nil {
		node := new(ListNode)
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		num := n1 + n2
		if num >= 10 {
			n := num % 10
			if add == 1 {
				node.Val = n + 1
			} else {
				node.Val = n
			}
			tail.Next = node
			tail = node
			//tail.Next.Val += 1 //报错
			//tail.Val += 1
			add = 1
		}

		if num < 10 {
			if add == 1 && num+1 < 10 {
				node.Val = num + 1
				add = 0
			} else if add == 1 && num+1 >= 10 {
				node.Val = (num + 1) % 10
				add = 1
			} else {
				node.Val = num
			}
			tail.Next = node
			tail = node
		}

	}
	if add == 1 {
		node := new(ListNode)
		node.Val = 1
		tail.Next = node
		tail = node
	}
	return head.Next
}

func main() {
	nums1 := []int{9, 9, 9, 9, 9, 9, 9, 9} //
	nums2 := []int{9, 9, 9, 9}             //

	node1 := BuildListNode(nums1)
	node2 := BuildListNode(nums2)
	fmt.Println(shownode(node1), shownode(node2))
	resnode := addTwoNumbers(node1, node2)
	fmt.Println(shownode(resnode)) //7 0 8

}
