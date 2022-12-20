package main

func main() {
	var i int
	for i < 5 {
		println(i)
		i++
		//exiting the loop early
		if i == 3 {
			continue
		}
		println("continuing...")
	}
	//controllers.RegisterControllers()
	//http.ListenAndServe(":3000", nil)
}
