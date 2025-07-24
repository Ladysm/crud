package server

import (
	"crud/api/router"
	"database/sql"
)

// encargada de iniciar el servidor web
func StartServer(addr string, db *sql.DB) {
	//Crea una nueva instancia del motor Gin con configuraciones por defecto

	// se pasa la conexión a las rutas
	r := router.InitRoutes(db)
	//arranca el servidor en el puerto 8080
	r.Run(addr)

}

// func StartServer() {
// 	fmt.Println("Nada aquí por ahora") // o lo puedes dejar vacío
// }
