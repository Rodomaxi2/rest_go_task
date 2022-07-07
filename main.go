package main

import (
	"net/http"

	"github.com/Rodomaxi2/rest_go_task/db"
	"github.com/Rodomaxi2/rest_go_task/models"
	"github.com/Rodomaxi2/rest_go_task/routes"

	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/user", routes.DeleteUserHandler).Methods("Delete")

	http.ListenAndServe(":3000", router)
}
