package ch01

import (
	"strings"
	"fmt"
)

/**
 * byte数组
 */
func ByteAry(){
	var str = "HelloWorld"
	//转换为大小写
    a := strings.ToLower(str)
    b := strings.ToUpper(str)
    fmt.Println(a,b)
    //字符串转化为byte数组
    var c []byte = []byte(str)
    fmt.Println(c)
	//byte数组转化为字符串
    var d = string(c)
    fmt.Println(d)
}
