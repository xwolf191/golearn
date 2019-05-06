package leetcode

import (
	"fmt"
	"strings"
)

/***
 * 151. Reverse Words in a String
Given an input string, reverse the string word by word.
Example 1:
Input: "the sky is blue"
Output: "blue is sky the"
Example 2:
Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:
Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/
func Reverse(s string) string {
	str := strings.TrimSpace(s)
	strAry := strings.Split(str, " ")
	fmt.Println(strAry)
	size := len(strAry) - 1
	res := ""
	for i := size; i >= 0; i-- {
		if len(strAry[i]) > 0 {
			res += strAry[i]
			if i != 0 {
				res += " "
			}
		}
	}
	return res
}
