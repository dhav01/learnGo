package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays and slices class of GoLang")
	var array [5]int
	array[0] = 5
	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)
	fmt.Println(b)

	//we can also initialize array using ellipses (...)
	x := [...]int{1, 3, 4, 6, 7, 8, 10} //compiler will identify len of arr based on elements in the arr declaration
	fmt.Println(len(x))                 //len(x) returns the size of array

	for _, num := range x {
		fmt.Print(num, " ")
	}

	//Slices in Golang
	numbers := []int{1, 2, 3, 4, 5} //syntax to create slice in Go

	slice := numbers[1:4] //creating slice from array/slice having numbers from 1st ind (inc) to 4th ind()exc
	fmt.Println(slice)

	temp := []int{10, 20}
	temp = append(temp, numbers...) //to copy all the ele from one
	// slice to another slice, first you need to unslice the array using ...
	fmt.Println(temp)

	var t []int //declaring an empty slice
	copy(t, numbers)
	fmt.Println(t) //As size of t is 0, its capacity is 0 so it cannot hold any ele

	xyz := []int{7, 8, 9, 10}
	copy(xyz, numbers)
	fmt.Println(xyz)
}
