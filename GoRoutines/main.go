package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to GoRoutines lecture!")

	//in order to work with concurreny and parallelism,
	//we need to understand GoRoutines

	/*
		 GoRoutines are a Golang wrapper on top of threads and managed by
		 Go runtime rather than the operating system.

		A Goroutine is much like a thread to accomplish multiple tasks,
		but consumes fewer resources than OS threads
	*/
	//without()
	//with()

	//runtime.GOMAXPROCS(4) //specifying Go can use 4 cores for execution
	// (to add parallelism)
	start := time.Now()
	go func() {
		for i := 1; i < 6; i++ {
			fmt.Print(i, " ")
		}
	}()
	go func() {
		for i := 1; i < 6; i++ {
			fmt.Print(i, " ")
		}
	}()

	go func() {
		for i := 1; i < 6; i++ {
			fmt.Print(i, " ")
		}
	}()

	fmt.Println()
	go func() {
		for i := 10; i < 16; i++ {
			fmt.Print(i, " ")
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println()

	fmt.Println(elapsedTime.String())
	time.Sleep(time.Second)
}

func without() {

	start := time.Now()

	func() {
		for i := 1; i < 6; i++ {
			fmt.Print(i, " ")
		}
	}()

	fmt.Println()
	func() {
		for i := 10; i < 16; i++ {
			fmt.Print(i, " ")
		}
	}()

	elapsedTime := time.Since(start)
	fmt.Println()
	fmt.Println(elapsedTime.String())
	time.Sleep(time.Second)
}

//func with() {

//	start := time.Now()
//
//	go func() {
//		for i := 1; i < 6; i++ {
//			fmt.Print(i, " ")
//		}
//	}()
//
//	fmt.Println()
//	go func() {
//		for i := 10; i < 16; i++ {
//			fmt.Print(i, " ")
//		}
//	}()
//
//	elapsedTime := time.Since(start)
//
//	fmt.Println()
//
//	fmt.Println(elapsedTime.String())
//	time.Sleep(time.Second)
//}
