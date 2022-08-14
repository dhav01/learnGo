package main

import "fmt"

func main() {

	name := "dhaval"
	greet(name)

	a := 5
	b := 10
	sum, diff := calculate(a, b) //we can also return more than one value from func
	fmt.Println("Sum is", sum, "and difference is", diff)
}

// (int,int) because this function returns two values of int datatype
func calculate(a int, b int) (int, int) {
	sum := a + b
	diff := a - b

	return sum, diff
}

func greet(name string) {
	fmt.Println("Welcome to Golang,", name, "\b!") //\b to remove space between name and !
}
