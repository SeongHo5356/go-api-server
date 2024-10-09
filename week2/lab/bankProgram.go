package lab

import (
	"errors"
	"fmt"
)

type BankAccount struct{
	balance int
}

//입금 
func (ba *BankAccount) Deposit(amount int) error { 
	if amount <= 0{
		return errors.New("입금액은 0보다 커야 합니다")
	}
	ba.balance += amount
	return nil
}

// 출금
func (ba *BankAccount) Withdraw(amount int) error { 
	if amount <= 0{
		return errors.New("출금액은 0보다 커야 합니다")
	}
	if ba.balance < amount {
		return errors.New("잔액이 부족합니다")
	}
	ba.balance -= amount
	return nil
}

func BankProgram(){
	account := BankAccount{balance:0}
	for {
		fmt.Print("입금 (1), 출금 (2), 종료 (Others) : ")
		var choice int
		fmt.Scan(&choice)

		switch choice{
		case 1:
			fmt.Print("입금할 금액을 입력하세요 : ")
			var amount int
			fmt.Scan(&amount)
			err := account.Deposit(amount)
			if err != nil {
				fmt.Println("오류 : ", err)
			}else{
				fmt.Printf("입금 성공! 현재 잔액: %d원\n", account.balance)
			}
		case 2:
			fmt.Print("출금할 금액을 입력하세요 : ")
			var amount int
			fmt.Scan(&amount)
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println("오류 : ", err)
			}else{
				fmt.Printf("출금 성공! 현재 잔액: %d원\n", account.balance)
			}
		default:
			fmt.Println("프로그램을 종료합니다.")
			return
		}
	}
}

