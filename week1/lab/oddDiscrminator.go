package lab

import "fmt"

func OddDiscriminate() {
	var num int

	fmt.Print("숫자를 입력하세요 : ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("%d는 짝수입니다.", num)
	} else {
		fmt.Printf("%d는 홀수입니다.", num)
	}
}
