package main

import "fmt"

func main() {

	// Long one
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// compact one
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\t", i)
	}

	// FizzBuzz challenge
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}
