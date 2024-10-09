package lab

import "fmt"

type User struct{
	Name string
	Age int
}

func change(a *User, b*User){
	temp := a.Name
	a.Name = b.Name
	b.Name = temp
}

func sorting(list []User){
	n := len(list)
	for i := 0; i<n-1; i++{
		for j:=0;j<n-i-1;j++{
			if list[j].Name > list[j+1].Name{
				change(&list[j], &list[j+1])
			}
		}
	}
}

func SortByUser(){
	list := []User{
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}

	fmt.Printf("%s", "before sorting\n")
	for _, user := range list{
		fmt.Println(user.Name)
	}

	sorting(list)
	fmt.Print("\n")

	fmt.Printf("%s", "after sorting\n")
	for _, user := range list{
		fmt.Println(user.Name)
	}
}