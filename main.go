package main

func main() {
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, v := range wellKnownPorts {
		println(v)
	}
	//controllers.RegisterControllers()
	//http.ListenAndServe(":3000", nil)
}
