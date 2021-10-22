//srting mamipulation and string conversion
package main

import (
	"fmt"
	"strconv"
)

func main()  {
	first,last:="Umar","Firoz"
	first=first+last
	fmt.Println(first)
	fmt.Println(last)
	var age int=20
	nameandage:=first+strconv.Itoa(age)
	first=first+nameandage
	fmt.Println("Name and age is",first)



}
