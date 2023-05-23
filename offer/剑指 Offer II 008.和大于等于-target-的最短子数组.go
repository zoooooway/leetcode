package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=剑指 Offer II 008 lang=golang
 * @lcpr version=
 *
 * [剑指 Offer II 008] 和大于等于 target 的最短子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// target和数组中的元素都为正整数
	// 要求连续的话应该可以使用滑动窗口来解

	res := math.MaxInt
	for s, l, r := 0, 0, 0; r < len(nums); {
		if s < target {
			s += nums[r]
			r++
		}

		for ; s >= target && l < r; l++ {
			if r-l < res {
				res = r - l
			}
			s -= nums[l]
		}
	}

	if res == math.MaxInt {
		return 0
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/
