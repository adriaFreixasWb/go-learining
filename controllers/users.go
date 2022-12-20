package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//a byte slice can also be defined as a string
	w.Write([]byte("Hello from UserController!"))
}

// constructor function
func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/`),
	}
}
