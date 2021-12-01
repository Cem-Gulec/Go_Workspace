package main

import "fmt"

func main() {

	// map[key]value
	emails := make(map[string]string)

	// assigning values/keys to map
	emails["deneme"] = "deneme@gmail.com"

	fmt.Println(emails["deneme"], len(emails))

	// delete
	delete(emails, "deneme")
	fmt.Println(emails)

	// declare and assign map
	emails_2 := map[string]string{"denemedeneme": "deneme2@gmailcom"}

	fmt.Println(emails_2)

}
