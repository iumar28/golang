//               TYPECASTING EXPLAINED
package main

import "fmt"

func main()  {
	var meint int=4
	var mefloat float64=5.6
	
	//fmt.Println(meint+mefloat) //This will throw an error beacuse of mismatched data type
	fmt.Println(meint+int(mefloat)) //This will run fine because we have converted the float into int using typecasting
	
	var number int=65
	var firstletter string=string(number)
	fmt.Println(firstletter) //Now the 65 is converted into the character 'A' as per ASCII code

	var mefloat3,meflaot2 float64=5.6,6.1
	fmt.Println("this is the float multiplication",meflaot2*mefloat3)
	fmt.Println("the multiplication has been converted into int type",int(mefloat*meflaot2))
	var numericString string="324145"
	var number2 uint8=numericString[0]
	fmt.Println(number2)


}
