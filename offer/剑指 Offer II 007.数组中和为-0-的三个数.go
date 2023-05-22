package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=剑指 Offer II 007 lang=golang
 * @lcpr version=
 *
 * [剑指 Offer II 007] 数组中和为 0 的三个数
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// 遍历处理
	// 顺序无关的话，应该可以先排序来简化
	// 问题在于如何去重呢

	// 复杂度过高，考虑使用二分查找

	// 手动排序
	ints := sort.IntSlice(nums)
	sort.Sort(ints)

	l := len(ints)

	var result [][]int
	for i := 0; i < l-2; i++ {
		if i != 0 && ints[i-1] == ints[i] {
			// 跳过重复
			continue
		}

		// 从尾部遍历
		for j := l - 1; j > i+1; j-- {
			if j != l-1 && ints[j] == ints[j+1] {
				// 跳过重复
				continue
			}

			v := ints[i] + ints[j]

			// 使用二分查找
			k := bsearch(ints, i+1, j-1, -v)
			if k != -1 {
				result = append(result, []int{ints[i], ints[j], ints[k]})
			}
		}

	}

	return result
}

func bsearch(ints []int, l, r, v int) int {
	for l <= r {
		m := (l + r) / 2

		if ints[m] > v {
			r = m - 1

		} else if ints[m] < v {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/
