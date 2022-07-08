package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Rodomaxi2/rest_go_task/db"
	"github.com/Rodomaxi2/rest_go_task/models"
	"github.com/gorilla/mux"
)

// Manejador para obtener todos los usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo array basado en el struct User
	var users []models.User
	// Variable que realiza la operacion de consulta en base de datos y la guarda en task
	db.DB.Find(&users)

	// Formato y envio de respuesta en json
	json.NewEncoder(w).Encode(&users)
	// Codigo de estado OK
	w.WriteHeader(http.StatusOK)
}

// Manejador para obtener un solo usuario
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo basado en el struct User
	var user models.User
	//Se crea una variable que almacena los parametros del request
	params := mux.Vars(r)

	// Se busca la primer coincidencia en task que coincida con el id obtenido de los parametros
	db.DB.First(&user, params["id"])

	//En caso de regresar un zero value (error)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// Asociacion del usuario buscado con sus tareas
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	// Formato y envio de respuesta en json
	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}

// Manejador para crear un usuario
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo basado en el struct User
	var user models.User
	// Decodificacion del json en el body del request y almacenado en user
	json.NewDecoder(r.Body).Decode(&user)

	// Se crea un row en con la estructura y tabla basado en user
	createdUser := db.DB.Create(&user)
	// La operacion puede ser erronea y se almacena en la propierda Error
	error := createdUser.Error

	// En caso de error
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
		return
	}

	// Formato y envio de respuesta en json
	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para borrar un usuario
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo basado en el struct Task
	var user models.User
	//Se crea una variable que almacena los parametros del request
	params := mux.Vars(r)

	// Se busca la primer coincidencia en user que coincida con el id obtenido de los parametros
	db.DB.First(&user, params["id"])

	//En caso de regresar un zero value (error)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	db.DB.Delete(&user) //Solo "oculta"
	// Realmente remueve el registro
	// db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusNoContent)

}
