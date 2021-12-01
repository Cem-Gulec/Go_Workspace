package main

import "fmt"

/* MAIN TYPES are listed below:

string
bool
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte -> alias for uint8
rune -> alias for int32
float32 float64
complex64 complex128

*/
func main() {

	// using var keyword
	var name string = "Cem"
	var number int32 = 25
	var isTrue = true
	// const isTrue = true
	isTrue = false

	fmt.Println(name, number)
	fmt.Printf("%T %T\n", number, isTrue)

	// Shorthand method
	//name2 := "Cem"
	//email := "cem.ggulecc@gmail.com"

	// or:
	name2, email := "Cem", "cem.ggulecc@gmail.com"

	fmt.Println(name2, email)

}
