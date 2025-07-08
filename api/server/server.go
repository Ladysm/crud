package server

import (
	"crud/api/router"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	//aqui estoy llamando ala función que genera la conexion a al DB
	db := InitDB()
	r := gin.Default()
	// se pasa la conexión a las rutas
	router.InitRoutes(r, db)
	r.Run(":8080")

}
