package main

import "fmt"

func main() {
	x := 15
	y := 25

	if x < y {
		fmt.Println("hello")
	} else {
		fmt.Println("yello")
	}

	isColored := false

	if isColored == false {
		fmt.Println("not colored")
	} else {
		fmt.Println("colored")
	}

	// Switch-case statements
	switch isColored {
	case false:
		fmt.Println("dadadada")
	default:
		fmt.Println("nani")
	}
}
