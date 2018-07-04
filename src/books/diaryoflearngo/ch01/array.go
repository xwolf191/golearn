package ch01

import "fmt"

/**
 * 数组定义、初始化、赋值、取值、遍历
 */
func AryBasic(){

	//数组定义
	var a [5]int
	fmt.Println("a:", a)
	//指定索引位置元素赋值
	a[3] = 32
	fmt.Println("a:", a)
	//取值
	fmt.Println("a[3]=",a[3])
    //数组初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

}

const LineLength = 5
const RowLength = 4
/**
 * 二维数组
 */
func TwoDimensionsAry(){
	  //定义五行四列的二维数组
	  a :=[LineLength][RowLength]int{{3,5,6,7},{8,15,14}}
	  fmt.Println("a=",a)
	  //遍历二维数组
	  for i:=0;i< LineLength;i++{
	  	for j:=0;j< RowLength;j++{
	  		fmt.Printf("a[%d][%d]=%d\n",i,j,a[i][j])
		}
	  }
}