package leetcode

/*
 * @lc app=leetcode.cn id=剑指 Offer II 003 lang=golang
 * @lcpr version=21907
 *
 * [剑指 Offer II 003] 前 n 个数字二进制中 1 的个数
 */

// @lc code=start
// 观察二进制数的 1 个数可得出以下规律
// 0 [ 1 ] [ 1  2 ] [ 1 2 2 3 ] [ 1 2 2 3 2 3 3 4 ] ...
// 从1开始后，下一个区间的 1 个数分布：前一个区间的 1 个数分布 + 前一个区间 1 个数分布都加 1：
// 比如前一个区间为 [ 1 2 2 3 ]， 那么下一个区间的 1 个数分布为 [ 1 2 2 3 ] + [ 2 3 3 4 ] -> [ 1 2 2 3 2 3 3 4 ]
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	dp := make([]int, n+1, n+1)
	u := 1
	d := 1
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; {
		for j := u; j <= d && i <= n; j++ {
			dp[i] = dp[j]
			i++
			if i > n {
				return dp
			}
		}
		for j := u; j <= d && i <= n; j++ {
			dp[i] = dp[j] + 1
			i++
			if i > n {
				return dp
			}
		}
		u = d + 1
		d = i - 1
	}

	return dp
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

*/
