package exercise

import "fmt"

func DeadlockChanOnlyPrint(){
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10)
	go squareIt(inputChannel, outputChannel)
	for i := 0; i<10;i++{
		inputChannel <- i
	}
	for i := range outputChannel{
		fmt.Println(i)
	}
}