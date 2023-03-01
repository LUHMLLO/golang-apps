package main

import "fmt"

func main() {
	//forks channel
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	//go rutine anonymous func
	go func() {
		myChannel <- "data"
	}()
	//go rutine anonymous func
	go func() {
		anotherChannel <- "another"
	}()

	//select joins channel
	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}

/*
   	main* <1stin-------------*-------------1stout> *channel
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
