package main

import "fmt"

func main() {
	// Arrays: fixed typed
	var fruitArr [2]string

	// Assigning values
	fruitArr[0] = "Elma"
	fruitArr[1] = "Armut"

	// Declare and assign at the same time
	fruitArr2 := [2]string{"mandalina", "portakal"}

	// Slices
	fruitSlice := []string{"mandalina", "portakal"}

	fmt.Println(fruitArr, fruitArr2, fruitSlice)
	fmt.Println(len(fruitSlice), fruitSlice[1:2])
}
