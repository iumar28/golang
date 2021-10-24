// Go program to illustrate how a function return multiple values
package main

import "fmt" //The package is imported

// myfunc return 3 values of int type
func myfunc(p, q int) (int, int, int) { //This function return 3 values (all 3 are of int type)  and accepts 2 values of int type.
	return p - q, p * q, p + q
}

// Main Method
func main() {

	// The return values are assigned into
	// three different variables
	var myvar1, myvar2, myvar3 = myfunc(4, 2) //The threee values that are return by the function into three variables

	// Display the values of the three variables.
	fmt.Printf("Result is: %d", myvar1)
	fmt.Printf("\nResult is: %d", myvar2)
	fmt.Printf("\nResult is: %d", myvar3)
}
