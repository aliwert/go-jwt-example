package main

import (
	"fmt"
	"github.com/aliwert/go-jwt-example/app"
	"github.com/aliwert/go-jwt-example/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/v1/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port:", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
