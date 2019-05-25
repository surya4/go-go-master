package main

import (
	"fmt"
	"log"
	"net/http"

	databases "./models"
	routes "./routes"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	// Get the "PORT" env variable
	port := "4000"
	fmt.Println("port running", port)

	// Initialize MySQL Client
	go databases.Connection()

	router := mux.NewRouter().StrictSlash(false)
	mainRoutes := router.PathPrefix("/api/v1/users").Subrouter()
	routes.GetUser(mainRoutes)

	fmt.Println("foo main")

	http.ListenAndServe(":" + port, router)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
