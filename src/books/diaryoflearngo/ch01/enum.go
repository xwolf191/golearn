package ch01

/**
 * 枚举类型
 */
func Enum(){
	const(
		Sunday  = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	a := Wednesday
	println(a)

}
