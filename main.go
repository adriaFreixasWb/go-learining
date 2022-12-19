package main

import "fmt"

func main() {

	//regular
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	//implicit
	otherArr := [3]int{4, 5, 6}
	fmt.Println(otherArr)
	fmt.Println(otherArr[1])
}
