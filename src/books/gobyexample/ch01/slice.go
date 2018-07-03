package ch01

import "fmt"

/**
 * slice 数据类型
 * @author xwolf
 */
func SliceBasic()  {
	//slice 初始化
	a := make([]string, 3)
	fmt.Println("a:", a)
	//赋值
	a[0]="a"
	a[1]="b"
	a[2]="c"
	fmt.Println("a:", a)
	//取值
	fmt.Println("a[1]:", a[1])
	//追加元素
	b := append(a,"d","e","f")
	//获取元素长度并输出
	fmt.Printf("b=%s,len=%d\n",b,len(b))

	//定义空slice,将b复制到c
	c := make([]string,len(b))
	copy(c,b)
	fmt.Println("c=",c)
    //[low,high] 截取slice 从low(包含)到high的元素
	d := c[1:]
	fmt.Println("d=",d)

	//直接定义slice
	e := []int{2,4,5,7}
	fmt.Println("e=",e)

	//二维slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)

}