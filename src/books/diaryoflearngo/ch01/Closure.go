package ch01

/**
 * 闭包
 */
func ClosePackage() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
/**
 * 递归
 */
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}