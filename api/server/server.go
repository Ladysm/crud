package server

import (
	"crud/api/router"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// encargada de iniciar el servidor web
func StartServer(db *sql.DB) {
	//Crea una nueva instancia del motor Gin con configuraciones por defecto
	r := gin.Default()
	// se pasa la conexión a las rutas
	router.InitRoutes(r, db)
	//arranca el servidor en el puerto 8080
	r.Run(":8080")

}

// func StartServer() {
// 	fmt.Println("Nada aquí por ahora") // o lo puedes dejar vacío
// }
