package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Pega os parâmetros do request
	userID := params["id"]

	// Busca o Usuário pelo id
	_, user := findUserByID(userID)
	if user != nil {
		// Define o cabeçalho da response
		w.Header().Set("Content-Type", "application/json")

		// Condifica o obj user em json
		json.NewEncoder(w).Encode(users)

		return
	} else {
		// Se o usuário não for encontrado, retorna um erro
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "User not found"}`))
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Gera um ID único
	newUser.ID = generateUserID()
	// Dá append na lista de users
	users = append(users, newUser)

	// retorna o usuário recém criado na response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	index, user := findUserByID(userID)

	if user != nil {

		var updatedUser User
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users[index].Name = updatedUser.Name
		users[index].Email = updatedUser.Email

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "User not found"}`))
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var delUser User
	err := json.NewDecoder(r.Body).Decode(&delUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	index, _ := findUserByID(delUser.ID)

	if index < 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "User not found"}`))
		return
	} else {
		users = deleteElement(users, index)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(delUser)

}

func findUserByID(id string) (int, *User) {

	// Busca o Usuário pelo id
	for index, user := range users {
		if user.ID == id {
			return index, &user
		}
	}
	return -1, nil
}

func generateUserID() string {
	return fmt.Sprintf("%d", len(users)+1)
}

func deleteElement(slice []User, index int) []User {
	return append(slice[:index], slice[index+1:]...)
}
