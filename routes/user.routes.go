package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Rodomaxi2/rest_go_task/db"
	"github.com/Rodomaxi2/rest_go_task/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	error := createdUser.Error

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleted user"))
}
