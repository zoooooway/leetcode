package leetcode

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=剑指 Offer II 002 lang=golang
 * @lcpr version=21907
 *
 * [剑指 Offer II 002] 二进制加法
 */

// @lc code=start
var m map[byte]string = map[byte]string{1: "1", 0: "0"}

func addBinary(a string, b string) string {
	u, d := sort(a, b)
	ui := len(u) - 1
	ud := len(d) - 1

	bs := make([]byte, ui+1, ui+1)
	var cy byte = 0
	var x byte = 0
	var y byte = 0
	var sb strings.Builder
	for ui >= 0 {
		fmt.Println(u[ui])
		if u[ui] == '1' {
			x = 1
		} else {
			x = 0
		}

		if ud < 0 {
			y = 0
		} else {
			if d[ud] == '1' {
				y = 1
			} else {
				y = 0
			}
		}

		s := x + y + cy
		if s >= 2 {
			cy = 1
			s -= 2
		} else {
			cy = 0
		}

		bs[ui] = s
		ui--
		ud--
	}

	for i := 0; i < len(bs); i++ {
		sb.WriteString(m[bs[i]])
	}

	if cy == 1 {
		return "1" + sb.String()
	}
	return sb.String()

}

func sort(a, b string) (string, string) {
	if len(a) < len(b) {
		return b, a
	}
	return a, b
}

// @lc code=end

/*
// @lcpr case=start
// "11"\n"10"\n
// @lcpr case=end

// @lcpr case=start
// "1010"\n"1011"\n
// @lcpr case=end

*/
