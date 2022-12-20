package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	//the coma at the end on LastName definition is needed to have your closing brance on a different line
	u2 := user{ID: 1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println(u2)
}
