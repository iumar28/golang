//Same package can't be imported using the same name in a file but it can be imported the alias of that package.
package main

import "fmt"
import alias "fmt"

func main(){
	fmt.Println("This is printed using fmt")
	alias.Println("This is printed using the alias of fmt")
	
	
}
