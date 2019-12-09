package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func checkAge(age int) bool {
	return age >= 18
}

func sellBeer(person person) string {
	if checkAge(person.Age) {
		return fmt.Sprintf("%s can buy beer", person.Name)
	}
	return fmt.Sprintf("%s CAN'T buy beer", person.Name)
}

func main() {
	person := person{
		Name: "Samuel",
		Age:  19,
	}
	fmt.Println(sellBeer(person))
}
