package main

import (
	"fmt"
	"time"
)

func someFunc(num int) error {
	fmt.Println(num)
	return nil
}

func main() {
	//go rutine async function
	go someFunc(5)

	//lets rutine execute
	time.Sleep(time.Second * 2)

	fmt.Println("hi")
}

/*
   	main*					*
		|					|
		|					|
		*----->fork			*
		|					|
		|					|
		*------------> child*starts
		|					|
		|					|
		* ---- process ---- *
		|					|
		|					|
		*<------------ child*ends
		|					|
		|					|
		*		 joins<-----*
		|					|
		|					|
   	main*					*
*/
