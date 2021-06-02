package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	log "github.com/sirupsen/logrus"
)

type person struct {
	name  string
	age   int
	email string
	phone string
	alive bool
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p

}

func main() {

	fmt.Println(person{name: randomdata.SillyName(), age: 30})

	sum := 0
	for i := 1; i < 4; i++ {
		sum += i
		peopleprinter()
	}

	logging()
}

func peopleprinter() {
	fmt.Println(randomdata.SillyName() + " lives at " + randomdata.Street() + randomdata.Address())
	fmt.Println(" ")
}

func logging() {
	log.Info("Sum equals to: 35")
	log.Warn("User kpmods was denied access to db02")
	log.Fatal("Server going to shut down")
}
