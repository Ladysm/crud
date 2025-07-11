// aqui indico que este archivo pertenece a la carpeta server
package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //  IMPORTANTE: sin esto, no se registra el driver
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {

	//aqui cargo las variable de entorno
	err := godotenv.Load("api/env/.env")
	if err != nil {
		log.Fatal("error al cargar archivo .env")
	}
	// leo las variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// esto es ikhe pa armar el string donde los simbolos %s reemplaza
	//cada uno de los items
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	// aqui se conecta
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("error al conectar", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("ping fallo ", err)
	}
	fmt.Println("conexión exitosa")
	return db
}
