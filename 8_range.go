package main

import "fmt"

func main() {

	// slice of id's
	ids := []int{15, 24, 14, 2, 1}

	// Loop through ids
	// i: index, id: element
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	// Loop through ids
	// i: index, id: element
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Adding ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Final sum: %d", sum)

	// Range with map
	emails := map[string]string{"deneme1": "deneme1@gmail.com", "deneme2": "deneme2@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	// just keys without values
	for key := range emails {
		fmt.Printf("Key: %s\n", key)
	}
}
