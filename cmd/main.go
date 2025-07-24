package main

import "crud/api/server"

func main() {
	//inciializar conexión db

	db := server.InitDB()
	// arranca el servidor HTTP con gin y pasando la conexión a la db para que las rutan las usen
	server.StartServer(":8080", db)
}
