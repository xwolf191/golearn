package array

import "strconv"

/**
<a href="https://leetcode.com/problems/add-binary/">Add Binary</a>
 二进制形式字符串运算
*/
func addBinary(a string, b string) string {
	s1, _ := strconv.ParseInt(a, 2, 64)
	s2, _ := strconv.ParseInt(b, 2, 64)
	i := s1 + s2
	return strconv.FormatInt(i, 2)
}
