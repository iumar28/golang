// Go program to illustrate how to create an anonymous function
package main

//The anonymous function is a function without name. These are used when we want to create an inline  function.
//you can assign an anonymous function to a variable. But condition that the type of the variable and function should be same.

import "fmt"

func main() {

	func() { // Anonymous function

		fmt.Println("Welcome! to GeeksforGeeks")
	}()

}
