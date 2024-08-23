package routes

import (
	"apirest/controllers"

	"github.com/gorilla/mux"
)

func setupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users", controllers.DeleteUser).Methods("DELETE")

	return router
}
