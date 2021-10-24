//srting mamipulation and string conversion
package main

import "fmt"

func main()  {
	type(
	 gram int
	 kilogram int
	 ton int
	)
	var(
		gold gram=50000
		vegetable kilogram=10
		wheat ton=15
	)
	// althoug the underlying type of gram,kilogram and ton is int yet they behave as if they are of different type
	// and we shall have to typecast them
	//gold=vegetable //This can't be done unless these are typecasted into the same type
	gold=gram(vegetable) //This is correct
	// Similar explanation can be given for the defined types accordingly
	fmt.Println(gold, vegetable, wheat)




}
