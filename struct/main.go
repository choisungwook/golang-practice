package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "hello", lastName: "world"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	var alpha person
	alpha.firstName = "abc"
	alpha.lastName = "def"
	fmt.Printf("%+v", alpha)
}
