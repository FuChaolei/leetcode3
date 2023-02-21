/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (64.20%)
 * Likes:    1006
 * Dislikes: 0
 * Total Accepted:    387K
 * Total Submissions: 602.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层序遍历结果：
 *
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{}
	res := [][]int{}
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		path := []int{}
		n := len(q)
		for n > 0 {
			n--
			tmp := q[0]
			path = append(path, tmp.Val)
			q = q[1:]
			if tmp.Left != nil {
				q = append(q, tmp.Left)
			}
			if tmp.Right != nil {
				q = append(q, tmp.Right)
			}
		}
		res = append(res, path)
	}
	return res
}

// @lc code=end

