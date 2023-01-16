package functions

import "reflect"

var Works = "Hi, i was imported from the functions folder"

func IsString(s any) string {
	if reflect.TypeOf(s).Kind() == reflect.String {
		return "is a string"
	} else {

		return "not a string"
	}
}
