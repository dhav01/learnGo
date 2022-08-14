package main

import "fmt"

func main() {
	fmt.Printf("Welcome to class of Pointers in Go\n")

	num := 10
	fmt.Println(num)  //prints the value of num
	fmt.Println(&num) // prints the address of num

	secondMultiply(num)
	fmt.Println(num)
	fmt.Println(&num)

	/*
		secondMultiply is known as pass by value
		Here we are passing a copy of variable hence there is no change
		in the value of variable

		multiply is known as pass by reference
		Here we are passing the address of variable hence whatever operation
		is done on the pointer, it affects the actual variable

		but if the type is *int, you cannot change it to string and vice versa
		is also true
	*/
	t := 343.43 //it is float64
	multiply(&t)
	fmt.Println(t)

	//we can also pass pointer of composite data type like arrays
	a := [3]int{1, 2, 3}
	changeValue(&a)
	fmt.Println(a)

}

func changeValue(p *[3]int) {
	(*p)[0] = 7
	p[1] *= 7 // it will multiply 7 to the value stored at 1st index address
	(*p)[2] = 7
}

func secondMultiply(val int) {
	val = 100
	/*
	   Go has special new function which creates allocates memory in the RAM and gives you the pointer to that memory
	    x := new(int)
	    x = reference to address of x
	    *x = value stored at address of x
	*/
}

func multiply(val *float64) { //val is the reference to the address of variable
	*val = 43.5 //*val means value stored at the reference to the address

	//val is the address of another variable whereas *val is the value stored at the address of another variable
	//*operator is known as dereferencing operator
}
