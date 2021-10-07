package main

import "fmt"

func main() {
	var var_name string = "This is one of the ways to assign the variable this string"
	fmt.Println(var_name)

	auto_detect := "94567" //This detects automatically the type of variable assigned.
	fmt.Println(auto_detect)

	typee := 3.1
	fmt.Printf("%v,%T", typee, typee) // '%v' formats the value in default format and '%T' represents the type of the value

}
