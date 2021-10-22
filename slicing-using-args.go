package main

import (
	"fmt"
	"os"
)

func main()  {
	//Args is an array of strings which contain all the command line arguments passed
	firstLetter :=os.Args[1]
	secondLetter :=os.Args[2]
	thirdLetter :=os.Args[3]


	fmt.Println(firstLetter," ", secondLetter," ", thirdLetter)


}
