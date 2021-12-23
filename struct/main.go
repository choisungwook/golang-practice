package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}

func (p *person) changename(name string) {
	(*p).firstName = name
}

func main() {
	alex := person{firstName: "hello", lastName: "world"}
	alex.print()

	var alpha person
	alpha.firstName = "abc"
	alpha.lastName = "def"

	alpha.print()

	beta := person{firstName: "hello", lastName: "world"}
	beta.changename("tttt")
	beta.print()

}
