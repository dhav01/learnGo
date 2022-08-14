package main

import "fmt"

func main() {

	//struct is like classes in Java
	//Syntax is
	//type name struct

	type person struct {
		name string
		age  int
	}

	type val struct {
		len int
		bre int
	}

	temp := val{bre: 10, len: 11}

	area := temp.bre * temp.len
	fmt.Println(area)

	person1 := person{name: "Raju", age: 22}
	fmt.Println(person1)
	fmt.Println(person1.name)
}
