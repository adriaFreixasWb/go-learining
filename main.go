package main

func main() {
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}
	//controllers.RegisterControllers()
	//http.ListenAndServe(":3000", nil)
}
