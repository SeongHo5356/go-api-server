package lab

import ("fmt")

func plus(x int, y int) int {
	return x + y
}
func minus(x int, y int) int{
	return x - y
}
func divide(x int, y int) int{
	return x / y
}
func multiply(x int, y int) int{
	return x * y
}

func Calculator(){
	var numA int
	var numB int
	var oper string
	fmt.Println("첫번째 숫자를 입력해주세요")
	fmt.Scan(&numA)
	fmt.Println("두번째 숫자를 입력해주세요")
	fmt.Scan(&numB)
	fmt.Println("연산자를 입력해주세요(+,-,*,/)")
	fmt.Scan(&oper)
	if oper == "+"{
		fmt.Printf("%d", plus(numA,numB))
	} else if oper == "-" {
		fmt.Printf("%d", minus(numA, numB))
	} else if oper == "*" {
		fmt.Printf("%d", multiply(numA, numB))
	} else {
		fmt.Printf("%d", divide(numA, numB))
	}
}
