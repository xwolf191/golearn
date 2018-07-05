package ch03

import (
	"fmt"
)

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

/**
 * 匿名函数传参
 * <code>
 *  ch03.FuncTest3(func() rune {
 *		 return 32+43
 *	 })
 * </code>
 */
func FuncTest3(fn func() rune) rune {
	return fn() +5
}

//定义函数类型
type Sum func( a int ) int

/**
 * 函数参数
 * 		ch03.FuncTest4(func(a int) int {
		return a + 5
	},10)
 */
func FuncTest4( a Sum,b int) int{
   return  a(b)
}

/**
 * 多参数列表
 *  a := []int{1,2,3,4,5}
	ch03.FuncTest5(a...)
 */
func FuncTest5(a... int) int{
	total := 0
	for _,b :=range  a {
		total += b
	}
    return total
}

/**
 *defer延迟按照FIFO顺序执行
 */
func FuncTest6(a int){
	defer println("a")
	defer println("b")
	defer func(){
		println(1000/a)
	}()
	defer println("c")
}
/**
 * 函数类型定义传参
 */
func FuncTest7(a int){
	defer println("a")
	defer println("b")
	defer func(){
		println(1000/a)
	}()

	defer println("c")
}