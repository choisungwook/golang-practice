package main

import "fmt"

func changeFirstValue(s []string) {
	s[0] = "hello"
}

func main() {
	a := []string{"a", "b"}

	fmt.Println(a)
	changeFirstValue(a)

	fmt.Println(a)
}
