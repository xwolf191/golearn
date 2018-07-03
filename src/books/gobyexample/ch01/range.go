package ch01

import "fmt"

func RangeBasic()  {
	//array遍历
	a := [5]int{1, 7, 4, 5}
	for _, num := range a {
		fmt.Println(num)
	}
	//slice遍历
	b := []int{1, 7, 4, 5}
	for i, bn := range b {
		fmt.Printf("i=%d,b[%d]=%d\n",i,i,bn)
	}
    //map遍历
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "0abcgo" {
		fmt.Println(i, c)
	}
}
