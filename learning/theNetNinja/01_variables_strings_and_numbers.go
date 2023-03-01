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

	string4 := ":= this is for first time declarations only, cannot be used outside function"
	fmt.Println(string4)

	//intergers
	var interger1 int = 30
	fmt.Println("static int: ", interger1)

	var interger2 = 20
	fmt.Println("unspecified int: ", interger2)

	interger3 := 10
	fmt.Println(":= int: ", interger3)

	//bits & memory
	var bit8 int8 = 127
	fmt.Println("int8: ", bit8)

	var bit16 int16 = 32767
	fmt.Println("int16: ", bit16)

	var bit32 int32 = 2147483647
	fmt.Println("int32: ", bit32)

	var bit64 int64 = 9223372036854775807
	fmt.Println("int64: ", bit64)

	var unsignedInterger uint = 1234567890
	fmt.Println("uint: ", unsignedInterger)

	var float32 float32 = 25.98
	fmt.Println("float32: ", float32)

	var float64 float64 = 16231238193123719238.67263
	fmt.Println("float64: ", float64)

	autoFloat := 12312.663
	fmt.Println(":= float: ", autoFloat)

}
