package main

import "fmt"

func main() {

	//regular
	arr := [3]int{1, 2, 3}
	//: means all indexes
	aSlice := arr[:]

	fmt.Println(arr, aSlice)

	//acts like a pointer
	arr[1] = 42
	aSlice[2] = 27
	fmt.Println(arr, aSlice)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 42, 27)
	fmt.Println(slice)

	//slice from slice form 1 to end
	s2 := slice[1:]
	//up to 2 elements
	s3 := slice[:2]
	//indexes 1 and 2
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)
}
