package leetcode

import (
	"bytes"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=剑指 Offer II 002 lang=golang
 * @lcpr version=21907
 *
 * [剑指 Offer II 002] 二进制加法
 */

// @lc code=start
func addBinary(a string, b string) string {
	var u string
	var d string
	if len(a) > len(b) {
		u = a
		d = b
	} else {
		u = b
		d = a
	}
	l := len(u)

	sum := make([]int, 0, l+1)
	bv := 0
	av := 0
	carry := 0

	di := len(d) - 1
	for i := l - 1; i >= 0; i-- {
		if di < 0 {
			bv = 0
		} else {
			bv = int(d[di] - '0')
			di--
		}

		if u[i] == '0' {
			av = 0
		} else {
			av = 1
		}

		s := bv + av + carry
		if s > 1 {
			carry = 1
			s -= 2
		} else {
			carry = 0
		}

		sum = append(sum, s)
	}

	if carry == 1 {
		sum = append(sum, 1)
	}
	// fmt.Println(sum)
	var bf bytes.Buffer
	for i := len(sum) - 1; i >= 0; i-- {
		bf.WriteString(fmt.Sprint(sum[i]))
	}

	return bf.String()
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
