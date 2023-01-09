package main

import "fmt"

func main() {

	//strings
	var string1 string = "Strings are double quoted"
	fmt.Println(string1)

	var string2 = "Go will automatically infer the type for us"
	fmt.Println(string2)

	var string3 string
	fmt.Println("default value of strings it's empty", string3)
	string3 = "u can set a value later"
	fmt.Println(string3)

	string4 := "this is for first time declarations only, cannot be used outside func"
	fmt.Println(string4)
}
