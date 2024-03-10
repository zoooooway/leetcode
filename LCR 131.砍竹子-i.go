package leetcode

/*
 * @lc app=leetcode.cn id=LCR 131 lang=golang
 * @lcpr version=30117
 *
 * [LCR 131] 砍竹子 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func cuttingBamboo(bamboo_len int) int {
	// f(1) = 1
	// f(2) = 2
	// f(3) = 3
	// f(4) = f(2) x f(2) = 4
	// f(5) = f(2) x f(3) = 6
	// f(6) = f(3) x f(3) = 9
	// f(7) = f(3) x f(4) = 12
	// f(8) = f(4) x f(4) = 16
	// f(9) = f(4) x f(5) = 24
	// ...
	if bamboo_len <= 3 {
		return bamboo_len
	}

	tb := [2]int{2, 2}
	for i := 4; i <= bamboo_len; i++ {
		temp := tb[0]
		tb[0] = tb[1]
		tb[1] = tb[0] * temp
		println("----", tb[0])
		println(tb[1])
	}
	return 33
}

// @lc code=end

/*
// @lcpr case=start
// 12\n
// @lcpr case=end

*/
