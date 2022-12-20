package main

func main() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}
	//controllers.RegisterControllers()
	//http.ListenAndServe(":3000", nil)
}
