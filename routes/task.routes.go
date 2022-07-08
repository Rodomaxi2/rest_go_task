package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Rodomaxi2/rest_go_task/db"
	"github.com/Rodomaxi2/rest_go_task/models"
	"github.com/gorilla/mux"
)

// Manejador para obtener todas las tareas
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo array basado en el struct Task
	var task []models.Task
	// Variable que realiza la operacion de consulta en base de datos y la guarda en task
	db.DB.Find(&task)

	// Formato y envio de respuesta en json
	json.NewEncoder(w).Encode(task)
	// Codigo de estado OK
	w.WriteHeader(http.StatusOK)
}

// Manejador para crear una tarea
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo basado en el struct Task
	var task models.Task
	// Decodificacion del json en el body del request y almacenado en task
	json.NewDecoder(r.Body).Decode(&task)
	// Se crea un row en con la estructura y tabla basado en task
	createdTask := db.DB.Create(&task)

	// La operacion puede ser erronea y se almacena
	err := createdTask.Error

	// En caso de error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	// Formato y envio de respuesta en json
	json.NewEncoder(w).Encode(&task)
}

// Manejador para obtener una tarea
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo basado en el struct Task
	var task models.Task
	//Se crea una variable que almacena los parametros del request
	params := mux.Vars(r)

	// Se busca la primer coincidencia en task que coincida con el id obtenido de los parametros
	db.DB.First(&task, params["id"])

	//En caso de regresar un zero value (error)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	// Formato y envio de respuesta en json
	json.NewEncoder(w).Encode(&task)
	w.WriteHeader(http.StatusOK)
}

// Manejador para eliminar una tarea
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Variable con tipo basado en el struct Task
	var task models.Task
	//Se crea una variable que almacena los parametros del request
	params := mux.Vars(r)

	// Se busca la primer coincidencia en task que coincida con el id obtenido de los parametros
	db.DB.First(&task, params["id"])

	//En caso de regresar un zero value (error)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	//Se activa la variable de borrado en el registro
	db.DB.Delete(&task)
	// Codigo de estado sin contenido que mostrar
	w.WriteHeader(http.StatusNoContent)
}
