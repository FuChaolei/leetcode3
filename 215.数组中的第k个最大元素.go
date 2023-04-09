/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.71%)
 * Likes:    1268
 * Dislikes: 0
 * Total Accepted:    426.7K
 * Total Submissions: 659.6K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 *
 *
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().Unix())
	n := len(nums) - 1
	return qsort(nums, k, 0, n)
}
func qsort(nums []int, k int, l int, r int) int {
	tmp := rand.Intn(r-l+1) + l
	nums[l], nums[tmp] = nums[tmp], nums[l]
	x, y := l, r
	tmp = nums[l]
	for x < y {
		for x < y && nums[y] <= tmp {
			y--
		}
		for x < y && nums[x] >= tmp {
			x++
		}
		nums[x], nums[y] = nums[y], nums[x]
	}
	nums[x], nums[l] = nums[l], nums[x]
	if x == k-1 {
		return nums[x]
	} else if x < k-1 {
		return qsort(nums, k, x+1, r)
	} else {
		return qsort(nums, k, l, x-1)
	}
}

// @lc code=end

