package ch01

import "fmt"

func Basic(){
	fmt.Println("Hello Go Basic\n")
}

func Condition(age int)  {

	if age<10 {
		fmt.Println("10<age<40\n")
	}
	if age < 40 && age >10 {
		fmt.Println("10<= age <=40\n")
	} else {
		 fmt.Println("a>40\n")
	}
}

func Switch(a int)  {
    switch(a){
	case 1:
		fmt.Println("a=1\n")
	case 2:
		fmt.Println("a=2\n")
	case 3:
		fmt.Println("a=3\n")
	default:
		fmt.Println("a=0\n")
	}
}