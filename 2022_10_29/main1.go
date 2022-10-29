package main

import (
	"fmt"
	"math"
	"strconv"
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
	s1 := []int{}
	s2 := []int{}
	n1, n2 := 0, 0

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	for i := 0; i < len(s1); i++ {
		if i != 0 {
			cf := math.Pow10(i)
			num, _ := strconv.Atoi(fmt.Sprintf("%1.0f", cf))
			n1 += num * s1[i]
		} else {
			n1 += s1[i]
		}
	}

	for i := 0; i < len(s2); i++ {
		if i != 0 {
			cf := math.Pow10(i)
			num, err := strconv.Atoi(fmt.Sprintf("%1.0f", cf))
			fmt.Println(err)
			n2 += num * s2[i]
		} else {
			n2 += s2[i]
		}
	}

	res := n1 + n2

	s := strconv.Itoa(res)
	head := new(ListNode)
	n, _ := strconv.Atoi(string(s[0]))
	head.Val = n
	for i := 1; i < len(s); i++ {
		node := new(ListNode)
		n, _ := strconv.Atoi(string(s[i]))
		node.Val = n
		node.Next = head
		head = node
	}
	fmt.Println(n1, n2)
	return head
}

func main() {
	//nums1 := []int{2, 4, 3}
	//nums2 := []int{5, 6, 4}
	//碰到这种长的就用问题，字符串转成int报错超出了范围
	nums1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	nums2 := []int{4, 3, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9, 0}
	node1 := BuildListNode(nums1)
	node2 := BuildListNode(nums2)
	resnode := addTwoNumbers(node1, node2)
	fmt.Println(shownode(resnode))

}
