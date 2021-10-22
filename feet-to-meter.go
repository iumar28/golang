//feet to meter
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	getnumber:=os.Args[1]

	feet,_:=strconv.ParseFloat(getnumber,64)

	meter:=feet*0.3048
	fmt.Printf("%g in feet is equal to meters  %g",feet,meter)

}
