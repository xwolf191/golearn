package ch01

import "fmt"

type Person struct {
	id int32
	name string
	age int8
	status byte
}

func Struct(){
	var person = Person{id:3,name:"老王",age:43,status:1}
	fmt.Println(person)
	fmt.Println(person.name)
	ps := &person
	ps.name= "老王.rewrw"
	fmt.Println(person.name)

	person.age = 45
	fmt.Println(person)
}