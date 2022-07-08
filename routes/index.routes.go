package routes

import "net/http"

// Manejador para ruta home "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello World!")))
}
