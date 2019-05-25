package routes

import (
	"fmt"
	"github.com/gorilla/mux"

	controllers "../controllers"
)

// GetUser routes
func GetUser(r *mux.Router) {
	fmt.Println("foo Router")
	r.HandleFunc("/user", controllers.UserController).Methods("GET")
}

