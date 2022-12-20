package main

import (
	"fmt"
)

func main() {
	fmt.Println("main started")
	port := 3000
	startWebServer(port)
}

func startWebServer(port int) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
}
