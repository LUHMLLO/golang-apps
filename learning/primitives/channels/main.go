package main

import "fmt"

func main() {
	//forks channel
	myChannel := make(chan string)

	//go rutine anonymous func
	go func() {
		myChannel <- "data"
	}()

	//joins channel
	msg := <-myChannel

	fmt.Println(msg)
}

/*
		|					 |
		|					 |
   	main* <1stin-----1stout> *channel
		|					 |
		|					 |
		*----->go rutine	 *
		|					 |
		|					 |
		*----->go rutine	 *
		|					 |
		|					 |
		*----->go rutine	 *
		|					 |
		|					 |
   	main*					 *
		|					 |
		|					 |
*/
