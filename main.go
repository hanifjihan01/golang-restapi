package main

import (
	"Golang-Cuy/configs"
	"Golang-Cuy/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	configs.ConnectDB()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	log.Println("Server running 8080")
	http.ListenAndServe(":8080", router)
}
