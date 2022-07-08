package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Configuracion de conexcion para la base de datos
var DSN = "host=localhost user=rodo password=mypassword dbname=my_task port=5432"

var DB *gorm.DB

func DBConnection() {
	var error error
	// Se realiza conexion a base de datos con los parametros de DSN
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		// En caso de error mostrarlo en el log
		log.Fatal(error)
	} else {
		//En caso de conectarse mostrar el siguiente mensaje
		log.Println("DB Connected")
	}
}
