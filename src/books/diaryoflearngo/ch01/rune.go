package ch01

/**
 * rune 定义、遍历
 */
func Rune(){
	var a []rune = []rune{4,5,54,548}
	for i, b := range a {
		println(i,b)
	}
}