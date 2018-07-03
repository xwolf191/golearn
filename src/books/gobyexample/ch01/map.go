package ch01

import "fmt"

func MapBasic()  {
    //定义map map[key]value.key,value指定具体的key,value的类型
	a := make(map[string]int)
	a["age"] = 32
	a["id"] = 12
	a["isDelete"] = 0
	fmt.Printf("a=%s,len=%d\n",a,len(a))
	fmt.Println("a[age]=",a["age"])
	//删除指定的key
	delete(a, "age")
	fmt.Printf("a=%s,len=%d\n",a,len(a))

	//是否包含key
	_,hasKey := a["id"]
	fmt.Println("hasKey:", hasKey)
    //直接创建map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
