package ch02

import "fmt"

/**
 * 表达式
 */
func Exp(){

	var a = 3 == 3
	var b = 5>6
	fmt.Println(a,b,!b,a&&b,a||b)
	var c = 1<<4
	fmt.Println(c)

}
