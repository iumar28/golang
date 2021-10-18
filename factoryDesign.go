package main

import "fmt"

const (
	student  = 1
	faculty= 2
)

type Gatekeeper interface {
	alarm() string
}

func checkPerson(m int) (Gatekeeper, error) {
	switch m {
	case student:
		return new(meStudent), nil
	case faculty:
		return new(meFaculty), nil
	default:
		return nil, fmt.Errorf("Don't Relax Bois this is suspicious!")
	}
}

type meStudent string

func (c *meStudent) alarm() string {
	return "Relax Bois This is student"
}

type meFaculty string

func (d *meFaculty) alarm() string {
	return "Relax Bois This is faculty"
}

func main() {
	obj1, _ := checkPerson(1)
	fmt.Printf("%T\n", obj1)
	fmt.Println(obj1.alarm())

	obj2, _ := checkPerson(2)
	fmt.Printf("%T\n", obj2)
	fmt.Println(obj2.alarm())

}
