package main

import "fmt"

func main() {
	var johnName *string = new(string)
	*johnName = "John"
	fmt.Println((*johnName))

	firstName := "Arthur"
	fmt.Println((firstName))

	//address of operator
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	//asing anopther value to first name, pointer addres remains the same
	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}
