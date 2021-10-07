package main

import "fmt"

const PI = 3.14

func main()
{
	const GFG = "Any_string" // GFG is a string variable 
	fmt.Println("Hello", GFG)  //Printing the statements

	fmt.Println("Happy", PI, "Day") // PI is a constant, whose value can't be changed during the program run

	const Correct= true
	fmt.Println("Go rules?", Correct)
}
