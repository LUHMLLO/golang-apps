package main

import "fmt"

func Foo() int {
	var (
		x   = 20
		y   = 2
		foo = "foo"
	)

	fmt.Println(foo)

	return x + y
}
