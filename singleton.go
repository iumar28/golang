package main

import (
	"fmt"
	"sync"
)

var lock=&sync.Mutex{}

type single struct {
}
var single_instance *single

func instance_getter() *single {
	if single_instance==nil	{
		lock.Lock()  //Here the lock variable is used in orde to prevent more than two threads simultaneously entering
		defer lock.Unlock() //defer is used so that it is unlocked after everything in the scope has been executed
		if single_instance==nil {
			fmt.Println("Creating the instance now now")
			single_instance = &single{}
		} else{
			fmt.Println("The instance has already been created")

		}
	} else {
		fmt.Println("The instance has already been created")
	}
	return single_instance
}

func main(){

	for i:=0;i<10;i++ {
		go instance_getter()
	}
	fmt.Scanln()

}
