/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (62.31%)
 * Likes:    4020
 * Dislikes: 0
 * Total Accepted:    619.1K
 * Total Submissions: 993.4K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	res := 0
	l, r := 0, len(height)-1
	lmax, rmax := height[0], height[r]
	for l < r {
		if lmax < rmax {
			res += lmax - height[l]
			l++
			lmax = max(height[l], lmax)
		} else {
			res += rmax - height[r]
			r--
			rmax = max(height[r], rmax)
		}
	}
	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

