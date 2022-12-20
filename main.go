package main

import (
	"fmt"
)

func main() {
	port := 3000
	err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port, numberOfRetries int) error {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries is", numberOfRetries)
	return nil
}
