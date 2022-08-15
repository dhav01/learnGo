package main

import "fmt"

func main() {
	//myChannel := make(chan string)

	fmt.Println("Welcome to Channels in Go")

	//channels are a way of communication between go routines in an efficient way
	//we can send and receive values from channel

	/*
							VERY VERY IMP NOTE
							- If there are no goroutines writing data to channel and main thread is
						     trying to read from channel, then it will create deadlock as it keeps
							 on waiting indefinitely

						 	- When we are sending data to channel, that thread is blocked
					 		  until some other thread is reading the data from it
							So in this case as well, it will create deadlock

				 		CHANNELS DOES NOT HAVE ANY STORAGE CAPACITY, SO IT NEEDS SOMEONE
						TO SEND AND READ THE MESSAGES IMMEDIATELY
					If there’s no receiver, a message is stuck with the sender

			SO ANOTHER SOLUTION IS CREATING BUFFERED CHANNELS

			Syntax:
			channelName := make(chan type, size)
			it will creat a channel with this much storage size so messages do not
		need to be disposed immediately

	*/

	newChannel := make(chan string, 2)

	newChannel <- "hello world"
	newChannel <- "hello go"

	fmt.Println(<-newChannel)
	fmt.Println(<-newChannel)

	//go func() {
	//	myChannel <- "Hi Dhaval"
	//	fmt.Println("Welcome from go routine1")
	//}()
	//
	//go func() {
	//	myChannel <- "Welcome to Golang"
	//	fmt.Println("Welcome from go routine2")
	//
	//}()
	//
	//for i := 0; i < 2; i++ {
	//	fmt.Printf("%s\n", <-myChannel)
	//}
}
