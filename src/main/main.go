package main

import (
	"fmt"
	"books/goaction/ch01"
	"log"
)

func main()  {

	log.Println("系统故障")
	for  i:=0;i<10 ;i++  {
		fmt.Printf("哈喽...go,i=%d\n",i)
	}
	ch01.Basic()
	ch01.Switch(1)
}
