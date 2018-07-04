package ch01

import "fmt"

func Point() {
	var a = 1
	//获取内存地址
	var b = &a
	fmt.Println("a=",&a)
	//获取内存地址对应的内容
	fmt.Println("b=",*b)
}


