/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (33.42%)
 * Likes:    3738
 * Dislikes: 0
 * Total Accepted:    638.8K
 * Total Submissions: 1.9M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for l := 0; l < len(nums); l++ {
		// if l+1 < len(nums) && nums[l] == nums[l+1] {
		// 	continue
		// }
		if l != 0 && nums[l] == nums[l-1] {
			continue
		}
		t := -nums[l]
		i, j := l+1, len(nums)-1
		for i < j {
			if nums[i]+nums[j] == t {
				tmp := []int{nums[l], nums[i], nums[j]}
				res = append(res, tmp)
				for i < j && nums[i+1] == nums[i] {
					i++
				}
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < t {
				i++
			} else {
				j--
			}
		}
	}
	return res
}

// @lc code=end

