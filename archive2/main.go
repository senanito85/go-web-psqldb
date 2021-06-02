package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

type person struct {
	name  string
	email string
	phone string
}

func newPerson(name string, email string, phone string) person {
	p := person{}
	return p
}

func main() {

	var arr []person
	sum := 0
	for i := 1; i <= 50; i++ {
		sum += i
		p := (person{name: randomdata.FullName(0), email: randomdata.Email(), phone: randomdata.PhoneNumber()})
		f := (person{name: randomdata.FullName(1), email: randomdata.Email(), phone: randomdata.PhoneNumber()})
		arr = append(arr, p)
		arr = append(arr, f)
	}

	//randomdata.Number(999)
	//for index, element := range arr {
	//	fmt.Println(index, "=>", element)
	//}

	for _, v := range arr {
		fmt.Println("--------------------------------")
		fmt.Println(v.name + "\n" + v.email + "\n" + v.phone)
		fmt.Println("--------------------------------")

	}

}
