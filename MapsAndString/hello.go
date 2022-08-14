package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Maps and String Lecture in Go")

	//String in Golang
	String := "Welcome to Go"
	String2 := `Another way to declare string`
	fmt.Println(String, "\n", String2)

	//for i, _ := range String {
	//	print(String[i], " ")
	//}
	//it will not work because in golang, string is a sequence of bytes
	//even len function will not give correct output as len func
	//will return number of bytes in the string

	fmt.Printf("Length of string is %d\n", len(String))

	String = strings.Replace(String, "W", "w", 2)
	//Syntax is [String to replace, old char, new char, # times to replace]

	for i, char := range String { //you can print using both ways
		fmt.Printf("%c ", String[i])
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	//Maps
	marks := map[string]int{"English": 40, "Physics": 55}
	// [string] => key is of type string
	//int => value is of type int

	//accessing and printing
	marks["English"] = 100 //it will overwrite prev value of key
	marks["science"] = 70  //adding element to map
	fmt.Println(marks["English"], marks)

	delete(marks, "science") //removing element from map

	for key, value := range marks {
		fmt.Println(key, value)
	}
}
