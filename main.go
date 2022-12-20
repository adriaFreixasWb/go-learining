package main

import (
	"fmt"
)

func main() {
	fmt.Println("main started")
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries is", numberOfRetries)
}
