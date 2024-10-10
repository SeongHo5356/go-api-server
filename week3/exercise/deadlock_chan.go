package exercise

import "fmt"

func squareIt(inputChan, outputChan chan int){
	for x := range inputChan{
		outputChan <- x*x
	}
}

func DeadlockChan(){
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	go squareIt(inputChannel, outputChannel)
	for i := 0; i<10;i++{
		inputChannel <- i
	}
	for i := range outputChannel{
		fmt.Println(i)
	}
}