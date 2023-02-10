package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode.cn/problems/reorder-list/description/
 *
 * algorithms
 * Medium (64.57%)
 * Likes:    1113
 * Dislikes: 0
 * Total Accepted:    227.4K
 * Total Submissions: 352.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	half := gethalf(head)
	half_fin := reverse(half)
	l1 := head
	merge(l1, half_fin)
}
func gethalf(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	res := s.Next
	s.Next = nil
	return res
}
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	cur := head
	pre := dummy
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = tmp
	}
	return dummy.Next
}

func merge(l1 *ListNode, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		tmp1 := l1.Next
		tmp2 := l2.Next
		l1.Next = l2
		l2.Next = tmp1
		l1 = tmp1
		l2 = tmp2
	}
}
func main() {
	l := []int{1, 2, 3, 4, 5}
	head := &ListNode{
		Val: 0,
	}
	cur := head
	for i, n := 0, len(l); i < n; i++ {
		cur.Next = &ListNode{Val: l[i]}
		cur = cur.Next
	}
	reorderList(head)
	tmp := head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

// @lc code=end
