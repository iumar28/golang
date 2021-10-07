// Go program to illustrate the use of function
package main

import "fmt"

func area(length, width int) int { //This function finds the area of the rectangle it accepts two parameters and returns the area

	Ar := length * width
	return Ar //the area is returned here
}

// Main function
func main() {

	// Display the area of the rectangle
	// with method calling
	fmt.Printf("Area of rectangle is : %d", area(12, 10)) //here, the area function is called with length=12 and width=10 is passed.
}
