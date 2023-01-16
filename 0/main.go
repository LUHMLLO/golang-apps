package main

import (
	"fmt"

	"github.com/LUHMLLO/project0/functions"
)

func main() {

	test_variable_type := 3

	fmt.Println(functions.Works)
	fmt.Println(functions.MaximumCount([]int{1, 2}))
	fmt.Println(functions.IsString(test_variable_type))

}
