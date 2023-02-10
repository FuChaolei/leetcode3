/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode-cn.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (56.57%)
 * Likes:    361
 * Dislikes: 0
 * Total Accepted:    211.9K
 * Total Submissions: 375.1K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 50000
 * -50000 <= nums[i] <= 50000
 *
 *
 */

// @lc code=start
func sortArray(nums []int) []int {
	r := len(nums) - 1
	qsort(nums, 0, r)
	return nums
}
func qsort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	tmp := rand.Intn(r-l+1) + l
	nums[l], nums[tmp] = nums[tmp], nums[l]
	x, y := l, r
	t := nums[l]
	for x < y {
		for x < y && nums[y] >= t {
			y--
		}
		for x < y && nums[x] <= t {
			x++
		}
		nums[x], nums[y] = nums[y], nums[x]
	}
	nums[l], nums[x] = nums[x], nums[l]
	qsort(nums, l, x-1)
	qsort(nums, x+1, r)
}

// @lc code=end

