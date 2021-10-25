package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type circle struct{
	radius float64
}

type rect struct{
	width float64
	height float64
}

func (r rect) area() float64  {
	return r.height * r.width
}
func (c circle) area() float64  {
	return math.Pi* c.radius * c.radius
}

func main()  {
	c1:=circle{5}
	c2:=rect{4,5}

	shapes:=[]shape{c1,c2}

	for _,shape:= range shapes{
		fmt.Println(shape.area())
	}


	
}
