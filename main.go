package main

import (
	"net/http"

	"github.com/Rodomaxi2/rest_go_task/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
