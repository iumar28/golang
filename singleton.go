package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func instanceGetter() *single {
	if single_instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if single_instance == nil {
			fmt.Println("Creating the instance now now")
			single_instance = &single{}
		} else {
			fmt.Println("The instance has already been created")

		}
	} else {
		fmt.Println("The instance has already been created")
	}
	return single_instance
}

func main() {

	for i := 0; i < 10; i++ {
		go instance_getter()
	}
	fmt.Scanln()

}
