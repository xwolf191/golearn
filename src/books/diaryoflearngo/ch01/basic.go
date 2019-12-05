package ch01

import (
	"fmt"
)

/**
 * go基本变量声明、初始化
 */
func Basic() {
	fmt.Println("Hello Go Basic")
	var a int = 1
	fmt.Printf("a=%d\n", a)
	//初始化多个变量
	var b, c = 3, 3
	fmt.Printf("%d+%d=%d\n", b, c, b+c)
	//默认值0
	var e int
	fmt.Printf("e=%d\n", e)
	//变量声明及初始化
	f := "var"
	fmt.Printf("f=%s\n", f)
	//强制转换
	var g byte = 123
	println(g)
}

/**
 * 条件判断
 */
func Condition(age int) {

	if age < 10 {
		fmt.Println("10<age<40")
	} else if age < 40 && age > 10 {
		fmt.Println("10<= age <=40")
	} else {
		fmt.Println("a>40")
	}
}

/**
  switch  default
*/
func Switch(a int) {
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	default:
		fmt.Println("a=0")
	}

}

/**
 *  循环
 */
func Cycle() {
	//定义常量
	const LENGTH = 5
	//定义数组
	var a = [LENGTH]int{3, 43, 3, 5}
	for i := 0; i < LENGTH; i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}
}
