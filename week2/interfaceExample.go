package structExample

import "fmt"

func Describe2(i interface{}){
	value, ok := i.int()
	if ok{
		fmt.Println()
	}
	else{

	}
}