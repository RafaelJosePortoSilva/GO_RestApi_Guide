package routes

import (
	"apirest/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users", DeleteUser).Methods("DELETE")

	return router
}
