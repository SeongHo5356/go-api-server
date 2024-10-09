package lab

import "fmt"

type Snack struct{
	Name string
	Price int
}

type Drink struct{
	Name string
	Price int
}

func SaleAndGetPrice(item interface{}) int {
	switch v:= item.(type){
	case Snack:
		return int(float64(v.Price) * 0.9)
	case Drink:	
		return int(float64(v.Price) * 0.8)
	default:
		return 0
	}
} 

func MartPriceCalculator(){
	chips := Snack{"Pringles", 4000}
	cracker := Snack{"Ace", 2500}
	soda := Drink{"Sprite", 1800}
	coffee := Drink{"TOP", 2700}

	var total int = 0
	total += SaleAndGetPrice(chips)
	total += SaleAndGetPrice(cracker)
	total += SaleAndGetPrice(soda)
	total += SaleAndGetPrice(coffee)

	fmt.Println(total)
}