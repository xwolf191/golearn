package ch03

import "fmt"

/**
 * 无参无返回值
 */
func FuncTest1(){
   fmt.Println("无参函数")
}

/**
 * 有参数、返回值
 */
func FuncTest2(a int) int{
	return a + 3
}

func FuncTest3(fn func() rune) rune {
	return fn()
}