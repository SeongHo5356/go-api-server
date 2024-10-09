package lab

import "fmt"

func ListFunctions(){

	fmt.Println("프로그램이 종료되었습니다.")
	defer fmt.Println("프로그램이 종료되었습니다.")

	arr := []int{3,5,1,2,0}
	var sum int = 0
	var max int = 0
	var min int = 15

	// 각 요소 출력, 총합, 최대/최소
	for _, ar := range arr {
		fmt.Printf("%d ", ar)
		sum += ar
		if max < ar{
			max = ar
		}
		if min > ar{
			min = ar
		}
	}

	fmt.Println("\n배열 요소의 합:", sum)
	fmt.Printf("최대값: %d, 최소값: %d\n", max, min)

	n := len(arr)
	switch {
	case n < 5:
		fmt.Println("배열의 길이가 5보다 짧습니다.")
	case n == 5:
		fmt.Println("배열의 길이가 5입니다.")
	default: // n > 5
		fmt.Println("배열의 길이가 5보다 깁니다.")
	}

}