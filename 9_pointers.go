package main

import "fmt"

func main() {

	// Everything in GoLang basically passed by value

	a := 14
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Reading value from address
	fmt.Println(b, *b, *&a)

	// Change value with pointer
	*b = 8
	fmt.Println(a)

}
