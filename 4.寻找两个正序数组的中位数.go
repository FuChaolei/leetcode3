/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (40.68%)
 * Likes:    4401
 * Dislikes: 0
 * Total Accepted:    482.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [0,0], nums2 = [0,0]
 * 输出：0.00000
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums1 = [], nums2 = [1]
 * 输出：1.00000
 *
 *
 * 示例 5：
 *
 *
 * 输入：nums1 = [2], nums2 = []
 * 输出：2.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0
 * 0
 * 1
 * -10^6
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n2 < n1 {
		return findMedianSortedArrays(nums2, nums1)
	}
	k := (n1 + n2 + 1) / 2
	l, r := 0, n1
	var m1, m2 int
	for l < r {
		m1 = l + (r-l+1)/2
		m2 = k - m1
		if nums1[m1-1] > nums2[m2] {
			r = m1 - 1
		} else {
			l = m1
		}
	}
	m1 = l
	m2 = k - l
	var ml, mr int
	if m1-1 < 0 {
		ml = nums2[m2-1]
	} else if m2-1 < 0 {
		ml = nums1[m1-1]
	} else if nums1[m1-1] < nums2[m2-1] {
		ml = nums2[m2-1]
	} else {
		ml = nums1[m1-1]
	}
	if (n1+n2)%2 == 1 {
		return float64(ml)
	}
	if m1 >= n1 {
		mr = nums2[m2]
	} else if m2 >= n2 {
		mr = nums1[m1]
	} else if nums1[m1] < nums2[m2] {
		mr = nums1[m1]
	} else {
		mr = nums2[m2]
	}
	return float64(ml+mr) / 2.0
}

// @lc code=end

