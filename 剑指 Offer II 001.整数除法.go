package leetcode

/*
 * @lc app=leetcode.cn id=剑指 Offer II 001 lang=golang
 * @lcpr version=21907
 *
 * [剑指 Offer II 001] 整数除法
 */

// 参考: https://www.onblogs.net/2019/09/19/%E4%BD%8D%E8%BF%90%E7%AE%97%E5%AE%9E%E7%8E%B0%E5%8A%A0%E5%87%8F%E4%B9%98%E9%99%A4%E5%9B%9B%E5%88%99%E8%BF%90%E7%AE%97Java/
// @lc code=start
func divide(a int, b int) int {
	// 求绝对值
	// 除数
	negative := 0
	if a < 0 {
		negative ^= 1
		a = ^a + 1
	}
	// 被除数
	if b < 0 {
		negative ^= 1
		b = ^b + 1
	}

	div := 0
	// 假设求66 / 7, 可以视作求	66 - (7 * (2 ^ 3 + 2 ^ 0)) - 3
	// 转换为求多少个7相加再加上余数能等于66, 3为余数. 对于int32, 循环从i = 31开始
	for i := 31; i >= 0; i-- {
		if a < b {
			break
		}

		// 使用 a >> i 可以避免超过 int32 的数据范围
		if (a >> i) >= b {
			a -= (b << i)
			div += (1 << i)
		}
	}

	// 判断正负
	if negative == 1 {
		div = ^div + 1
	}

	if div < -2147483648 || div > 2147483647 {
		div = 2147483647
	}

	return div
}

// @lc code=end

/*
// @lcpr case=start
// 15\n2\n
// @lcpr case=end

// @lcpr case=start
// 7\n-3\n
// @lcpr case=end

// @lcpr case=start
// 0\n1\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
