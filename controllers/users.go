package controllers

import "net/http"

type userController struct{}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//a byte slice can also be defined as a string
	w.Write([]byte("Hello from UserController!"))
}
