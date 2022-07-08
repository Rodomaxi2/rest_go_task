package main

import (
	"net/http"

	"github.com/Rodomaxi2/rest_go_task/db"
	"github.com/Rodomaxi2/rest_go_task/models"
	"github.com/Rodomaxi2/rest_go_task/routes"

	"github.com/gorilla/mux"
)

func main() {

	// Instancia de la base de datos y conexion
	db.DBConnection()

	// Migracion de las tablas
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	// Instancia de router de gorilla/mux
	router := mux.NewRouter()

	// Creacion de rutas y manejadores para operaciones
	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("Delete")

	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	// Iniciacion del servidor
	http.ListenAndServe(":3000", router)
}
