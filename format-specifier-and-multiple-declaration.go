// Multiple declaaration and the formar specifier
package main

import "fmt"

func main()  {
	var(
		city string="Los Anegles"
		distancekm int=345
		distancem int=700
		isgood bool=true
	)
	fmt.Printf("The distance between my city and the %[1]v is %[2]v km and %[3]v metre and it is a good city %[4]v",city,distancekm,distancem,isgood)
	fmt.Printf("\nthe data type of city is %T",city)

}
