package main

import "fmt"

func main() {
	fmt.Println("Welcome to Closures in Go")
	//before jumping to Closures, first learn anonymous func

	func() {
		fmt.Println("hello from anonymous")
	}()
	//here () means we are calling the function immediately

	//anonymous func are regular func which do not have any name

	innerFunc := greet()     //we store the func returned by the greet
	innerFunc("Hello World") //calling the variable like a function

	/*
		Closure is an inner function that has access to variables in the scope
		in which it was created (variables in the outer func) even after
		outer func finishes execution
	*/

	closure := increment()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	// o/p: 1 2 3
	/*
				Not only we are able to access the outer variable of anonymous func (i),
				we are incrementing it
				HOW ?
			As func() is an inner func, it has access to all the variables of its scope
		 (outer func) hence it is able to access and update the value
	*/

	//second := increment()
	//fmt.Println(second())
}

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Syntax of function: func funcName() returnType{}
func greet() func(string) {
	return func(message string) {
		fmt.Println(message, "from inner function")
	}
}

//That means greet func returns a func of type func(string)
