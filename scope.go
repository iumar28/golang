package main

import "fmt"  //This is the file scope

const one2=12 //This is the package scope

//This is the package scope
func  main()  { //The block scope has started here
	var block = "scope" //This block can be used within the main function only
	fmt.Println("This is the blocked scope"," ",one2, " ",block)

}
