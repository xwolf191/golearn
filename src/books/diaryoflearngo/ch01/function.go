package ch01

/**
 * 多返回值
 */
func Vals() (int, int) {
    return 3, 7
}

/**
 * 参数列表
 */
func Sum(nums ...int)(int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/**
 * 参数列表
 */
func Add(a []int)(int) {
	return Sum(a...)
}