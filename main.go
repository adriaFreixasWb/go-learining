package main

import "fmt"

func main() {
	const c = 3
	fmt.Println(c + 3)

	//a bunch of code

	fmt.Println(c + 1.2)

	//type const
	const num int = 3

	fmt.Println(float32(num) + 1.2)
}
