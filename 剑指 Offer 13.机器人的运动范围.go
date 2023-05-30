package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=剑指 Offer 13 lang=golang
 * @lcpr version=
 *
 * [剑指 Offer 13] 机器人的运动范围
 */

// @lc code=start
func movingCount(m int, n int, k int) int {
	// 尝试暴力求解
	// 1 <= n,m <= 100
	// 0 <= k <= 20
	// 向右下搜索
	visited := make(map[string]bool)
	queue := []idx{{0, 0}}

	for len(queue) > 0 {
		poll := queue[0]
		queue = queue[1:]

		key := strconv.FormatInt(int64(poll[0]), 10) + "," + strconv.FormatInt(int64(poll[1]), 10)
		if _, ok := visited[key]; ok {
			continue
		}

		if poll[0] >= m || poll[1] >= n {
			continue
		}

		x := poll[0]
		if x > 9 {
			x = x/10 + x%10
		}
		y := poll[1]
		if y > 9 {
			y = y/10 + y%10
		}

		if x+y > k {
			continue
		}

		visited[key] = true

		queue = append(queue, idx{poll[0] + 1, poll[1]}, idx{poll[0], poll[1] + 1})
	}

	return len(visited)
}

type idx [2]int

// @lc code=end

/*
// @lcpr case=start
// 2\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 3\n1\n0\n
// @lcpr case=end

*/
